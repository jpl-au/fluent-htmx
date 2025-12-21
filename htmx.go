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

// HtmxWrapper provides fluent chaining for HTMX attributes.
// It wraps any node.Node implementation and adds HTMX-specific methods.
// All methods return *HtmxWrapper to enable method chaining.
type HtmxWrapper struct {
	node node.Node
}

// New creates a new HtmxWrapper around a node.
// The wrapper delegates all node.Node interface methods to the wrapped node.
func New(n node.Node) *HtmxWrapper {
	return &HtmxWrapper{node: n}
}

// Render delegates to the wrapped node's Render method.
func (h *HtmxWrapper) Render(w ...io.Writer) []byte {
	return h.node.Render(w...)
}

// RenderBuilder delegates to the wrapped node's RenderBuilder method.
func (h *HtmxWrapper) RenderBuilder(buf *bytes.Buffer) {
	h.node.RenderBuilder(buf)
}

// Nodes delegates to the wrapped node's Nodes method.
func (h *HtmxWrapper) Nodes() []node.Node {
	return h.node.Nodes()
}

// SetAttribute delegates to the wrapped node's SetAttribute method.
func (h *HtmxWrapper) SetAttribute(key string, value string) {
	h.node.SetAttribute(key, value)
}
