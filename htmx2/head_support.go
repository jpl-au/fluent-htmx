package htmx

// Head-support extension. Uses the hx-head attribute. Enable it with HxExt("head-support") and
// include the head-support extension script; without it the attribute is a silent no-op.

// HxHead sets the hx-head attribute, which controls how a response <head> is applied. On
// the response <head> itself, "merge" keeps exact matches, adds new elements and removes
// missing ones, and "append" adds without removing; without the attribute a boosted request
// merges and any other appends. On an individual head element, "re-eval" removes and re-adds
// it on every request, and hx-preserve="true" keeps it through a merge that omits it.
func (h *Wrapper) HxHead(mode string) *Wrapper {
	h.element.SetAttribute("hx-head", mode)

	return h
}
