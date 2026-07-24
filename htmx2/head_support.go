package htmx

// Head-support extension. Uses the hx-head attribute. Enable it with HxExt("head-support") and
// include the head-support extension script; without it the attribute is a silent no-op.

// HxHead sets the hx-head attribute for controlling head element merge behaviour.
// Valid values: "merge" (default algorithm), "append" (append all), "re-eval" (force re-evaluation).
func (h *Wrapper) HxHead(mode string) *Wrapper {
	h.element.SetAttribute("hx-head", mode)

	return h
}
