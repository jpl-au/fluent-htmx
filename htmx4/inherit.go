package htmx

// Mod is an attribute modifier that controls whether an attribute inherits down the DOM
// tree. Pass it to any inheritable Hx* setter - those whose signature ends in ...Mod.
//
// Without a modifier an attribute applies only to the element it is set on. A modifier
// opts the attribute into inheritance, so descendant elements pick it up as well as the
// element itself.
//
// A handful of setters take no modifier because htmx reads them by presence or by id on a
// single element rather than resolving them down the tree: HxPreserve, HxSwapOOB,
// HxHistoryElt, HxIgnore, HxMorphSkip, HxMorphSkipChildren and HxOn. Everything else - including HxConfig, HxStatus and
// HxSelectOOB - is resolved with inheritance and accepts a modifier. The request verbs,
// HxAction and HxMethod technically inherit too, but the setters omit the modifier because
// an inherited verb cannot by itself initiate a request. For anything not covered by a
// typed setter, fall back to SetAttribute.
//
// The modifier values below assume htmx's default ":" meta character. If you change it via
// Config().MetaCharacter, set inherited attributes with SetAttribute using the configured
// separator instead, as these constants are not derived from the config.
type Mod string

const (
	// Inherited makes the attribute apply to its element and inherit to all descendants,
	// by appending :inherited to the attribute name (e.g. hx-confirm:inherited).
	Inherited Mod = ":inherited"

	// InheritedAppend appends this element's value to the value inherited from an ancestor
	// that carries Inherited or InheritedAppend, and passes the combined value on to its own
	// descendants (e.g. hx-include:inherited:append). A descendant that sets the plain
	// attribute still replaces; it needs Append or InheritedAppend to append. Use it for the
	// value-merging attributes such as HxInclude, HxHeaders and HxVals.
	InheritedAppend Mod = ":inherited:append"

	// Append makes a descendant append to an ancestor's inherited value without itself
	// re-propagating to its own descendants (e.g. hx-include:append). This is the terminal
	// counterpart to InheritedAppend; reach for it only when you need that distinction.
	Append Mod = ":append"
)

// modifiedKey appends any modifier suffixes to a base attribute name. With no modifiers it
// returns the base name unchanged, so the common (non-inherited) call is unaffected.
func modifiedKey(base string, mods []Mod) string {
	for _, m := range mods {
		base += string(m)
	}

	return base
}
