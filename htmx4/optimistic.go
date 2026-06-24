package htmx

// Optimistic UI extension. Uses the hx-optimistic attribute. The extension is bundled in
// htmax.js; with the core htmx.js build, include the optimistic extension script yourself.

// HxOptimistic shows placeholder content the moment a request starts, then replaces it with
// the real response when it arrives. The selector points to an element whose inner HTML is
// cloned into the swap target optimistically; htmx removes the placeholder once the response
// is swapped in. Use it to make the UI feel instant for actions that usually succeed.
func (h *Wrapper) HxOptimistic(selector string) *Wrapper {
	h.element.SetAttribute("hx-optimistic", selector)

	return h
}
