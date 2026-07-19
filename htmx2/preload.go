package htmx

// Preload enables preloading on an element with the specified trigger.
// Common values: "mousedown" (default), "mouseover", "always", or custom event names.
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
