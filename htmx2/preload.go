package htmx

// Preload extension. Uses the preload / preload-images attributes. Enable it with HxExt("preload")
// and include the preload extension script; without it the attributes are silent no-ops.

// Preload enables preloading on an element with the specified trigger: "mousedown" (the
// default, and what a bare attribute means), "mouseover" (after a 100ms hover), a custom
// event name such as "preload:init", or "always" to preload on every trigger rather than
// once; "always" combines with another value, as in "always mouseover". Only GET requests
// preload. The attribute is inherited, so a container preloads every link within it.
func (h *Wrapper) Preload(trigger string) *Wrapper {
	h.element.SetAttribute("preload", trigger)

	return h
}

// PreloadImages opts the preload extension into also fetching images found in a
// preloaded response, warming the browser image cache alongside the HTML itself.
// It defaults to off because preloading images costs extra bandwidth; enable it
// when a preloaded fragment is image-heavy and you want those images ready the
// moment the fragment is swapped in.
func (h *Wrapper) PreloadImages(enabled bool) *Wrapper {
	value := boolFalse
	if enabled {
		value = boolTrue
	}

	h.element.SetAttribute("preload-images", value)

	return h
}
