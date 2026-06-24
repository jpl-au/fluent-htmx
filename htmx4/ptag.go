package htmx

// Polling-tag extension (hx-ptag). Not bundled in htmax.js - load it as a separate script
// (dist/ext/hx-ptag.js).

// HxPtag sets a polling tag on an element that polls (for example hx-trigger="every 3s"). The
// server echoes the tag and, when the polled content has not changed, htmx skips the swap.
// Use an opaque version value such as a build id or a timestamp. Example: HxPtag("v42").
func (h *Wrapper) HxPtag(tag string) *Wrapper {
	h.element.SetAttribute("hx-ptag", tag)

	return h
}
