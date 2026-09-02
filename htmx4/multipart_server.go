package htmx

import (
	"bytes"
	"fmt"
	stdmultipart "mime/multipart"
	"net/http"
	"net/textproto"

	"github.com/jpl-au/fluent-htmx/htmx4/multipart"
	"github.com/jpl-au/fluent-htmx/htmx4/swap"
	"github.com/jpl-au/fluent/node"
)

// MultipartWriter streams a multipart htmx response, one part per swap. It pairs with the
// client-side multipart extension (MultipartConnect): each part carries its own HX-* headers, so
// a single response can update several targets. A multipart.Mixed stream swaps parts in order; a
// multipart.Parallel stream lets the client swap each part as it arrives.
type MultipartWriter struct {
	mw *stdmultipart.Writer
	f  http.Flusher
}

// NewMultipart begins a multipart htmx response of the given type. It sets the Content-Type,
// including the boundary it generates, and returns a writer for streaming parts. It returns an
// error if the ResponseWriter cannot flush, which multipart streaming requires to deliver each
// part as it is written rather than buffering the whole response.
func NewMultipart(w http.ResponseWriter, t multipart.Type) (*MultipartWriter, error) {
	f, ok := w.(http.Flusher)
	if !ok {
		return nil, fmt.Errorf("ResponseWriter does not implement http.Flusher")
	}

	mw := stdmultipart.NewWriter(w)
	w.Header().Set("Content-Type", fmt.Sprintf("%s; boundary=%s", t, mw.Boundary()))
	// A streamed, reconnecting response must not be cached or replayed by an intermediary.
	w.Header().Set("Cache-Control", "no-cache")

	return &MultipartWriter{mw: mw, f: f}, nil
}

// WritePart streams one part: n is rendered as the part body, and each PartOption sets an HX-*
// header directing how that part is swapped. The part is flushed immediately so the client can
// act on it without waiting for the rest of the response. With no options the part swaps using
// the request's own target and swap style.
//
// A nil node writes a header-only part, but that is not a no-op swap: the extension still runs a
// swap with empty content, so a nil node clears the target under the default swap style. For a
// part that only fires an event (via PartTrigger) and changes nothing, pair the nil node with
// PartSwap(swap.None).
func (m *MultipartWriter) WritePart(n node.Node, opts ...PartOption) error {
	// Every part declares Content-Type: text/html. Besides being accurate, it guarantees the
	// part carries at least one header, which the client's multipart parser relies on to
	// separate the header block from the body: a header-less part leaves only a single CRLF
	// after the boundary line, and the parser would then read body bytes as headers.
	header := textproto.MIMEHeader{"Content-Type": {"text/html; charset=utf-8"}}
	for _, opt := range opts {
		opt(header)
	}

	part, err := m.mw.CreatePart(header)
	if err != nil {
		return fmt.Errorf("failed to create multipart part: %w", err)
	}

	if n != nil {
		var buf bytes.Buffer
		n.RenderBuilder(&buf)

		if _, err := part.Write(buf.Bytes()); err != nil {
			return fmt.Errorf("failed to write multipart part: %w", err)
		}
	}

	m.f.Flush()

	return nil
}

// Close writes the closing boundary that ends the multipart response and flushes it. No further
// parts may be written after Close.
func (m *MultipartWriter) Close() error {
	if err := m.mw.Close(); err != nil {
		return fmt.Errorf("failed to close multipart response: %w", err)
	}

	m.f.Flush()

	return nil
}

// PartOption sets an HX-* header on a single multipart part, directing how that part is swapped.
// The options set the override headers (HX-Retarget, HX-Reswap, HX-Reselect) so a part's own
// directives always win over any default the request carried.
type PartOption func(textproto.MIMEHeader)

// PartTarget swaps the part into the element matching selector, via HX-Retarget.
func PartTarget(selector string) PartOption {
	// Assign the map directly rather than via MIMEHeader.Set: Set would canonicalise the key to
	// Hx-Retarget, whereas htmx's convention (and every other HX-* header here) is HX-Retarget.
	// The extension lowercases header names, so both resolve, but the wire form should match htmx.
	return func(h textproto.MIMEHeader) { h[HXRetargetHeader] = []string{selector} }
}

// PartSwap sets how the part is swapped into its target, overriding the request's swap style. It
// sets HX-Reswap.
func PartSwap(strategy swap.Strategy) PartOption {
	return func(h textproto.MIMEHeader) { h[HXReswapHeader] = []string{string(strategy)} }
}

// PartSelect swaps only the fragment of the part matching selector, via HX-Reselect.
func PartSelect(selector string) PartOption {
	return func(h textproto.MIMEHeader) { h[HXReselectHeader] = []string{selector} }
}

// PartTrigger fires client events through the part's HX-Trigger header. They fire before the
// part's content is swapped. The value is a comma-separated list of event names. The 4.0.0
// extension cannot parse the JSON object form: it calls an undefined htmx.__HCON and the
// part fails with htmx:multipart:error, so send names only.
func PartTrigger(events string) PartOption {
	return func(h textproto.MIMEHeader) { h[HXTriggerHeader] = []string{events} }
}

// PartRefresh makes the client reload the page through the part's HX-Refresh header. The
// part's content is not swapped and the connection closes, so send it with a nil node.
func PartRefresh() PartOption {
	return func(h textproto.MIMEHeader) { h[HXRefreshHeader] = []string{boolTrue} }
}

// PartRedirect makes the client navigate to url with a full page load through the part's
// HX-Redirect header. The part's content is not swapped and the connection closes, so send
// it with a nil node.
func PartRedirect(url string) PartOption {
	return func(h textproto.MIMEHeader) { h[HXRedirectHeader] = []string{url} }
}

// PartLocation makes the client fetch url and swap it in without a full page load through
// the part's HX-Location header. The part's content is not swapped, and the connection stays
// open, so send it with a nil node. The value must be a plain path without spaces or commas:
// the 4.0.0 extension routes the object form, and any value holding a space or comma,
// through an undefined htmx.__HCON and the part fails with htmx:multipart:error.
func PartLocation(url string) PartOption {
	return func(h textproto.MIMEHeader) { h[HXLocationHeader] = []string{url} }
}

// PartPushURL pushes url into the browser history when the part swaps, through the part's
// HX-Push-Url header. The value is as for HxPushURL: "true" for the request's own URL,
// "false" to stop a push, or a URL.
func PartPushURL(url string) PartOption {
	return func(h textproto.MIMEHeader) { h[HXPushURLHeader] = []string{url} }
}

// PartReplaceURL replaces the current browser history entry with url when the part swaps,
// through the part's HX-Replace-Url header. The value is as for PartPushURL.
func PartReplaceURL(url string) PartOption {
	return func(h textproto.MIMEHeader) { h[HXReplaceURLHeader] = []string{url} }
}

// PartID gives the part an HX-Part-ID header. The client remembers the id of the last part it
// swapped and sends it back as HX-Last-Part-ID when it reconnects, so a handler can resume the
// stream after that part instead of starting again. A part without the header leaves the
// remembered id as it was, and an empty id clears it, so the next reconnect carries no
// HX-Last-Part-ID. Ids are the server's own scheme; a monotonic counter or a timestamp both
// work.
func PartID(id string) PartOption {
	return func(h textproto.MIMEHeader) { h[HXPartIDHeader] = []string{id} }
}
