package htmx

// Preload enables preloading on an element with the specified trigger.
// Common values: "mousedown" (default), "mouseover", "always", or custom event names.
func (h *Wrapper) Preload(trigger string) *Wrapper {
	h.element.SetAttribute("preload", trigger)

	return h
}

// PreloadImages enables preloading of linked image resources from preloaded HTML fragments.
func (h *Wrapper) PreloadImages(enabled bool) *Wrapper {
	value := boolFalse
	if enabled {
		value = boolTrue
	}

	h.element.SetAttribute("preload-images", value)

	return h
}
