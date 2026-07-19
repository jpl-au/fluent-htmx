package htmx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/jpl-au/fluent-htmx/htmx4/swap"
	"github.com/jpl-au/fluent/node"
)

// triggerEvent represents a single HTMX trigger event.
type triggerEvent struct {
	Name    string
	Details any
}

// TriggerBuilder accumulates client-side events to fire after a response and writes them
// as the HX-Trigger response header. All events fire immediately when the response is
// received. Use it to notify the page that something happened server-side (for example an
// item was saved) without coupling that signal to the swapped content.
type TriggerBuilder struct {
	w        http.ResponseWriter
	triggers []triggerEvent
}

// HxRequest returns true if the request was initiated by HTMX.
// HTMX sends the HX-Request header with every AJAX request it makes.
func HxRequest(r *http.Request) bool {
	return r.Header.Get(HXRequestHeader) == boolTrue
}

// Handle executes the closure if request is from HTMX and returns true to signal early return.
// This enables clean separation of HTMX partial responses from full page responses:
//
//	func MyHandler(w http.ResponseWriter, r *http.Request) {
//	    if htmx.Handle(r, func() {
//	        // HTMX request: render partial and set headers
//	        partial.Render(w)
//	        htmx.HxPushURL(w, "/new-path")
//	    }) {
//	        return // Early return for HTMX requests
//	    }
//	    // Non-HTMX request: render full page
//	    fullPage.Render(w)
//	}
func Handle(r *http.Request, fn func()) bool {
	if HxRequest(r) {
		fn()

		return true
	}

	return false
}

// HxBoosted returns true if the request came from an element with hx-boost enabled.
// Boosted requests behave like standard navigation but use AJAX - the server may want
// to return a full page layout for boosted requests but a partial for regular HTMX requests.
func HxBoosted(r *http.Request) bool {
	return r.Header.Get(HXBoostedHeader) == boolTrue
}

// HxCurrentURL returns the URL the user was on when the request was made.
// Useful for server-side decisions like highlighting the active navigation item.
func HxCurrentURL(r *http.Request) string {
	return r.Header.Get(HXCurrentURLHeader)
}

// HxHistoryRestoreRequest returns true when the user navigated back/forward and the page
// was not found in the local history cache. The server should return a full page response.
func HxHistoryRestoreRequest(r *http.Request) bool {
	return r.Header.Get(HXHistoryRestoreRequestHeader) == boolTrue
}

// HxPrompt returns the user's text input from a browser prompt, read from the HX-Prompt
// request header.
//
// hx-prompt is not in htmx 4 core, so the core build never sends this header. It was restored
// in beta5 as the hx-prompt extension (dist/ext/hx-prompt.js), which shows the prompt and sends
// HX-Prompt; pair this reader with the client [Wrapper.HxPrompt] setter and load that script. With
// neither in play this returns an empty string unless something else on your side sets HX-Prompt.
func HxPrompt(r *http.Request) string {
	return r.Header.Get(HXPromptHeader)
}

// HxTarget returns the target element the response will be swapped into.
// The value format is "tagName#id" (e.g. "div#result").
func HxTarget(r *http.Request) string {
	return r.Header.Get(HXTargetHeader)
}

// HxSource returns the element that triggered the request, sent in the HX-Source header
// with the value format "tagName#id" (e.g. "button#submit").
func HxSource(r *http.Request) string {
	return r.Header.Get(HXSourceHeader)
}

// HxRequestType returns the HX-Request-Type header, which htmx sets to "full" for a
// full-page request or "partial" for a partial swap. Use it to vary the response - for
// example returning a whole page layout for a full request and just a fragment for a
// partial one. Returns empty string for non-htmx requests.
func HxRequestType(r *http.Request) string {
	return r.Header.Get(HXRequestTypeHeader)
}

// HxRedirect performs a client-side redirect. For HTMX requests, it sets the HX-Redirect
// response header and writes a 200 status, then htmx navigates client-side. (htmx honours
// HX-Redirect at any status; 200 is used here so the htmx path returns an ordinary response
// carrying the header. The code argument applies to the standard HTTP redirect below.)
// For standard requests, it uses a standard HTTP redirect with the given code.
func HxRedirect(w http.ResponseWriter, r *http.Request, url string, code int) {
	if HxRequest(r) {
		w.Header().Set(HXRedirectHeader, url)
		w.WriteHeader(http.StatusOK)
	} else {
		http.Redirect(w, r, url, code)
	}
}

// HxPushURL pushes a new URL into the browser's history stack after the swap.
// Unlike client-side hx-push-url, this lets the server control the URL based on
// request processing (e.g. pushing a canonical URL after a form submission).
func HxPushURL(w http.ResponseWriter, url string) {
	w.Header().Set(HXPushURLHeader, url)
}

// HxLocation performs a client-side redirect without a full page reload.
// The value can be a plain URL string or a JSON object with path, source, event,
// target, swap, select, values and headers properties for fine-grained control.
func HxLocation(w http.ResponseWriter, url string) {
	w.Header().Set(HXLocationHeader, url)
}

// HxReplaceURL replaces the current URL in the browser's location bar without adding a history entry.
// Unlike HxPushURL, the user cannot navigate back to the previous URL.
func HxReplaceURL(w http.ResponseWriter, url string) {
	w.Header().Set(HXReplaceURLHeader, url)
}

// HxRefresh triggers a full page refresh on the client.
// Use sparingly - typically after operations that affect global state
// where a partial swap would leave the page inconsistent.
func HxRefresh(w http.ResponseWriter) {
	w.Header().Set(HXRefreshHeader, boolTrue)
}

// HxRetarget overrides the client-side hx-target, redirecting the swap to a different element.
// Useful when the server needs to change the swap target based on the response
// (e.g. showing an error in a different location than the success content).
func HxRetarget(w http.ResponseWriter, selector string) {
	w.Header().Set(HXRetargetHeader, selector)
}

// HxReswap overrides the client-side hx-swap strategy from the server.
// For example, the server can change "innerHTML" to "outerHTML" to replace the
// entire target element when returning an error state.
func HxReswap(w http.ResponseWriter, strategy swap.Strategy) {
	w.Header().Set(HXReswapHeader, string(strategy))
}

// HxReselect overrides the client-side hx-select, choosing a different fragment
// of the response to swap in. Useful when the server wants to override which part
// of a full page response is extracted.
func HxReselect(w http.ResponseWriter, selector string) {
	w.Header().Set(HXReselectHeader, selector)
}

// NewTrigger creates a new TriggerBuilder for accumulating HTMX trigger events.
// Call AddTrigger to queue events, then Write to send the response with the
// HX-Trigger header set.
func NewTrigger(w http.ResponseWriter) *TriggerBuilder {
	return &TriggerBuilder{
		w:        w,
		triggers: make([]triggerEvent, 0),
	}
}

// AddTrigger queues an event to fire when the response is received.
// If details is non-nil, it is included as a JSON object alongside the event name.
// Call Write to finalise and send.
func (tb *TriggerBuilder) AddTrigger(eventName string, details any) *TriggerBuilder {
	tb.triggers = append(tb.triggers, triggerEvent{Name: eventName, Details: details})

	return tb
}

// Write renders the node, sets the HX-Trigger header, and writes the response.
// Simple events (no details) are comma-separated; if any event has details,
// all events are marshaled into a single JSON object per the HTMX spec.
func (tb *TriggerBuilder) Write(n node.Node, code int) error {
	if len(tb.triggers) > 0 {
		headerValue, err := encodeTriggers(tb.triggers)
		if err != nil {
			return err
		}

		tb.w.Header().Set(HXTriggerHeader, headerValue)
	}

	tb.w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tb.w.WriteHeader(code)

	var buf bytes.Buffer
	n.RenderBuilder(&buf)

	if _, err := tb.w.Write(buf.Bytes()); err != nil {
		return fmt.Errorf("failed to write response: %w", err)
	}

	return nil
}

// encodeTriggers renders a set of trigger events into the HX-Trigger header value.
// If any event has details, all events are marshaled as a JSON object (simple events
// get a boolean true value); otherwise the names are joined with commas.
func encodeTriggers(events []triggerEvent) (string, error) {
	var simpleEvents []string

	detailedEvents := make(map[string]any)
	hasDetailedEvent := false

	for _, event := range events {
		if event.Details != nil {
			detailedEvents[event.Name] = event.Details
			hasDetailedEvent = true
		} else {
			simpleEvents = append(simpleEvents, event.Name)
		}
	}

	if !hasDetailedEvent {
		return strings.Join(simpleEvents, ","), nil
	}

	for _, se := range simpleEvents {
		detailedEvents[se] = true
	}

	jsonBytes, err := json.Marshal(detailedEvents)
	if err != nil {
		return "", fmt.Errorf("failed to marshal HTMX trigger details: %w", err)
	}

	return string(jsonBytes), nil
}

// Response renders the node and writes it as the response body with the given HTTP
// status code and a text/html content type. It is the plain full-response helper for
// returning a server-rendered fragment or page to htmx; reach for NewTrigger and
// TriggerBuilder.Write instead when the response also needs to fire client-side events
// through the HX-Trigger header.
func Response(w http.ResponseWriter, n node.Node, code int) error {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(code)

	var buf bytes.Buffer
	n.RenderBuilder(&buf)

	if _, err := w.Write(buf.Bytes()); err != nil {
		return fmt.Errorf("failed to write response: %w", err)
	}

	return nil
}
