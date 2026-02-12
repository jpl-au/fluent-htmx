package htmx

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// triggerEvent represents a single HTMX trigger event.
type triggerEvent struct {
	Name    string
	Details any
}

// TriggerBuilder accumulates HTMX trigger events across the three timing phases
// (immediate, after-swap, after-settle) and writes them as response headers.
type TriggerBuilder struct {
	w                   http.ResponseWriter
	triggers            []triggerEvent
	triggersAfterSettle []triggerEvent
	triggersAfterSwap   []triggerEvent
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
// Boosted requests behave like standard navigation but use AJAX — the server may want
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

// HxPrompt returns the user's text input from a browser prompt shown via hx-prompt.
// Returns empty string if no prompt was shown.
func HxPrompt(r *http.Request) string {
	return r.Header.Get(HXPromptHeader)
}

// HxTarget returns the ID of the element the response will be swapped into.
// The server can use this to vary the response — e.g. return different content
// depending on which part of the page is being updated.
func HxTarget(r *http.Request) string {
	return r.Header.Get(HXTargetHeader)
}

// HxTriggerName returns the name attribute of the element that triggered the request.
// Useful for distinguishing which form or input initiated the request when multiple
// elements target the same endpoint.
func HxTriggerName(r *http.Request) string {
	return r.Header.Get(HXTriggerNameHeader)
}

// HxTrigger returns the ID of the element that triggered the request.
// Together with HxTriggerName, this identifies exactly which element initiated the request.
func HxTrigger(r *http.Request) string {
	return r.Header.Get(HXTriggerHeader)
}

// HxRedirect performs a client-side redirect. For HTMX requests, it sets the HX-Redirect
// response header with a 200 status — HTMX processes redirects client-side so the response
// must be 200 for the header to be read. For standard requests, it uses a standard HTTP redirect.
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
// handler, target, swap, and values properties for fine-grained control.
func HxLocation(w http.ResponseWriter, url string) {
	w.Header().Set(HXLocationHeader, url)
}

// HxReplaceURL replaces the current URL in the browser's location bar without adding a history entry.
// Unlike HxPushURL, the user cannot navigate back to the previous URL.
func HxReplaceURL(w http.ResponseWriter, url string) {
	w.Header().Set(HXReplaceURLHeader, url)
}

// HxRefresh triggers a full page refresh on the client.
// Use sparingly — typically after operations that affect global state
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
func HxReswap(w http.ResponseWriter, method string) {
	w.Header().Set(HXReswapHeader, method)
}

// HxReselect overrides the client-side hx-select, choosing a different fragment
// of the response to swap in. Useful when the server wants to override which part
// of a full page response is extracted.
func HxReselect(w http.ResponseWriter, selector string) {
	w.Header().Set(HXReselectHeader, selector)
}

// NewTrigger creates a new TriggerBuilder for accumulating HTMX trigger events.
// Call AddTrigger, AddTriggerAfterSwap, or AddTriggerAfterSettle to queue events,
// then call Write to send the response with all trigger headers set.
func NewTrigger(w http.ResponseWriter) *TriggerBuilder {
	return &TriggerBuilder{
		w:                   w,
		triggers:            make([]triggerEvent, 0),
		triggersAfterSettle: make([]triggerEvent, 0),
		triggersAfterSwap:   make([]triggerEvent, 0),
	}
}

// addTrigger is an internal helper to reduce repetition in the public trigger methods.
func (tb *TriggerBuilder) addTrigger(header string, eventName string, details interface{}) *TriggerBuilder {
	switch header {
	case HXTriggerHeader:
		tb.triggers = append(tb.triggers, triggerEvent{Name: eventName, Details: details})
	case HXTriggerAfterSettleHeader:
		tb.triggersAfterSettle = append(tb.triggersAfterSettle, triggerEvent{Name: eventName, Details: details})
	case HXTriggerAfterSwapHeader:
		tb.triggersAfterSwap = append(tb.triggersAfterSwap, triggerEvent{Name: eventName, Details: details})
	}

	return tb
}

// Write sets the accumulated trigger headers and writes the response.
// Simple events (no details) are comma-separated; if any event has details,
// all events are marshaled into a single JSON object per the HTMX spec.
func (tb *TriggerBuilder) Write(content string, code int) error {
	triggerHeaders := map[string][]triggerEvent{
		HXTriggerHeader:            tb.triggers,
		HXTriggerAfterSettleHeader: tb.triggersAfterSettle,
		HXTriggerAfterSwapHeader:   tb.triggersAfterSwap,
	}

	for headerKey, events := range triggerHeaders {
		if len(events) == 0 {
			continue
		}

		var simpleEvents []string

		detailedEvents := make(map[string]interface{})
		hasDetailedEvent := false

		for _, event := range events {
			if event.Details != nil {
				detailedEvents[event.Name] = event.Details
				hasDetailedEvent = true
			} else {
				simpleEvents = append(simpleEvents, event.Name)
			}
		}

		var finalHeaderValue string

		if hasDetailedEvent {
			// If any event has details, all events must be marshaled into a JSON object.
			// Simple events are added with a boolean true value.
			for _, se := range simpleEvents {
				detailedEvents[se] = true
			}

			jsonBytes, err := json.Marshal(detailedEvents)
			if err != nil {
				return fmt.Errorf("failed to marshal HTMX trigger details for %s: %w", headerKey, err)
			}

			finalHeaderValue = string(jsonBytes)
		} else {
			// Only simple events, use comma-separated string.
			finalHeaderValue = strings.Join(simpleEvents, ",")
		}

		tb.w.Header().Set(headerKey, finalHeaderValue)
	}

	tb.w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tb.w.WriteHeader(code)

	if _, err := tb.w.Write([]byte(content)); err != nil {
		return fmt.Errorf("failed to write response: %w", err)
	}

	return nil
}

// AddTrigger queues an event to fire immediately when the response is received.
// If details is non-nil, it is included as a JSON object alongside the event name.
// Call Write to finalise and send.
func (tb *TriggerBuilder) AddTrigger(eventName string, details interface{}) *TriggerBuilder {
	return tb.addTrigger(HXTriggerHeader, eventName, details)
}

// AddTriggerAfterSettle queues an event to fire after the DOM has settled.
// Settling occurs after new content attributes have been applied — use this
// for operations that depend on final attribute values.
func (tb *TriggerBuilder) AddTriggerAfterSettle(eventName string, details interface{}) *TriggerBuilder {
	return tb.addTrigger(HXTriggerAfterSettleHeader, eventName, details)
}

// AddTriggerAfterSwap queues an event to fire after the content swap completes.
// Use this for operations that depend on the new content being in the DOM
// (e.g. initialising JavaScript components on swapped elements).
func (tb *TriggerBuilder) AddTriggerAfterSwap(eventName string, details interface{}) *TriggerBuilder {
	return tb.addTrigger(HXTriggerAfterSwapHeader, eventName, details)
}

// Response writes a simple HTML response with the given status code.
func Response(w http.ResponseWriter, content string, code int) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(code)
	_, _ = w.Write([]byte(content))
}
