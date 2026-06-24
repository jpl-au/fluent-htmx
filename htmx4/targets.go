package htmx

// Multi-target extension. Uses the hx-targets attribute. The extension is bundled in htmax.js;
// with the core htmx.js build, include the targets extension script yourself.

// HxTargets swaps the response into every element matching the selector, rather than a single
// hx-target. The response fragment is cloned once per match. The selector is resolved
// relative to the requesting element. Useful for updating several places on the page from one
// request without out-of-band swaps.
func (h *Wrapper) HxTargets(selector string) *Wrapper {
	h.element.SetAttribute("hx-targets", selector)

	return h
}
