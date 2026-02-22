package htmx

import "fmt"

// HxTargetError sets the hx-target-error attribute to handle both 4xx and 5xx error responses.
func (h *Wrapper) HxTargetError(selector string) *Wrapper {
	h.element.SetAttribute("hx-target-error", selector)

	return h
}

// HxTargetCode sets the hx-target-[CODE] attribute for a specific HTTP response code.
// Example: HxTargetCode(404, "#not-found") sets hx-target-404="#not-found".
func (h *Wrapper) HxTargetCode(code int, selector string) *Wrapper {
	attr := fmt.Sprintf("hx-target-%d", code)
	h.element.SetAttribute(attr, selector)

	return h
}

// HxTargetCodePattern sets the hx-target-[PATTERN] attribute for wildcard response codes.
// Example: HxTargetCodePattern("4*", "#client-error") matches all 4xx codes.
// Example: HxTargetCodePattern("5*", "#server-error") matches all 5xx codes.
func (h *Wrapper) HxTargetCodePattern(pattern string, selector string) *Wrapper {
	attr := "hx-target-" + pattern
	h.element.SetAttribute(attr, selector)

	return h
}
