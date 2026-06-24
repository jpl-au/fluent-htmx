package htmx

// Preload extension. Uses the hx-preload attribute. The extension is bundled in htmax.js;
// with the core htmx.js build, include the preload extension script yourself.

// Preload triggers the element's GET request early so the response is ready on click.
// The value is a trigger spec - common values are "mousedown" (default) and "mouseover",
// optionally with a timeout modifier.
func (h *Wrapper) Preload(trigger string) *Wrapper {
	h.element.SetAttribute("hx-preload", trigger)

	return h
}
