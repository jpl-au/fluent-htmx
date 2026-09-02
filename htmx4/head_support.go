package htmx

// HxHead sets the hx-head attribute, which controls <head> merge behaviour. Place it on the
// response <head> tag with "merge" (keep matching elements, add new ones, remove missing
// ones) or "append" (add all to the existing head); place "re-eval" on an individual head
// element to re-add it on every request.
//
// hx-head is an htmx 4 extension shipped as a separate script (dist/ext/hx-head.js). It is
// not one of the extensions bundled in this htmax.js build (sse, ws, preload,
// browser-indicator, download, hx-pending, hx-targets, hx-live). Without the hx-head
// extension loaded, core htmx only swaps in the response title and ignores this attribute.
func (h *Wrapper) HxHead(mode string) *Wrapper {
	h.element.SetAttribute("hx-head", mode)

	return h
}
