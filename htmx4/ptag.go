package htmx

import "net/http"

// Polling-tag extension (hx-ptag). Not bundled in htmax.js - load it as a separate script
// (dist/ext/hx-ptag.js).

// HxPtag sets a polling tag on an element that polls (for example hx-trigger="every 3s"). The
// server echoes the tag and, when the polled content has not changed, htmx skips the swap.
// Use an opaque version value such as a build id or a timestamp. Example: HxPtag("v42").
func (h *Wrapper) HxPtag(tag string) *Wrapper {
	h.element.SetAttribute("hx-ptag", tag)

	return h
}

// HxPTag returns the polling tag the client sends as the HX-PTag request header: the value
// of the element's hx-ptag attribute on its first poll, then the tag from the last response.
// It is "" on a first poll of an element with no hx-ptag attribute, or when the extension is
// not loaded.
func HxPTag(r *http.Request) string {
	return r.Header.Get(HXPTagHeader)
}

// HxPTagResponse sets the HX-PTag response header, the tag the client stores and sends back
// on its next poll.
func HxPTagResponse(w http.ResponseWriter, tag string) {
	w.Header().Set(HXPTagHeader, tag)
}

// HxPTagUnchanged is the whole polling exchange in one call. When the client's tag equals
// tag, it answers 304 Not Modified, which skips the swap, and returns true so the handler
// can return. Otherwise it sets the HX-PTag response header to tag and returns false, and
// the handler renders the content.
//
//	if htmx.HxPTagUnchanged(w, r, version) {
//	    return
//	}
func HxPTagUnchanged(w http.ResponseWriter, r *http.Request, tag string) bool {
	if HxPTag(r) == tag {
		w.WriteHeader(http.StatusNotModified)
		return true
	}
	HxPTagResponse(w, tag)

	return false
}
