package htmx

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

// SSEWriter sends Server-Sent Events over an HTTP response. It pairs with the
// client-side SSE extension (SSEConnect, SSESwap, SSEClose) to enable real-time
// updates from server to browser.
type SSEWriter struct {
	w io.Writer
	f http.Flusher
}

// NewSSE initialises a Server-Sent Events stream. It sets the required headers
// and returns a writer for sending events. Returns an error if the
// ResponseWriter does not support flushing, which is required for SSE to
// deliver events immediately.
func NewSSE(w http.ResponseWriter) (*SSEWriter, error) {
	f, ok := w.(http.Flusher)
	if !ok {
		return nil, fmt.Errorf("ResponseWriter does not implement http.Flusher")
	}

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	return &SSEWriter{w: w, f: f}, nil
}

// Send writes a named SSE event with the given data payload. Multi-line data is
// handled automatically per the SSE specification - each line is sent with its
// own data: prefix. The response is flushed after each event to ensure immediate
// delivery to the client.
//
// The event name should match what the client expects in sse-swap or sse-close
// attributes.
func (s *SSEWriter) Send(event string, data string) error {
	if _, err := fmt.Fprintf(s.w, "event: %s\n", event); err != nil {
		return fmt.Errorf("failed to write SSE event: %w", err)
	}

	for line := range strings.SplitSeq(data, "\n") {
		if _, err := fmt.Fprintf(s.w, "data: %s\n", line); err != nil {
			return fmt.Errorf("failed to write SSE data: %w", err)
		}
	}

	if _, err := fmt.Fprint(s.w, "\n"); err != nil {
		return fmt.Errorf("failed to write SSE terminator: %w", err)
	}

	s.f.Flush()

	return nil
}
