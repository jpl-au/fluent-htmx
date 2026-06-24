package htmx

import (
	"net/http"
	"net/http/httptest"
	"testing"
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

	if err := sse.Send("message", "<div>hello</div>"); err != nil {
		t.Fatalf("Send() returned error: %v", err)
	}

	want := "event: message\ndata: <div>hello</div>\n\n"
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

	if err := sse.Send("update", "<div>\n  <p>line one</p>\n  <p>line two</p>\n</div>"); err != nil {
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

	if err := sse.Send("first", "one"); err != nil {
		t.Fatalf("Send(first) returned error: %v", err)
	}

	if err := sse.Send("second", "two"); err != nil {
		t.Fatalf("Send(second) returned error: %v", err)
	}

	want := "event: first\ndata: one\n\nevent: second\ndata: two\n\n"
	if got := w.Body.String(); got != want {
		t.Errorf("body = %q, want %q", got, want)
	}
}
