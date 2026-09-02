package htmx

// Pending content extension. Uses the hx-pending attribute. The extension is bundled in
// htmax.js; with the core htmx.js build, include dist/ext/hx-pending.js yourself.

// HxPending shows placeholder content while a request is in flight. The selector points at
// an element, usually a <template>, whose children are cloned into a div that is placed at
// the swap target: over the target's children for an innerHTML swap, or beside the target
// for the other styles. Each request parameter is copied onto the div as a data-* attribute,
// so the placeholder can echo what was submitted. The div is removed when the response is
// swapped in or the request fails.
func (h *Wrapper) HxPending(selector string) *Wrapper {
	h.element.SetAttribute("hx-pending", selector)

	return h
}
