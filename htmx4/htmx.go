// Package htmx provides fluent wrappers for HTMX attributes and server-side helpers,
// targeting htmx 4.
//
// This package is organised into these main components:
//   - htmx.go: the Wrapper type and all client-side hx-* attribute methods (this file)
//   - inherit.go: the attribute inheritance modifier (Mod, Inherited, InheritedAppend)
//   - server.go: server-side helpers for handling HTMX requests and responses
//   - config.go: HTMX configuration builder for generating htmx.config settings
//
// Extension support is provided in separate files. These are bundled into the htmax.js build
// unless noted otherwise:
//   - ws.go: WebSocket extension (hx-ws:connect, hx-ws:send)
//   - sse.go: Server-Sent Events extension (hx-sse:connect, hx-sse:close)
//   - sse_server.go: Server-side SSE writer for sending events to the browser
//   - preload.go: Preload extension (hx-preload)
//   - optimistic.go: Optimistic UI extension (hx-optimistic)
//   - targets.go: Multi-target extension (hx-targets)
//   - live.go: Reactive expressions extension (hx-live)
//   - browser_indicator.go: Native loading-indicator extension (hx-browser-indicator)
//   - download.go: Download extension (swap.Download, HxDownload, HX-Download)
//   - status.go: Status-code swap control (hx-status:CODE)
//
// These wrap htmx 4 extensions shipped as separate scripts (not bundled in htmax.js):
//   - ptag.go: Polling-tag extension (hx-ptag)
//   - history_cache.go: History-cache extension (hx-history)
//   - csp.go: CSP nonce extension (hx-nonce)
//   - head_support.go: hx-head extension (hx-head)
//   - prompt.go: Prompt extension (hx-prompt), restored from htmx 2 in beta5
//   - swap.Upsert: Upsert extension (hx-swap="upsert")
//
// Two builds, and what happens if an extension is not loaded. htmx 4 ships two scripts:
// htmax.js bundles the eight popular extensions (ws, sse, preload, optimistic, targets, live,
// browser-indicator, download) so their attributes work with no extra includes, while the core
// htmx.js build contains none of them - with it you include each extension's script yourself,
// the same as the always-separate extensions (hx-head, hx-ptag, hx-history, hx-nonce, hx-prompt, upsert).
//
// An extension method only writes an attribute (or, server-side, a response header) - it never
// loads any JavaScript. If the matching extension is not present, htmx does not recognise the
// attribute and ignores it: the element still renders, core htmx still works, and the
// extension's behaviour simply does not happen. So using an extension method without its
// extension is a harmless no-op, never an error - the cost is silent, which is why each
// extension method's doc names the script it needs. (Unlike htmx 2, htmx 4 does not need hx-ext
// to switch an included extension on; the Config().Extensions allowlist only restricts which run.)
//
// Attribute inheritance: by default an attribute applies only to the element it is set on.
// Inheritable setters take an optional modifier - pass htmx.Inherited to inherit the
// attribute to descendant elements, or htmx.InheritedAppend so a descendant appends to the
// inherited value instead of replacing it. See [Mod].
//
// Usage:
//
//	import (
//	    "github.com/jpl-au/fluent/html5/div"
//	    "github.com/jpl-au/fluent-htmx/htmx4"
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
//	    DefaultTimeout(5000).
//	    Transitions(true)
//	metaTag, err := cfg.ToMetaTag()
package htmx

import (
	"bytes"
	"io"

	"github.com/jpl-au/fluent-htmx/htmx4/swap"
	"github.com/jpl-au/fluent-htmx/htmx4/sync"
	"github.com/jpl-au/fluent/node"
)

// Wrapper adds HTMX attribute methods to any node.Element.
// It delegates all node.Element methods to the wrapped element so it can be
// used anywhere the original element would be. All HTMX methods return *Wrapper
// to enable fluent chaining.
type Wrapper struct {
	element node.Element
}

// New wraps an element so htmx's hx-* attribute methods chain onto it, returning a
// *Wrapper whose fluent setters build up the element's client-side behaviour. The
// Wrapper delegates the full node.Element surface to the wrapped element, so it renders
// identically and can be used anywhere the original element could.
func New(n node.Element) *Wrapper {
	return &Wrapper{element: n}
}

// node.Element delegation - all calls pass through to the wrapped element.

func (h *Wrapper) Render(w io.Writer)                    { h.element.Render(w) }
func (h *Wrapper) WriteTo(w io.Writer) (int64, error)    { return h.element.WriteTo(w) }
func (h *Wrapper) RenderBytes() []byte                   { return h.element.RenderBytes() }
func (h *Wrapper) RenderBuilder(buf *bytes.Buffer)       { h.element.RenderBuilder(buf) }
func (h *Wrapper) Nodes() []node.Node                    { return h.element.Nodes() }
func (h *Wrapper) SetAttribute(key string, value string) { h.element.SetAttribute(key, value) }

// SetAttributeRaw sets a custom attribute without escaping, delegating to the
// wrapped element (node.Element's raw hatch).
func (h *Wrapper) SetAttributeRaw(key string, value string) { h.element.SetAttributeRaw(key, value) }
func (h *Wrapper) RenderOpen(buf *bytes.Buffer)             { h.element.RenderOpen(buf) }
func (h *Wrapper) RenderClose(buf *bytes.Buffer)            { h.element.RenderClose(buf) }

// HxGet issues an AJAX GET request to url when the element is triggered, then
// swaps the returned HTML into the page. By default the response replaces the
// element's own inner content; pair it with HxTarget and HxSwap to place the
// content elsewhere or change how it is inserted. GET sends no body and should
// stay idempotent, so use it to fetch and show server-rendered fragments rather
// than to change state.
func (h *Wrapper) HxGet(url string) *Wrapper {
	h.element.SetAttribute("hx-get", url)

	return h
}

// HxPost issues an AJAX POST request to url when the element is triggered, then
// swaps the returned HTML into the page. Enclosed form values, plus anything
// added with hx-include, are serialised into the request body, making POST the
// verb for creating data or otherwise mutating state. The response swaps into
// the element itself unless HxTarget and HxSwap redirect it.
func (h *Wrapper) HxPost(url string) *Wrapper {
	h.element.SetAttribute("hx-post", url)

	return h
}

// HxPut issues an AJAX PUT request to url when the element is triggered, then
// swaps the returned HTML into the page. PUT carries the enclosed form values as
// its body and conventionally replaces a resource wholesale; htmx sends the real
// method, so no _method override field is needed. The response swaps into the
// element unless HxTarget and HxSwap redirect it.
func (h *Wrapper) HxPut(url string) *Wrapper {
	h.element.SetAttribute("hx-put", url)

	return h
}

// HxPatch issues an AJAX PATCH request to url when the element is triggered, then
// swaps the returned HTML into the page. PATCH carries the enclosed form values
// as its body and conventionally applies a partial update to a resource; htmx
// sends the real method directly. The response swaps into the element unless
// HxTarget and HxSwap redirect it.
func (h *Wrapper) HxPatch(url string) *Wrapper {
	h.element.SetAttribute("hx-patch", url)

	return h
}

// HxDelete issues an AJAX DELETE request to url when the element is triggered,
// then swaps the returned HTML into the page. Use it to remove a resource; the
// server's response, often the surrounding list re-rendered or empty to clear
// the element, swaps into the element unless HxTarget and HxSwap redirect it.
//
// Like GET, a DELETE does not include the enclosing form's inputs; add
// HxInclude("closest form") where that form data is needed.
func (h *Wrapper) HxDelete(url string) *Wrapper {
	h.element.SetAttribute("hx-delete", url)

	return h
}

// HxAction specifies the request URL when the HTTP method is set separately with HxMethod.
// Use this when the verb is dynamic; otherwise prefer the verb methods (HxGet, HxPost, ...).
func (h *Wrapper) HxAction(url string) *Wrapper {
	h.element.SetAttribute("hx-action", url)

	return h
}

// HxMethod specifies the HTTP method for a request whose URL is set with HxAction.
// Example: HxMethod("post").
func (h *Wrapper) HxMethod(method string) *Wrapper {
	h.element.SetAttribute("hx-method", method)

	return h
}

// HxSwap controls how the response content is swapped into the DOM.
// Use the predefined swap package constants or swap.Custom() for strategies with modifiers.
func (h *Wrapper) HxSwap(strategy swap.Strategy, mods ...Mod) *Wrapper {
	h.element.SetAttribute(modifiedKey("hx-swap", mods), string(strategy))

	return h
}

// HxTarget specifies a CSS selector for the element that will receive the swapped content.
// Without this, the element that triggers the request is the swap target.
func (h *Wrapper) HxTarget(target string, mods ...Mod) *Wrapper {
	h.element.SetAttribute(modifiedKey("hx-target", mods), target)

	return h
}

// HxTrigger specifies which DOM events cause the element to issue a request.
// Supports standard DOM events, HTMX-specific events, and modifiers.
// Example: "keyup changed delay:500ms" triggers on keyup after 500ms of inactivity.
//
// Request queuing is controlled with HxSync, not a hx-trigger modifier
// (e.g. HxSync(sync.Custom("this:queue all"))).
func (h *Wrapper) HxTrigger(trigger string, mods ...Mod) *Wrapper {
	h.element.SetAttribute(modifiedKey("hx-trigger", mods), trigger)

	return h
}

// HxBoost converts standard links and forms into AJAX requests,
// swapping the response into the page body without a full page reload.
func (h *Wrapper) HxBoost(enabled bool, mods ...Mod) *Wrapper {
	value := boolFalse
	if enabled {
		value = boolTrue
	}

	h.element.SetAttribute(modifiedKey("hx-boost", mods), value)

	return h
}

// HxConfirm shows a browser confirmation dialog before issuing the request.
// The request is only sent if the user confirms.
func (h *Wrapper) HxConfirm(message string, mods ...Mod) *Wrapper {
	h.element.SetAttribute(modifiedKey("hx-confirm", mods), message)

	return h
}

// HxVals adds extra JSON-encoded values to the request parameters.
// Example: `{"key": "value"}`. For computed values use a "js:" prefix.
func (h *Wrapper) HxVals(values string, mods ...Mod) *Wrapper {
	h.element.SetAttribute(modifiedKey("hx-vals", mods), values)

	return h
}

// HxHeaders adds extra JSON-encoded headers to the AJAX request.
// Example: `{"X-Custom-Header": "value"}`.
func (h *Wrapper) HxHeaders(headers string, mods ...Mod) *Wrapper {
	h.element.SetAttribute(modifiedKey("hx-headers", mods), headers)

	return h
}

// HxIndicator specifies a CSS selector for an element to show while the request is in flight.
// The targeted element receives the htmx-request class, which can be styled to show a spinner.
func (h *Wrapper) HxIndicator(indicator string, mods ...Mod) *Wrapper {
	h.element.SetAttribute(modifiedKey("hx-indicator", mods), indicator)

	return h
}

// HxPushURL pushes a URL into the browser history after the request completes.
// Accepts "true" to push the fetched URL, "false" to disable, or a custom URL string.
func (h *Wrapper) HxPushURL(value string, mods ...Mod) *Wrapper {
	h.element.SetAttribute(modifiedKey("hx-push-url", mods), value)

	return h
}

// HxSelect picks a CSS selector from the response HTML to swap in,
// discarding the rest of the response. Useful when the server returns a full page
// but only a fragment is needed.
func (h *Wrapper) HxSelect(selector string, mods ...Mod) *Wrapper {
	h.element.SetAttribute(modifiedKey("hx-select", mods), selector)

	return h
}

// HxSelectOOB selects content from the response for out-of-band swaps.
// These elements are swapped into matching targets elsewhere in the DOM,
// independently of the primary swap target.
//
// hx-select-oob is read off the requesting element and resolves with inheritance,
// so it accepts the inheritance modifier - unlike HxSwapOOB, which htmx matches by id
// on the response and so takes no modifier.
func (h *Wrapper) HxSelectOOB(selector string, mods ...Mod) *Wrapper {
	h.element.SetAttribute(modifiedKey("hx-select-oob", mods), selector)

	return h
}

// HxSwapOOB marks response content for out-of-band swapping.
// The element is swapped into the DOM by matching its ID, regardless of the primary swap target.
// Typically set on server-rendered response fragments rather than request elements.
func (h *Wrapper) HxSwapOOB(value string) *Wrapper {
	h.element.SetAttribute("hx-swap-oob", value)

	return h
}

// HxReplaceURL replaces the current URL in the browser location bar without adding a history entry.
// Unlike HxPushURL, the user cannot navigate back to the previous URL.
func (h *Wrapper) HxReplaceURL(url string, mods ...Mod) *Wrapper {
	h.element.SetAttribute(modifiedKey("hx-replace-url", mods), url)

	return h
}

// HxInclude includes values from other elements in the request using a CSS selector.
// Useful for submitting inputs that are outside the triggering element's form.
func (h *Wrapper) HxInclude(selector string, mods ...Mod) *Wrapper {
	h.element.SetAttribute(modifiedKey("hx-include", mods), selector)

	return h
}

// HxSync coordinates requests between this element and another element matched by the selector.
// Prevents race conditions when multiple elements can trigger overlapping requests.
// Request queuing is also expressed here, via a queue strategy.
//
//	htmx.New(el).HxSync(sync.Drop)
//	htmx.New(el).HxSync(sync.Custom("this:queue all"))
func (h *Wrapper) HxSync(strategy sync.Strategy, mods ...Mod) *Wrapper {
	h.element.SetAttribute(modifiedKey("hx-sync", mods), string(strategy))

	return h
}

// HxConfig sets per-element request configuration as JSON or htmx's key:value (HCON) syntax.
// Example: HxConfig(`{"timeout": 5000}`) or HxConfig("timeout:5000 credentials:'include'").
//
// htmx resolves hx-config with inheritance, so it accepts the inheritance modifier. The
// mode key is ignored per element - htmx never lets a per-element config override the
// security-sensitive fetch mode.
func (h *Wrapper) HxConfig(config string, mods ...Mod) *Wrapper {
	h.element.SetAttribute(modifiedKey("hx-config", mods), config)

	return h
}

// HxEncoding changes the request encoding type.
// The only non-default value is "multipart/form-data" for file uploads.
func (h *Wrapper) HxEncoding(encoding string, mods ...Mod) *Wrapper {
	h.element.SetAttribute(modifiedKey("hx-encoding", mods), encoding)

	return h
}

// HxPreserve keeps the element unchanged during swaps by matching its ID.
// Useful for persistent elements like video players or iframes that should
// survive content updates around them.
//
// htmx selects preserved elements by the presence of hx-preserve, never by its value, so
// this is a no-argument setter. To stop preserving an element, omit the call entirely:
// there is no "off" value - hx-preserve="false" would still preserve the element.
func (h *Wrapper) HxPreserve() *Wrapper {
	h.element.SetAttribute("hx-preserve", boolTrue)

	return h
}

// HxHistoryElt designates this element as the snapshot source for history navigation.
// When the user navigates back, htmx re-fetches and swaps into this element instead of <body>.
func (h *Wrapper) HxHistoryElt() *Wrapper {
	h.element.SetAttribute("hx-history-elt", boolTrue)

	return h
}

// HxIgnore prevents htmx from processing this element and all of its children.
// The attribute's presence alone is sufficient; its value is ignored, and it cannot be
// re-enabled by any descendant content.
func (h *Wrapper) HxIgnore() *Wrapper {
	h.element.SetAttribute("hx-ignore", boolTrue)

	return h
}

// HxDisable disables the matched form elements while a request is in flight.
// Useful for preventing duplicate submissions by disabling the submit button.
func (h *Wrapper) HxDisable(selector string, mods ...Mod) *Wrapper {
	h.element.SetAttribute(modifiedKey("hx-disable", mods), selector)

	return h
}

// HxValidate forces the browser's native form validation before issuing the request.
// The request is blocked if any included input fails its validation constraints.
func (h *Wrapper) HxValidate(validate bool, mods ...Mod) *Wrapper {
	value := boolFalse
	if validate {
		value = boolTrue
	}

	h.element.SetAttribute(modifiedKey("hx-validate", mods), value)

	return h
}

// HxOn attaches an inline event handler directly to the element (locality of behaviour).
// The syntax is hx-on:EVENT (a single colon), so the handler lives with the element it
// controls rather than in a separate script tag. Pass the event name as htmx dispatches
// it, e.g. "click" or "htmx:after:swap".
// Example: HxOn("click", "alert('hi')") → hx-on:click="alert('hi')".
func (h *Wrapper) HxOn(event string, handler string) *Wrapper {
	h.element.SetAttribute("hx-on:"+event, handler)

	return h
}
