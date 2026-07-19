package htmx

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/jpl-au/fluent/node"
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

// Send writes a named SSE event carrying the rendered node as its data payload.
// The node is rendered with RenderBuilder and each physical line of the output is
// sent with its own data: prefix, so a multi-line fragment stays one valid event
// per the SSE specification. A nil node sends a single empty data line, which suits
// signal-only events such as one that closes the stream. The response is flushed
// after each event to ensure immediate delivery to the client.
//
// Send is the fluent path. When the payload is not a fluent node, for example
// markup from another template engine or a cached fragment, use SendBytes.
//
// The event name should match what the client expects in sse-swap or sse-close
// attributes.
func (s *SSEWriter) Send(event string, n node.Node) error {
	if n == nil {
		return s.SendBytes(event, nil)
	}

	var buf bytes.Buffer
	n.RenderBuilder(&buf)

	return s.SendBytes(event, buf.Bytes())
}

// SendBytes writes a named SSE event whose data is the given bytes, without
// involving fluent. It is the escape hatch for payloads that are not fluent
// nodes: markup from another template engine, a cached fragment, or plain text.
// Each newline-separated line of data is written with its own data: prefix so a
// multi-line payload stays one valid event, and a nil or empty payload sends a
// single empty data line. The response is flushed after the event.
func (s *SSEWriter) SendBytes(event string, data []byte) error {
	if _, err := fmt.Fprintf(s.w, "event: %s\n", event); err != nil {
		return fmt.Errorf("failed to write SSE event: %w", err)
	}

	for line := range bytes.SplitSeq(data, []byte("\n")) {
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
