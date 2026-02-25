// Client-side HTMX attribute methods.
// This file contains all the hx-* attribute setters for use in the browser.

package htmx

// HxGet issues an AJAX GET request to the given URL and swaps the response into the DOM.
func (h *Wrapper) HxGet(url string) *Wrapper {
	h.element.SetAttribute("hx-get", url)

	return h
}

// HxPost issues an AJAX POST request to the given URL and swaps the response into the DOM.
func (h *Wrapper) HxPost(url string) *Wrapper {
	h.element.SetAttribute("hx-post", url)

	return h
}

// HxPut issues an AJAX PUT request to the given URL and swaps the response into the DOM.
func (h *Wrapper) HxPut(url string) *Wrapper {
	h.element.SetAttribute("hx-put", url)

	return h
}

// HxPatch issues an AJAX PATCH request to the given URL and swaps the response into the DOM.
func (h *Wrapper) HxPatch(url string) *Wrapper {
	h.element.SetAttribute("hx-patch", url)

	return h
}

// HxDelete issues an AJAX DELETE request to the given URL and swaps the response into the DOM.
func (h *Wrapper) HxDelete(url string) *Wrapper {
	h.element.SetAttribute("hx-delete", url)

	return h
}

// HxSwap controls how the response content is swapped into the DOM.
// Use the predefined Swap* constants or CustomSwap() for strategies with modifiers.
func (h *Wrapper) HxSwap(strategy SwapStrategy) *Wrapper {
	h.element.SetAttribute("hx-swap", string(strategy))

	return h
}

// HxTarget specifies a CSS selector for the element that will receive the swapped content.
// Without this, the element that triggers the request is the swap target.
func (h *Wrapper) HxTarget(target string) *Wrapper {
	h.element.SetAttribute("hx-target", target)

	return h
}

// HxTrigger specifies which DOM events cause the element to issue a request.
// Supports standard DOM events, HTMX-specific events, and modifiers.
// Example: "keyup changed delay:500ms" triggers on keyup after 500ms of inactivity.
func (h *Wrapper) HxTrigger(trigger string) *Wrapper {
	h.element.SetAttribute("hx-trigger", trigger)

	return h
}

// HxBoost converts standard links and forms into AJAX requests,
// swapping the response into the page body without a full page reload.
func (h *Wrapper) HxBoost(enabled bool) *Wrapper {
	value := boolFalse
	if enabled {
		value = boolTrue
	}

	h.element.SetAttribute("hx-boost", value)

	return h
}

// HxConfirm shows a browser confirmation dialog before issuing the request.
// The request is only sent if the user confirms.
func (h *Wrapper) HxConfirm(message string) *Wrapper {
	h.element.SetAttribute("hx-confirm", message)

	return h
}

// HxVals adds extra JSON-encoded values to the request parameters.
// Example: `{"key": "value"}`.
func (h *Wrapper) HxVals(values string) *Wrapper {
	h.element.SetAttribute("hx-vals", values)

	return h
}

// HxHeaders adds extra JSON-encoded headers to the AJAX request.
// Example: `{"X-Custom-Header": "value"}`.
func (h *Wrapper) HxHeaders(headers string) *Wrapper {
	h.element.SetAttribute("hx-headers", headers)

	return h
}

// HxIndicator specifies a CSS selector for an element to show while the request is in flight.
// The targeted element receives the htmx-request class, which can be styled to show a spinner.
func (h *Wrapper) HxIndicator(indicator string) *Wrapper {
	h.element.SetAttribute("hx-indicator", indicator)

	return h
}

// HxPushURL pushes a URL into the browser history after the request completes.
// Accepts "true" to push the fetched URL, "false" to disable, or a custom URL string.
// Examples:
//   - HxPushURL("true") → pushes the fetched URL
//   - HxPushURL("false") → disables URL pushing
//   - HxPushURL("/custom/path") → pushes a custom URL
func (h *Wrapper) HxPushURL(value string) *Wrapper {
	h.element.SetAttribute("hx-push-url", value)

	return h
}

// HxExt enables one or more HTMX extensions on the element.
// Multiple extensions are comma-separated. Extensions must be loaded via script tag first.
// Example: "ws" or "sse,preload".
func (h *Wrapper) HxExt(extensions string) *Wrapper {
	h.element.SetAttribute("hx-ext", extensions)

	return h
}

// HxSelect picks a CSS selector from the response HTML to swap in,
// discarding the rest of the response. Useful when the server returns a full page
// but only a fragment is needed.
func (h *Wrapper) HxSelect(selector string) *Wrapper {
	h.element.SetAttribute("hx-select", selector)

	return h
}

// HxSelectOOB selects content from the response for out-of-band swaps.
// These elements are swapped into matching targets elsewhere in the DOM,
// independently of the primary swap target.
func (h *Wrapper) HxSelectOOB(selector string) *Wrapper {
	h.element.SetAttribute("hx-select-oob", selector)

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
func (h *Wrapper) HxReplaceURL(url string) *Wrapper {
	h.element.SetAttribute("hx-replace-url", url)

	return h
}

// HxParams filters which request parameters are submitted.
// Accepts "*" (all), "none", or a comma-separated list of parameter names.
// Prefix with "not " to exclude specific parameters: "not name,email".
func (h *Wrapper) HxParams(params string) *Wrapper {
	h.element.SetAttribute("hx-params", params)

	return h
}

// HxInclude includes values from other elements in the request using a CSS selector.
// Useful for submitting inputs that are outside the triggering element's form.
func (h *Wrapper) HxInclude(selector string) *Wrapper {
	h.element.SetAttribute("hx-include", selector)

	return h
}

// HxSync coordinates requests between this element and another element matched by the selector.
// Prevents race conditions when multiple elements can trigger overlapping requests.
// Strategies: "drop", "abort", "replace", "queue first", "queue last", "queue all".
func (h *Wrapper) HxSync(strategy string) *Wrapper {
	h.element.SetAttribute("hx-sync", strategy)

	return h
}

// HxPrompt shows a browser prompt dialog before issuing the request.
// The user's input is sent to the server via the HX-Prompt request header.
func (h *Wrapper) HxPrompt(message string) *Wrapper {
	h.element.SetAttribute("hx-prompt", message)

	return h
}

// HxEncoding changes the request encoding type.
// The only non-default value is "multipart/form-data" for file uploads.
func (h *Wrapper) HxEncoding(encoding string) *Wrapper {
	h.element.SetAttribute("hx-encoding", encoding)

	return h
}

// HxPreserve keeps the element unchanged during swaps by matching its ID.
// Useful for persistent elements like video players or iframes that should
// survive content updates around them.
func (h *Wrapper) HxPreserve(preserve bool) *Wrapper {
	value := boolFalse
	if preserve {
		value = boolTrue
	}

	h.element.SetAttribute("hx-preserve", value)

	return h
}

// HxHistory prevents the page from being saved to the browser's localStorage history cache.
// Use HxHistory("false") on pages that contain sensitive data to avoid caching it.
func (h *Wrapper) HxHistory(value string) *Wrapper {
	h.element.SetAttribute("hx-history", value)

	return h
}

// HxHistoryElt designates this element as the snapshot source for history navigation.
// By default HTMX snapshots the entire body; this narrows it to a specific element.
func (h *Wrapper) HxHistoryElt() *Wrapper {
	h.element.SetAttribute("hx-history-elt", boolTrue)

	return h
}

// HxDisable prevents HTMX from processing this element and all of its children.
// The attribute's presence alone is sufficient; its value is ignored.
// This cannot be overridden by any descendant content.
func (h *Wrapper) HxDisable() *Wrapper {
	h.element.SetAttribute("hx-disable", boolTrue)

	return h
}

// HxDisabledElt adds the "disabled" attribute to the matched elements while a request is in flight.
// Useful for preventing duplicate form submissions by disabling the submit button.
func (h *Wrapper) HxDisabledElt(selector string) *Wrapper {
	h.element.SetAttribute("hx-disabled-elt", selector)

	return h
}

// HxDisinherit prevents child elements from inheriting specific HTMX attributes.
// Accepts a space-separated list of attribute names or "*" to disable all inheritance.
func (h *Wrapper) HxDisinherit(attributes string) *Wrapper {
	h.element.SetAttribute("hx-disinherit", attributes)

	return h
}

// HxInherit re-enables inheritance for specific attributes that were disabled by a parent's HxDisinherit("*").
func (h *Wrapper) HxInherit(attributes string) *Wrapper {
	h.element.SetAttribute("hx-inherit", attributes)

	return h
}

// HxValidate forces the browser's native form validation before issuing the request.
// The request is blocked if any included input fails its validation constraints.
func (h *Wrapper) HxValidate(validate bool) *Wrapper {
	value := boolFalse
	if validate {
		value = boolTrue
	}

	h.element.SetAttribute("hx-validate", value)

	return h
}

// HxRequest overrides HTMX request behaviour with a JSON configuration string.
// Supported keys: "timeout" (ms), "credentials" (bool), "noHeaders" (bool).
// Example: `{"timeout": 5000}`.
func (h *Wrapper) HxRequest(config string) *Wrapper {
	h.element.SetAttribute("hx-request", config)

	return h
}

// HxVars is deprecated in favour of HxVals. It evaluates JavaScript expressions
// to compute request values, which requires allowEval to be enabled.
func (h *Wrapper) HxVars(variables string) *Wrapper {
	h.element.SetAttribute("hx-vars", variables)

	return h
}

// HxOn attaches an inline event handler directly to the element (locality of behaviour).
// Uses the hx-on::event attribute syntax so the handler lives with the element it controls,
// rather than in a separate script tag.
// Example: HxOn("after-swap", "console.log('swapped')") → hx-on::after-swap="console.log('swapped')".
func (h *Wrapper) HxOn(event string, handler string) *Wrapper {
	h.element.SetAttribute("hx-on::"+event, handler)

	return h
}
