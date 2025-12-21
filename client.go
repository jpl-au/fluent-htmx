// Client-side HTMX attribute methods.
// This file contains all the hx-* attribute setters for use in the browser.

package htmx

// HxGet Sets the hx-get attribute for htmx GET requests
func (h *HtmxWrapper) HxGet(url string) *HtmxWrapper {
	h.node.SetAttribute("hx-get", url)
	return h
}

// HxPost Sets the hx-post attribute for htmx POST requests
func (h *HtmxWrapper) HxPost(url string) *HtmxWrapper {
	h.node.SetAttribute("hx-post", url)
	return h
}

// HxPut Sets the hx-put attribute for htmx PUT requests
func (h *HtmxWrapper) HxPut(url string) *HtmxWrapper {
	h.node.SetAttribute("hx-put", url)
	return h
}

// HxPatch Sets the hx-patch attribute for htmx PATCH requests
func (h *HtmxWrapper) HxPatch(url string) *HtmxWrapper {
	h.node.SetAttribute("hx-patch", url)
	return h
}

// HxDelete Sets the hx-delete attribute for htmx DELETE requests
func (h *HtmxWrapper) HxDelete(url string) *HtmxWrapper {
	h.node.SetAttribute("hx-delete", url)
	return h
}

// HxSwap Sets the hx-swap attribute to control how content is swapped
func (h *HtmxWrapper) HxSwap(strategy string) *HtmxWrapper {
	h.node.SetAttribute("hx-swap", strategy)
	return h
}

// HxTarget Sets the hx-target attribute to specify swap target
func (h *HtmxWrapper) HxTarget(target string) *HtmxWrapper {
	h.node.SetAttribute("hx-target", target)
	return h
}

// HxTrigger Sets the hx-trigger attribute to specify trigger events
func (h *HtmxWrapper) HxTrigger(trigger string) *HtmxWrapper {
	h.node.SetAttribute("hx-trigger", trigger)
	return h
}

// HxBoost Sets the hx-boost attribute to enable progressive enhancement
func (h *HtmxWrapper) HxBoost(enabled bool) *HtmxWrapper {
	value := "false"
	if enabled {
		value = "true"
	}
	h.node.SetAttribute("hx-boost", value)
	return h
}

// HxConfirm Sets the hx-confirm attribute for confirmation prompts
func (h *HtmxWrapper) HxConfirm(message string) *HtmxWrapper {
	h.node.SetAttribute("hx-confirm", message)
	return h
}

// HxVals Sets the hx-vals attribute for additional values
func (h *HtmxWrapper) HxVals(values string) *HtmxWrapper {
	h.node.SetAttribute("hx-vals", values)
	return h
}

// HxHeaders Sets the hx-headers attribute for additional headers
func (h *HtmxWrapper) HxHeaders(headers string) *HtmxWrapper {
	h.node.SetAttribute("hx-headers", headers)
	return h
}

// HxIndicator Sets the hx-indicator attribute for loading indicators
func (h *HtmxWrapper) HxIndicator(indicator string) *HtmxWrapper {
	h.node.SetAttribute("hx-indicator", indicator)
	return h
}

// HxPushURL Sets the hx-push-url attribute for URL management.
// Accepts a URL string, "true", or "false".
// Examples:
//   - HxPushURL("true") → hx-push-url="true" (pushes the fetched URL)
//   - HxPushURL("false") → hx-push-url="false" (disables URL pushing)
//   - HxPushURL("/custom/path") → hx-push-url="/custom/path" (pushes custom URL)
func (h *HtmxWrapper) HxPushURL(value string) *HtmxWrapper {
	h.node.SetAttribute("hx-push-url", value)
	return h
}

// HxExt Sets the hx-ext attribute for HTMX extensions
func (h *HtmxWrapper) HxExt(extensions string) *HtmxWrapper {
	h.node.SetAttribute("hx-ext", extensions)
	return h
}

// HxSelect Sets the hx-select attribute to choose which part of the response to swap
func (h *HtmxWrapper) HxSelect(selector string) *HtmxWrapper {
	h.node.SetAttribute("hx-select", selector)
	return h
}

// HxSelectOOB Sets the hx-select-oob attribute for out-of-band element selection
func (h *HtmxWrapper) HxSelectOOB(selector string) *HtmxWrapper {
	h.node.SetAttribute("hx-select-oob", selector)
	return h
}

// HxSwapOOB Sets the hx-swap-oob attribute for out-of-band swaps
func (h *HtmxWrapper) HxSwapOOB(value string) *HtmxWrapper {
	h.node.SetAttribute("hx-swap-oob", value)
	return h
}

// HxReplaceUrl Sets the hx-replace-url attribute to replace URL without adding to history
func (h *HtmxWrapper) HxReplaceUrl(url string) *HtmxWrapper {
	h.node.SetAttribute("hx-replace-url", url)
	return h
}

// HxParams Sets the hx-params attribute to filter request parameters
func (h *HtmxWrapper) HxParams(params string) *HtmxWrapper {
	h.node.SetAttribute("hx-params", params)
	return h
}

// HxInclude Sets the hx-include attribute to include additional form values
func (h *HtmxWrapper) HxInclude(selector string) *HtmxWrapper {
	h.node.SetAttribute("hx-include", selector)
	return h
}

// HxSync Sets the hx-sync attribute to synchronise multiple requests
func (h *HtmxWrapper) HxSync(strategy string) *HtmxWrapper {
	h.node.SetAttribute("hx-sync", strategy)
	return h
}

// HxPrompt Sets the hx-prompt attribute to prompt user for input
func (h *HtmxWrapper) HxPrompt(message string) *HtmxWrapper {
	h.node.SetAttribute("hx-prompt", message)
	return h
}

// HxEncoding Sets the hx-encoding attribute for request encoding type
func (h *HtmxWrapper) HxEncoding(encoding string) *HtmxWrapper {
	h.node.SetAttribute("hx-encoding", encoding)
	return h
}

// HxPreserve Sets the hx-preserve attribute to preserve elements during swaps
func (h *HtmxWrapper) HxPreserve(preserve bool) *HtmxWrapper {
	value := "false"
	if preserve {
		value = "true"
	}
	h.node.SetAttribute("hx-preserve", value)
	return h
}

// HxHistory Sets the hx-history attribute to prevent sensitive data caching.
// Accepts "false" to prevent the page from being cached in localStorage.
// Example: HxHistory("false")
func (h *HtmxWrapper) HxHistory(value string) *HtmxWrapper {
	h.node.SetAttribute("hx-history", value)
	return h
}

// HxHistoryElt Sets the hx-history-elt attribute to specify history snapshot element.
// This is a boolean flag attribute - its presence designates the element for history snapshots.
// No value is required.
func (h *HtmxWrapper) HxHistoryElt() *HtmxWrapper {
	h.node.SetAttribute("hx-history-elt", "true")
	return h
}

// HxDisable Sets the hx-disable attribute to prevent htmx processing.
// The value of the attribute is ignored - its presence disables htmx processing
// for the element and all its children. This cannot be reversed by any content beneath it.
func (h *HtmxWrapper) HxDisable() *HtmxWrapper {
	h.node.SetAttribute("hx-disable", "true")
	return h
}

// HxDisabledElt Sets the hx-disabled-elt attribute to disable elements during request
func (h *HtmxWrapper) HxDisabledElt(selector string) *HtmxWrapper {
	h.node.SetAttribute("hx-disabled-elt", selector)
	return h
}

// HxDisinherit Sets the hx-disinherit attribute to control attribute inheritance
func (h *HtmxWrapper) HxDisinherit(attributes string) *HtmxWrapper {
	h.node.SetAttribute("hx-disinherit", attributes)
	return h
}

// HxInherit Sets the hx-inherit attribute to explicitly specify inheritance
func (h *HtmxWrapper) HxInherit(attributes string) *HtmxWrapper {
	h.node.SetAttribute("hx-inherit", attributes)
	return h
}

// HxValidate Sets the hx-validate attribute to enable form validation
func (h *HtmxWrapper) HxValidate(validate bool) *HtmxWrapper {
	value := "false"
	if validate {
		value = "true"
	}
	h.node.SetAttribute("hx-validate", value)
	return h
}

// HxRequest Sets the hx-request attribute to configure request behaviour
func (h *HtmxWrapper) HxRequest(config string) *HtmxWrapper {
	h.node.SetAttribute("hx-request", config)
	return h
}

// HxVars Sets the hx-vars attribute to dynamically compute request values (deprecated)
func (h *HtmxWrapper) HxVars(variables string) *HtmxWrapper {
	h.node.SetAttribute("hx-vars", variables)
	return h
}

// HxOn sets the hx-on::event attribute for inline event handling
// Example: HxOn("after-swap", "console.log('swapped')")
// Results in: hx-on::after-swap="console.log('swapped')"
func (h *HtmxWrapper) HxOn(event string, handler string) *HtmxWrapper {
	h.node.SetAttribute("hx-on::"+event, handler)
	return h
}
