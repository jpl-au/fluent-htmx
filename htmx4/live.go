package htmx

// Reactive expressions extension. Uses the hx-live attribute and the hx-live:* bindings. The
// extension is bundled in htmax.js; with the core htmx.js build, include dist/ext/hx-live.js
// yourself.
//
// Every expression is JavaScript, re-run whenever an input, change or DOM mutation happens.
// Inside an expression, q() selects elements reactively and debounce() delays re-evaluation.
// HxLive runs an expression for its effect. The binding methods write the expression's value
// to one place on the element: its text, its HTML, a class, its style, or an attribute.

// HxLive runs a JavaScript expression whenever its inputs change. Use it for an effect, such
// as updating another element through q(). Example: HxLive("q('#total').textContent = q('#qty').value * unitPrice").
func (h *Wrapper) HxLive(expression string) *Wrapper {
	h.element.SetAttribute("hx-live", expression)

	return h
}

// HxLiveText sets the element's text content to the expression's value whenever it changes.
// Example: HxLiveText("q('#qty').value * unitPrice").
func (h *Wrapper) HxLiveText(expression string) *Wrapper {
	h.element.SetAttribute("hx-live:text", expression)

	return h
}

// HxLiveHTML sets the element's inner HTML to the expression's value whenever it changes. The
// value is not escaped, so it must come from a trusted source.
func (h *Wrapper) HxLiveHTML(expression string) *Wrapper {
	h.element.SetAttribute("hx-live:html", expression)

	return h
}

// HxLiveClass sets the element's classes from the expression. The expression returns either
// a space-separated class string or an object of class name to boolean, and the extension
// adds and removes only the classes it manages.
// Example: HxLiveClass("{active: q('#tab').value === 'home'}").
func (h *Wrapper) HxLiveClass(expression string) *Wrapper {
	h.element.SetAttribute("hx-live:class", expression)

	return h
}

// HxLiveClassToggle adds the named class while the expression is truthy and removes it
// otherwise. Example: HxLiveClassToggle("hidden", "q('#items').length === 0").
func (h *Wrapper) HxLiveClassToggle(class string, expression string) *Wrapper {
	h.element.SetAttribute("hx-live:."+class, expression)

	return h
}

// HxLiveStyle sets inline styles from the expression. The expression returns either a CSS
// declaration string such as "color: red; width: 10px" or an object of property to value,
// and the extension removes any style it set earlier that the new value no longer names.
func (h *Wrapper) HxLiveStyle(expression string) *Wrapper {
	h.element.SetAttribute("hx-live:style", expression)

	return h
}

// HxLiveAttr sets the named attribute to the expression's value whenever it changes. A
// boolean value adds or removes the attribute. Example: HxLiveAttr("disabled", "q('#agree').checked === false").
func (h *Wrapper) HxLiveAttr(name string, expression string) *Wrapper {
	h.element.SetAttribute("hx-live:"+name, expression)

	return h
}
