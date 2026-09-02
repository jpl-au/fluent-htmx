package htmx

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/jpl-au/fluent/node"
)

// SSEWriter sends Server-Sent Events over an HTTP response. It pairs with the
// client-side SSE extension (SSEConnect, SSEClose) to enable real-time updates from
// server to browser. An unnamed message (Swap) is swapped into the connecting element's
// target; a named event (Send) is dispatched as a DOM event and not swapped. There is no
// sse-swap in htmx 4. The client asks for the stream with Accept: text/event-stream and
// opens it on the element's hx-trigger, or on load.
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

// Event is one Server-Sent Event. Name is what the client listens for: htmx dispatches each
// named event as a DOM event, and an unnamed event's data is swapped automatically. ID, when
// set, is remembered by the browser and sent back as Last-Event-ID when the connection is
// re-established, so the handler can resume from that point. Retry, when set, tells the client
// how long to wait before reconnecting. Data is written one data: line per newline-separated
// line, so a multi-line fragment stays one valid event.
type Event struct {
	Name  string
	ID    string
	Retry time.Duration
	Data  []byte
}

// ReleaseEvent is the event name that tells the htmx SSE extension to release the request
// that opened the stream, so the element's request cycle completes while the stream stays
// open. The stream's config decides when release happens on its own (releaseOn); sending this
// event releases it now. See [SSEWriter.Release].
const ReleaseEvent = "hx:release"

// Swap sends an unnamed message carrying the rendered node, which the client swaps
// into the connecting element's target with its hx-swap style. This is the call for
// pushing content; a named event from Send is dispatched as a DOM event and is not
// swapped. A nil node sends an empty message, which the client ignores.
func (s *SSEWriter) Swap(n node.Node) error {
	return s.Send("", n)
}

// SwapBytes sends an unnamed message carrying the given bytes, which the client
// swaps into the connecting element's target. It is Swap for a payload that is not
// a fluent node.
func (s *SSEWriter) SwapBytes(data []byte) error {
	return s.SendBytes("", data)
}

// Send writes a named SSE event carrying the rendered node as its data payload.
// The client dispatches a named event as a DOM event of that name on the connecting
// element, with the data in the event detail, and does not swap it; listen with
// HxTrigger or HxOn. Use Swap to push content into the page. A nil node sends a
// single empty data line, which suits signal-only events such as the one named in
// hx-sse:close. An empty event name sends an unnamed message, the same as Swap.
// The response is flushed after each event.
//
// Send is the fluent path. When the payload is not a fluent node, use SendBytes;
// to set an event id or retry interval, use SendEvent.
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
// A nil or empty payload sends a single empty data line. The response is flushed
// after the event.
func (s *SSEWriter) SendBytes(event string, data []byte) error {
	return s.SendEvent(Event{Name: event, Data: data})
}

// SendEvent writes one event with every field the SSE format allows: the name, an
// id the browser echoes back as Last-Event-ID on reconnect, a retry interval, and
// the data. Fields left at their zero value are omitted, except that an event with
// no data still sends one empty data line so the client delivers it. The response
// is flushed after the event.
func (s *SSEWriter) SendEvent(e Event) error {
	if e.Name != "" {
		if _, err := fmt.Fprintf(s.w, "event: %s\n", e.Name); err != nil {
			return fmt.Errorf("failed to write SSE event: %w", err)
		}
	}
	if e.ID != "" {
		if _, err := fmt.Fprintf(s.w, "id: %s\n", e.ID); err != nil {
			return fmt.Errorf("failed to write SSE id: %w", err)
		}
	}
	if e.Retry > 0 {
		if _, err := fmt.Fprintf(s.w, "retry: %d\n", e.Retry.Milliseconds()); err != nil {
			return fmt.Errorf("failed to write SSE retry: %w", err)
		}
	}

	for line := range bytes.SplitSeq(e.Data, []byte("\n")) {
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

// Release sends the hx:release event, which completes the request that opened the
// stream on the client while leaving the stream open. Use it when the element's
// request indicators and after-request handlers should run once the stream is
// established rather than when it ends.
func (s *SSEWriter) Release() error {
	return s.SendEvent(Event{Name: ReleaseEvent})
}
