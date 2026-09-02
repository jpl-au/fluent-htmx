package htmx

// Preload extension. Uses the hx-preload attribute. The extension is bundled in htmax.js;
// with the core htmx.js build, include the preload extension script yourself.

// Preload triggers the element's GET request early so the response is ready on click.
// The value is a trigger spec: "mousedown" is the default and "mouseover" the eager
// choice, and a timeout modifier such as "mousedown timeout:2s" sets how long the
// preloaded response stays usable, 5 seconds by default. Only GET requests preload.
// A boosted link preloads on mousedown with no attribute at all.
func (h *Wrapper) Preload(trigger string, mods ...Mod) *Wrapper {
	h.element.SetAttribute(modifiedKey("hx-preload", mods), trigger)

	return h
}
