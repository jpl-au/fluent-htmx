package htmx

// HxHead sets the hx-head attribute for controlling head element merge behaviour.
// Valid values: "merge" (default algorithm), "append" (append all), "re-eval" (force re-evaluation).
func (h *Wrapper) HxHead(mode string) *Wrapper {
	h.node.SetAttribute("hx-head", mode)

	return h
}
