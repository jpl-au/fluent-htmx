package htmx

// Multi-target extension. Uses the hx-targets attribute. The extension is bundled in htmax.js;
// with the core htmx.js build, include the targets extension script yourself.

// HxTargets swaps the response into every element matching the selector, rather than a single
// hx-target, which it overrides. The response fragment is cloned once per match, and the
// element's hx-swap style applies to each. The selector may use the extended forms such as
// find or closest and is resolved relative to the requesting element. With no match htmx
// logs a warning and falls back to the normal target.
func (h *Wrapper) HxTargets(selector string, mods ...Mod) *Wrapper {
	h.element.SetAttribute(modifiedKey("hx-targets", mods), selector)

	return h
}
