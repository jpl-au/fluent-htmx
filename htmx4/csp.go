package htmx

// Content-Security-Policy extension (hx-csp). Not bundled in htmax.js - load it as a separate
// script (dist/ext/hx-csp.js).

// HxNonce stamps a CSP nonce on the element. With the hx-csp extension loaded, htmx processes
// an element only when its hx-nonce matches the page's CSP nonce; an element with a missing or
// mismatched nonce is left inert. Stamp it from your templating engine on every element that
// carries htmx attributes. This is unrelated to Config().InlineScriptNonce, which nonces
// htmx's own injected inline scripts.
func (h *Wrapper) HxNonce(nonce string) *Wrapper {
	h.element.SetAttribute("hx-nonce", nonce)

	return h
}
