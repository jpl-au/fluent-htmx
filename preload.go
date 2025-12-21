package htmx

// Preload enables preloading on an element with the specified trigger
// Common values: "mousedown" (default), "mouseover", "always", or custom event names
func (h *HtmxWrapper) Preload(trigger string) *HtmxWrapper {
	h.node.SetAttribute("preload", trigger)
	return h
}

// PreloadImages enables preloading of linked image resources from preloaded HTML fragments
func (h *HtmxWrapper) PreloadImages(enabled bool) *HtmxWrapper {
	if enabled {
		h.node.SetAttribute("preload-images", "true")
	}
	return h
}
