package htmx

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/jpl-au/fluent/text"
)

func TestNewSSE(t *testing.T) {
	w := httptest.NewRecorder()

	sse, err := NewSSE(w)
	if err != nil {
		t.Fatalf("NewSSE() returned error: %v", err)
	}

	if sse == nil {
		t.Fatal("NewSSE() returned nil writer")
	}

	if got := w.Header().Get("Content-Type"); got != "text/event-stream" {
		t.Errorf("Content-Type = %q, want %q", got, "text/event-stream")
	}

	if got := w.Header().Get("Cache-Control"); got != "no-cache" {
		t.Errorf("Cache-Control = %q, want %q", got, "no-cache")
	}

	if got := w.Header().Get("Connection"); got != "keep-alive" {
		t.Errorf("Connection = %q, want %q", got, "keep-alive")
	}
}

// nonFlusher is a minimal ResponseWriter that does not implement http.Flusher.
type nonFlusher struct {
	header http.Header
}

func (nf *nonFlusher) Header() http.Header         { return nf.header }
func (nf *nonFlusher) Write(b []byte) (int, error) { return len(b), nil }
func (nf *nonFlusher) WriteHeader(int)             {}

func TestNewSSENoFlusher(t *testing.T) {
	w := &nonFlusher{header: http.Header{}}

	_, err := NewSSE(w)
	if err == nil {
		t.Error("NewSSE() should return error when ResponseWriter lacks Flusher")
	}
}

func TestSSESend(t *testing.T) {
	w := httptest.NewRecorder()

	sse, err := NewSSE(w)
	if err != nil {
		t.Fatalf("NewSSE() returned error: %v", err)
	}

	if err := sse.Send("message", text.RawText("<div>hello</div>")); err != nil {
		t.Fatalf("Send() returned error: %v", err)
	}

	want := "event: message\ndata: <div>hello</div>\n\n"
	if got := w.Body.String(); got != want {
		t.Errorf("body = %q, want %q", got, want)
	}
}

func TestSSESendNil(t *testing.T) {
	w := httptest.NewRecorder()

	sse, err := NewSSE(w)
	if err != nil {
		t.Fatalf("NewSSE() returned error: %v", err)
	}

	// A nil node is a signal-only event, for example the done event that
	// closes the stream. It sends one empty data line and no payload.
	if err := sse.Send("done", nil); err != nil {
		t.Fatalf("Send() returned error: %v", err)
	}

	want := "event: done\n\n"
	if got := w.Body.String(); got != want {
		t.Errorf("body = %q, want %q", got, want)
	}
}

func TestSSESendMultiline(t *testing.T) {
	w := httptest.NewRecorder()

	sse, err := NewSSE(w)
	if err != nil {
		t.Fatalf("NewSSE() returned error: %v", err)
	}

	// A node whose rendered output spans several lines must produce one
	// data: line per physical line, all within a single event.
	n := text.RawText("<div>\n  <p>line one</p>\n  <p>line two</p>\n</div>")
	if err := sse.Send("update", n); err != nil {
		t.Fatalf("Send() returned error: %v", err)
	}

	want := "event: update\ndata: <div>\ndata:   <p>line one</p>\ndata:   <p>line two</p>\ndata: </div>\n\n"
	if got := w.Body.String(); got != want {
		t.Errorf("body = %q, want %q", got, want)
	}
}

func TestSSESendMultipleEvents(t *testing.T) {
	w := httptest.NewRecorder()

	sse, err := NewSSE(w)
	if err != nil {
		t.Fatalf("NewSSE() returned error: %v", err)
	}

	if err := sse.Send("first", text.Text("one")); err != nil {
		t.Fatalf("Send(first) returned error: %v", err)
	}

	if err := sse.Send("second", text.Text("two")); err != nil {
		t.Fatalf("Send(second) returned error: %v", err)
	}

	want := "event: first\ndata: one\n\nevent: second\ndata: two\n\n"
	if got := w.Body.String(); got != want {
		t.Errorf("body = %q, want %q", got, want)
	}
}

func TestSSESendBytes(t *testing.T) {
	w := httptest.NewRecorder()

	sse, err := NewSSE(w)
	if err != nil {
		t.Fatalf("NewSSE() returned error: %v", err)
	}

	// SendBytes is the escape hatch: the payload is raw bytes from any source,
	// not a fluent node. Multi-line data still splits into one data: line each.
	if err := sse.SendBytes("message", []byte("<div>\n  <p>raw</p>\n</div>")); err != nil {
		t.Fatalf("SendBytes() returned error: %v", err)
	}

	want := "event: message\ndata: <div>\ndata:   <p>raw</p>\ndata: </div>\n\n"
	if got := w.Body.String(); got != want {
		t.Errorf("body = %q, want %q", got, want)
	}
}

func TestSSESendBytesNil(t *testing.T) {
	w := httptest.NewRecorder()

	sse, err := NewSSE(w)
	if err != nil {
		t.Fatalf("NewSSE() returned error: %v", err)
	}

	// A nil payload sends one empty data line, matching Send with a nil node.
	if err := sse.SendBytes("done", nil); err != nil {
		t.Fatalf("SendBytes() returned error: %v", err)
	}

	want := "event: done\n\n"
	if got := w.Body.String(); got != want {
		t.Errorf("body = %q, want %q", got, want)
	}
}

func TestSSESendEvent(t *testing.T) {
	w := httptest.NewRecorder()
	sse, err := NewSSE(w)
	if err != nil {
		t.Fatalf("NewSSE() returned error: %v", err)
	}

	err = sse.SendEvent(Event{Name: "tick", ID: "42", Retry: 2 * time.Second, Data: []byte("<p>a</p>\n<p>b</p>")})
	if err != nil {
		t.Fatalf("SendEvent() returned error: %v", err)
	}

	want := "event: tick\nid: 42\nretry: 2000\ndata: <p>a</p>\ndata: <p>b</p>\n\n"
	if got := w.Body.String(); got != want {
		t.Errorf("body = %q, want %q", got, want)
	}
}

// A block with only an id updates the client's Last-Event-ID without dispatching a message,
// and a fully empty event still sends one empty data line so the client delivers it.
func TestSSESendEventIDOnly(t *testing.T) {
	w := httptest.NewRecorder()
	sse, err := NewSSE(w)
	if err != nil {
		t.Fatalf("NewSSE() returned error: %v", err)
	}

	if err := sse.SendEvent(Event{ID: "42"}); err != nil {
		t.Fatalf("SendEvent() returned error: %v", err)
	}

	if got, want := w.Body.String(), "id: 42\n\n"; got != want {
		t.Errorf("body = %q, want %q", got, want)
	}
}

func TestSSESendEventEmpty(t *testing.T) {
	w := httptest.NewRecorder()
	sse, err := NewSSE(w)
	if err != nil {
		t.Fatalf("NewSSE() returned error: %v", err)
	}

	if err := sse.SendEvent(Event{}); err != nil {
		t.Fatalf("SendEvent() returned error: %v", err)
	}

	if got, want := w.Body.String(), "data: \n\n"; got != want {
		t.Errorf("body = %q, want %q", got, want)
	}
}

func TestSSERelease(t *testing.T) {
	w := httptest.NewRecorder()
	sse, err := NewSSE(w)
	if err != nil {
		t.Fatalf("NewSSE() returned error: %v", err)
	}

	if err := sse.Release(); err != nil {
		t.Fatalf("Release() returned error: %v", err)
	}

	if got, want := w.Body.String(), "event: hx:release\n\n"; got != want {
		t.Errorf("body = %q, want %q", got, want)
	}
}

// An unnamed message is what the client swaps, so Swap must write no event line.
func TestSSESwap(t *testing.T) {
	w := httptest.NewRecorder()
	sse, err := NewSSE(w)
	if err != nil {
		t.Fatalf("NewSSE() returned error: %v", err)
	}

	if err := sse.Swap(text.RawText("<p>x</p>")); err != nil {
		t.Fatalf("Swap() returned error: %v", err)
	}
	if err := sse.SwapBytes([]byte("<p>y</p>")); err != nil {
		t.Fatalf("SwapBytes() returned error: %v", err)
	}

	if got, want := w.Body.String(), "data: <p>x</p>\n\ndata: <p>y</p>\n\n"; got != want {
		t.Errorf("body = %q, want %q", got, want)
	}
}
