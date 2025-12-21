package htmx

import "fmt"

// HxTargetError sets the hx-target-error attribute to handle both 4xx and 5xx error responses
func (h *HtmxWrapper) HxTargetError(selector string) *HtmxWrapper {
	h.node.SetAttribute("hx-target-error", selector)
	return h
}

// HxTargetCode sets the hx-target-[CODE] attribute for a specific HTTP response code
// Example: HxTargetCode(404, "#not-found") sets hx-target-404="#not-found"
func (h *HtmxWrapper) HxTargetCode(code int, selector string) *HtmxWrapper {
	attr := fmt.Sprintf("hx-target-%d", code)
	h.node.SetAttribute(attr, selector)
	return h
}

// HxTargetCodePattern sets the hx-target-[PATTERN] attribute for wildcard response codes
// Example: HxTargetCodePattern("4*", "#client-error") matches all 4xx codes
// Example: HxTargetCodePattern("5*", "#server-error") matches all 5xx codes
func (h *HtmxWrapper) HxTargetCodePattern(pattern string, selector string) *HtmxWrapper {
	attr := fmt.Sprintf("hx-target-%s", pattern)
	h.node.SetAttribute(attr, selector)
	return h
}
