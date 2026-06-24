package htmx

// Native browser loading-indicator extension. Uses the hx-browser-indicator attribute. The
// extension is bundled in htmax.js; with the core htmx.js build, include the browser-indicator
// extension script yourself.

// HxBrowserIndicator drives the browser's own native page-loading indicator during the
// request when enabled, instead of (or as well as) an hx-indicator element. Handy for
// boosted, navigation-like requests where the native progress bar is the expected feedback.
func (h *Wrapper) HxBrowserIndicator(enabled bool) *Wrapper {
	value := boolFalse
	if enabled {
		value = boolTrue
	}

	h.element.SetAttribute("hx-browser-indicator", value)

	return h
}
