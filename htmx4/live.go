package htmx

// Reactive expressions extension. Uses the hx-live attribute. The extension is bundled in
// htmax.js; with the core htmx.js build, include the live extension script yourself.

// HxLive binds a reactive JavaScript expression to the element. The expression is recomputed
// whenever its inputs change - on input, change, or DOM mutation - keeping the element in sync
// with client-side state without a server round-trip. Inside the expression, q() selects
// elements reactively. Example: HxLive("q('#qty').value * unitPrice").
func (h *Wrapper) HxLive(expression string) *Wrapper {
	h.element.SetAttribute("hx-live", expression)

	return h
}
