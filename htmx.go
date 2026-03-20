// Package htmx provides fluent wrappers for HTMX attributes and server-side helpers.
//
// This package is organised into these main components:
//   - htmx.go: Core wrapper struct and constructor (this file)
//   - client.go: Client-side HTMX attribute methods (hx-get, hx-post, etc.)
//   - server.go: Server-side helpers for handling HTMX requests and responses
//   - config.go: HTMX configuration builder for generating htmx.config settings
//
// Extension support is provided in separate files:
//   - ws.go: WebSocket extension (ws-connect, ws-send)
//   - sse.go: Server-Sent Events extension (sse-connect, sse-swap, sse-close)
//   - sse_server.go: Server-side SSE writer for sending events to the browser
//   - preload.go: Preload extension (preload, preload-images)
//   - response_targets.go: Response targets extension (hx-target-error, hx-target-*)
//   - head_support.go: Head support extension (hx-head)
//
// Usage:
//
//	import (
//	    "github.com/jpl-au/fluent/html5/div"
//	    "github.com/jpl-au/fluent-htmx"
//	)
//
//	// Using HTMX attributes
//	d := div.New()
//	w := htmx.New(d)
//	w.HxGet("/api/data").HxTarget("#result").HxSwap("innerHTML")
//
//	// Configuring HTMX
//	cfg := htmx.Config().
//	    DefaultSwapStyle("outerHTML").
//	    Timeout(5000).
//	    GlobalViewTransitions(true)
//	metaTag, err := cfg.ToMetaTag()
package htmx

import (
	"bytes"
	"io"

	"github.com/jpl-au/fluent/node"
)

// Wrapper adds HTMX attribute methods to any node.Element.
// It delegates all node.Element methods to the wrapped element so it can be
// used anywhere the original element would be. All HTMX methods return *Wrapper
// to enable fluent chaining.
type Wrapper struct {
	element node.Element
}

// New wraps an element to enable HTMX attribute chaining.
func New(n node.Element) *Wrapper {
	return &Wrapper{element: n}
}

// node.Element delegation - all calls pass through to the wrapped element.

func (h *Wrapper) Render(w ...io.Writer) []byte          { return h.element.Render(w...) }
func (h *Wrapper) RenderBuilder(buf *bytes.Buffer)       { h.element.RenderBuilder(buf) }
func (h *Wrapper) Nodes() []node.Node                    { return h.element.Nodes() }
func (h *Wrapper) SetAttribute(key string, value string) { h.element.SetAttribute(key, value) }
func (h *Wrapper) RenderOpen(buf *bytes.Buffer)          { h.element.RenderOpen(buf) }
func (h *Wrapper) RenderClose(buf *bytes.Buffer)         { h.element.RenderClose(buf) }
