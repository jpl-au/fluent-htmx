// Package htmx provides a set of helpers for working with HTMX in Go applications.
// It simplifies handling HTMX-specific request headers and constructing HTMX-specific responses.
package htmx

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strings"
)

// triggerEvent represents a single HTMX trigger event.
type triggerEvent struct {
	Name    string
	Details any
}

// TriggerBuilder is a helper for building HTMX trigger responses.
// It allows for chaining of trigger additions and finalizes the headers.
type TriggerBuilder struct {
	w                   http.ResponseWriter
	triggers            []triggerEvent
	triggersAfterSettle []triggerEvent
	triggersAfterSwap   []triggerEvent
}

// HxRequest checks if the request was initiated by HTMX by checking for the HX-Request header.
func HxRequest(r *http.Request) bool {
	slog.Debug("Checking HTMX request", "HX-Request", r.Header.Get(HXRequestHeader))
	return strings.ToLower(r.Header.Get(HXRequestHeader)) == "true"
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

// HxBoosted checks if the request was made from an element with hx-boost="true".
// It reads the HX-Boosted header.
func HxBoosted(r *http.Request) bool {
	return r.Header.Get(HXBoostedHeader) == "true"
}

// HxCurrentURL returns the URL of the page that the request was sent from.
// It reads the HX-Current-URL header.
func HxCurrentURL(r *http.Request) string {
	return r.Header.Get(HXCurrentURLHeader)
}

// HxHistoryRestoreRequest checks if the request is for history restoration after a miss in the local history cache.
// It returns true if the HX-History-Restore-Request header is set to "true".
func HxHistoryRestoreRequest(r *http.Request) bool {
	return r.Header.Get(HXHistoryRestoreRequest) == "true"
}

// HxPrompt returns the user's response to a prompt shown via hx-prompt.
// It reads the HX-Prompt header.
func HxPrompt(r *http.Request) string {
	return r.Header.Get(HXPromptHeader)
}

// HxTarget returns the ID of the target element for the request.
// It reads the HX-Target header.
func HxTarget(r *http.Request) string {
	return r.Header.Get(HXTargetHeader)
}

// HxTriggerName returns the name attribute of the element that triggered the request.
// It reads the HX-Trigger-Name header.
func HxTriggerName(r *http.Request) string {
	return r.Header.Get(HXTriggerNameHeader)
}

// HxTrigger returns the ID of the element that triggered the request.
// It reads the HX-Trigger header.
func HxTrigger(r *http.Request) string {
	return r.Header.Get(HXTriggerHeader)
}

// HxRedirect performs a client-side redirect. For HTMX requests, it sets the HX-Redirect
// response header. For standard requests, it uses a standard HTTP 3xx redirect.
func HxRedirect(w http.ResponseWriter, r *http.Request, url string, code int) {
	if HxRequest(r) {
		w.Header().Set(HXRedirectHeader, url)
		w.WriteHeader(http.StatusOK)
	} else {
		http.Redirect(w, r, url, code)
	}
}

// HxPushURL pushes a new URL into the browser's history stack.
// It sets the HX-Push-Url response header.
func HxPushURL(w http.ResponseWriter, url string) {
	w.Header().Set(HXPushURLHeader, url)
}

// HxLocation performs a client-side redirect without a full page reload.
// It sets the HX-Location response header. The url parameter can be a string or a JSON object
// with path, source, event, handler, target, swap, and values properties.
func HxLocation(w http.ResponseWriter, url string) {
	w.Header().Set(HXLocationHeader, url)
}

// HxReplaceURL replaces the current URL in the browser's location bar.
// It sets the HX-Replace-Url response header.
func HxReplaceURL(w http.ResponseWriter, url string) {
	w.Header().Set(HXReplaceURLHeader, url)
}

// HxRefresh sends a response that triggers a client-side refresh of the page.
// It sets the HX-Refresh response header to "true".
func HxRefresh(w http.ResponseWriter) {
	w.Header().Set(HXRefreshHeader, "true")
}

// HxRetarget changes the target of the response to a different element on the page.
// It sets the HX-Retarget response header to the provided CSS selector.
func HxRetarget(w http.ResponseWriter, selector string) {
	w.Header().Set(HXRetargetHeader, selector)
}

// HxReswap changes the swapping method for the response.
// It sets the HX-Reswap response header to the provided method (e.g., "innerHTML", "beforeend").
func HxReswap(w http.ResponseWriter, method string) {
	w.Header().Set(HXReswapHeader, method)
}

// HxReselect allows you to choose which part of the response is used to be swapped in.
// It sets the HX-Reselect response header to the provided CSS selector.
// This overrides an existing hx-select on the triggering element.
func HxReselect(w http.ResponseWriter, selector string) {
	w.Header().Set(HXReselectHeader, selector)
}

// NewTrigger creates a new TriggerBuilder instance.
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

// Write formats and sets the accumulated trigger headers on the http.ResponseWriter.
// It then writes the provided content with the given status code.
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

	tb.w.WriteHeader(code)
	tb.w.Write([]byte(content))
	return nil
}

// AddTrigger adds an event to the HX-Trigger response header.
// This function can be called multiple times on the same response to trigger multiple events.
// If 'details' is not nil, it will be marshaled to a JSON string and included with the event.
func (tb *TriggerBuilder) AddTrigger(eventName string, details interface{}) *TriggerBuilder {
	return tb.addTrigger(HXTriggerHeader, eventName, details)
}

// AddTriggerAfterSettle adds an event to the HX-Trigger-After-Settle response header.
// This function can be called multiple times. Events are triggered after the HTMX settling phase.
// If 'details' is not nil, it will be marshaled to a JSON string.
func (tb *TriggerBuilder) AddTriggerAfterSettle(eventName string, details interface{}) *TriggerBuilder {
	return tb.addTrigger(HXTriggerAfterSettleHeader, eventName, details)
}

// AddTriggerAfterSwap adds an event to the HX-Trigger-After-Swap response header.
// This function can be called multiple times. Events are triggered after the content swap phase.
// If 'details' is not nil, it will be marshaled to a JSON string.
func (tb *TriggerBuilder) AddTriggerAfterSwap(eventName string, details interface{}) *TriggerBuilder {
	return tb.addTrigger(HXTriggerAfterSwapHeader, eventName, details)
}

// Response is a convenience function to write a simple HTMX response with a given status code.
func Response(w http.ResponseWriter, content string, code int) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(code)
	w.Write([]byte(content))
}