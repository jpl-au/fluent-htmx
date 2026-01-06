// Client-side HTMX attribute methods.
// This file contains all the hx-* attribute setters for use in the browser.

package htmx

// HxGet Sets the hx-get attribute for htmx GET requests.
func (h *Wrapper) HxGet(url string) *Wrapper {
	h.node.SetAttribute("hx-get", url)

	return h
}

// HxPost Sets the hx-post attribute for htmx POST requests.
func (h *Wrapper) HxPost(url string) *Wrapper {
	h.node.SetAttribute("hx-post", url)

	return h
}

// HxPut Sets the hx-put attribute for htmx PUT requests.
func (h *Wrapper) HxPut(url string) *Wrapper {
	h.node.SetAttribute("hx-put", url)

	return h
}

// HxPatch Sets the hx-patch attribute for htmx PATCH requests.
func (h *Wrapper) HxPatch(url string) *Wrapper {
	h.node.SetAttribute("hx-patch", url)

	return h
}

// HxDelete Sets the hx-delete attribute for htmx DELETE requests.
func (h *Wrapper) HxDelete(url string) *Wrapper {
	h.node.SetAttribute("hx-delete", url)

	return h
}

// HxSwap Sets the hx-swap attribute to control how content is swapped.
func (h *Wrapper) HxSwap(strategy string) *Wrapper {
	h.node.SetAttribute("hx-swap", strategy)

	return h
}

// HxTarget Sets the hx-target attribute to specify swap target.
func (h *Wrapper) HxTarget(target string) *Wrapper {
	h.node.SetAttribute("hx-target", target)

	return h
}

// HxTrigger Sets the hx-trigger attribute to specify trigger events.
func (h *Wrapper) HxTrigger(trigger string) *Wrapper {
	h.node.SetAttribute("hx-trigger", trigger)

	return h
}

// HxBoost Sets the hx-boost attribute to enable progressive enhancement.
func (h *Wrapper) HxBoost(enabled bool) *Wrapper {
	value := boolFalse
	if enabled {
		value = boolTrue
	}

	h.node.SetAttribute("hx-boost", value)

	return h
}

// HxConfirm Sets the hx-confirm attribute for confirmation prompts.
func (h *Wrapper) HxConfirm(message string) *Wrapper {
	h.node.SetAttribute("hx-confirm", message)

	return h
}

// HxVals Sets the hx-vals attribute for additional values.
func (h *Wrapper) HxVals(values string) *Wrapper {
	h.node.SetAttribute("hx-vals", values)

	return h
}

// HxHeaders Sets the hx-headers attribute for additional headers.
func (h *Wrapper) HxHeaders(headers string) *Wrapper {
	h.node.SetAttribute("hx-headers", headers)

	return h
}

// HxIndicator Sets the hx-indicator attribute for loading indicators.
func (h *Wrapper) HxIndicator(indicator string) *Wrapper {
	h.node.SetAttribute("hx-indicator", indicator)

	return h
}

// HxPushURL Sets the hx-push-url attribute for URL management.
// Accepts a URL string, "true", or "false".
// Examples:
//   - HxPushURL("true") → hx-push-url="true" (pushes the fetched URL)
//   - HxPushURL("false") → hx-push-url="false" (disables URL pushing)
//   - HxPushURL("/custom/path") → hx-push-url="/custom/path" (pushes custom URL)
func (h *Wrapper) HxPushURL(value string) *Wrapper {
	h.node.SetAttribute("hx-push-url", value)

	return h
}

// HxExt Sets the hx-ext attribute for HTMX extensions.
func (h *Wrapper) HxExt(extensions string) *Wrapper {
	h.node.SetAttribute("hx-ext", extensions)

	return h
}

// HxSelect Sets the hx-select attribute to choose which part of the response to swap.
func (h *Wrapper) HxSelect(selector string) *Wrapper {
	h.node.SetAttribute("hx-select", selector)

	return h
}

// HxSelectOOB Sets the hx-select-oob attribute for out-of-band element selection.
func (h *Wrapper) HxSelectOOB(selector string) *Wrapper {
	h.node.SetAttribute("hx-select-oob", selector)

	return h
}

// HxSwapOOB Sets the hx-swap-oob attribute for out-of-band swaps.
func (h *Wrapper) HxSwapOOB(value string) *Wrapper {
	h.node.SetAttribute("hx-swap-oob", value)

	return h
}

// HxReplaceURL Sets the hx-replace-url attribute to replace URL without adding to history.
func (h *Wrapper) HxReplaceURL(url string) *Wrapper {
	h.node.SetAttribute("hx-replace-url", url)

	return h
}

// HxParams Sets the hx-params attribute to filter request parameters.
func (h *Wrapper) HxParams(params string) *Wrapper {
	h.node.SetAttribute("hx-params", params)

	return h
}

// HxInclude Sets the hx-include attribute to include additional form values.
func (h *Wrapper) HxInclude(selector string) *Wrapper {
	h.node.SetAttribute("hx-include", selector)

	return h
}

// HxSync Sets the hx-sync attribute to synchronise multiple requests.
func (h *Wrapper) HxSync(strategy string) *Wrapper {
	h.node.SetAttribute("hx-sync", strategy)

	return h
}

// HxPrompt Sets the hx-prompt attribute to prompt user for input.
func (h *Wrapper) HxPrompt(message string) *Wrapper {
	h.node.SetAttribute("hx-prompt", message)

	return h
}

// HxEncoding Sets the hx-encoding attribute for request encoding type.
func (h *Wrapper) HxEncoding(encoding string) *Wrapper {
	h.node.SetAttribute("hx-encoding", encoding)

	return h
}

// HxPreserve Sets the hx-preserve attribute to preserve elements during swaps.
func (h *Wrapper) HxPreserve(preserve bool) *Wrapper {
	value := boolFalse
	if preserve {
		value = boolTrue
	}

	h.node.SetAttribute("hx-preserve", value)

	return h
}

// HxHistory Sets the hx-history attribute to prevent sensitive data caching.
// Accepts "false" to prevent the page from being cached in localStorage.
// Example: HxHistory("false").
func (h *Wrapper) HxHistory(value string) *Wrapper {
	h.node.SetAttribute("hx-history", value)

	return h
}

// HxHistoryElt Sets the hx-history-elt attribute to specify history snapshot element.
// This is a boolean flag attribute - its presence designates the element for history snapshots.
// No value is required.
func (h *Wrapper) HxHistoryElt() *Wrapper {
	h.node.SetAttribute("hx-history-elt", boolTrue)

	return h
}

// HxDisable Sets the hx-disable attribute to prevent htmx processing.
// The value of the attribute is ignored - its presence disables htmx processing
// for the element and all its children. This cannot be reversed by any content beneath it.
func (h *Wrapper) HxDisable() *Wrapper {
	h.node.SetAttribute("hx-disable", boolTrue)

	return h
}

// HxDisabledElt Sets the hx-disabled-elt attribute to disable elements during request.
func (h *Wrapper) HxDisabledElt(selector string) *Wrapper {
	h.node.SetAttribute("hx-disabled-elt", selector)

	return h
}

// HxDisinherit Sets the hx-disinherit attribute to control attribute inheritance.
func (h *Wrapper) HxDisinherit(attributes string) *Wrapper {
	h.node.SetAttribute("hx-disinherit", attributes)

	return h
}

// HxInherit Sets the hx-inherit attribute to explicitly specify inheritance.
func (h *Wrapper) HxInherit(attributes string) *Wrapper {
	h.node.SetAttribute("hx-inherit", attributes)

	return h
}

// HxValidate Sets the hx-validate attribute to enable form validation.
func (h *Wrapper) HxValidate(validate bool) *Wrapper {
	value := boolFalse
	if validate {
		value = boolTrue
	}

	h.node.SetAttribute("hx-validate", value)

	return h
}

// HxRequest Sets the hx-request attribute to configure request behaviour.
func (h *Wrapper) HxRequest(config string) *Wrapper {
	h.node.SetAttribute("hx-request", config)

	return h
}

// HxVars Sets the hx-vars attribute to dynamically compute request values (deprecated).
func (h *Wrapper) HxVars(variables string) *Wrapper {
	h.node.SetAttribute("hx-vars", variables)

	return h
}

// HxOn sets the hx-on::event attribute for inline event handling.
// Example: HxOn("after-swap", "console.log('swapped')")
// Results in: hx-on::after-swap="console.log('swapped')".
func (h *Wrapper) HxOn(event string, handler string) *Wrapper {
	h.node.SetAttribute("hx-on::"+event, handler)

	return h
}
