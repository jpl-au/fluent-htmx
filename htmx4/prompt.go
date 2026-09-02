package htmx

// Prompt extension (hx-prompt). Restored from htmx 2 in htmx 4 as an extension - not bundled
// in htmx.js; load it as a separate script (dist/ext/hx-prompt.js).

// HxPrompt shows the browser prompt dialog with question before the request is sent and puts
// the answer in the HX-Prompt request header. Cancelling the dialog aborts the request; an
// empty answer is sent. The prompt runs before hx-confirm, and the htmx:prompt event fires
// with the answer so a handler can veto the request. Override window.htmxPrompt on the
// client to supply a custom, synchronous dialog in place of the native prompt.
func (h *Wrapper) HxPrompt(question string, mods ...Mod) *Wrapper {
	h.element.SetAttribute(modifiedKey("hx-prompt", mods), question)

	return h
}
