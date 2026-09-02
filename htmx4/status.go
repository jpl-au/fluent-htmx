package htmx

// HxStatus sets per-status-code swap behaviour via the hx-status:CODE attribute.
//
// The code may be an exact status (e.g. "404"), a single-digit wildcard ("50x"),
// or a range wildcard ("5xx"); rules are evaluated in order of specificity. The
// noSwap config, 204 and 304 by default, is checked first, so a rule for one of
// those codes never runs. The ":" in the attribute name is not derived from
// Config().MetaCharacter.
//
// The spec uses htmx's key:value (HCON) syntax with the keys swap:, target:,
// select:, push:, replace: and transition:. Example:
//
//	htmx.New(form).
//	    HxStatus("422", "swap:innerHTML target:#errors").
//	    HxStatus("5xx", "swap:none push:false")
//
// htmx resolves hx-status:CODE with inheritance, so it accepts the inheritance modifier.
func (h *Wrapper) HxStatus(code string, spec string, mods ...Mod) *Wrapper {
	h.element.SetAttribute(modifiedKey("hx-status:"+code, mods), spec)

	return h
}
