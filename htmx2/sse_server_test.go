package htmx

import (
	"net/http"
	"net/http/httptest"
	"testing"

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

	want := "event: done\ndata: \n\n"
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

	want := "event: done\ndata: \n\n"
	if got := w.Body.String(); got != want {
		t.Errorf("body = %q, want %q", got, want)
	}
}
