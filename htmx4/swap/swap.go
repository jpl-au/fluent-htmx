// Package swap defines HTMX swap strategies that control how response content
// is inserted into the DOM.
package swap

// Strategy is the typed hx-swap value that controls how response content is inserted
// into the DOM; the named constants in this package are its valid values. HxSwap and
// the config default-swap setters take a Strategy rather than a free string, so swap
// styles are checked at compile time instead of failing silently at runtime.
type Strategy string

const (
	InnerHTML   Strategy = "innerHTML"   // Replace the inner html of the target element
	OuterHTML   Strategy = "outerHTML"   // Replace the entire target element with the response
	BeforeBegin Strategy = "beforebegin" // Insert the response before the target element
	AfterBegin  Strategy = "afterbegin"  // Insert the response before the first child of the target element
	BeforeEnd   Strategy = "beforeend"   // Insert the response after the last child of the target element
	AfterEnd    Strategy = "afterend"    // Insert the response after the target element
	Delete      Strategy = "delete"      // Deletes the target element regardless of the response
	None        Strategy = "none"        // Does not target any part of the DOM
)

// Morphing and text swap styles.
const (
	InnerMorph  Strategy = "innerMorph"  // Morph the children of the target using the idiomorph algorithm
	OuterMorph  Strategy = "outerMorph"  // Morph the target element itself using the idiomorph algorithm
	TextContent Strategy = "textContent" // Set the target's text content (no HTML parsing)
	OuterSync   Strategy = "outerSync"   // Replace the target's attributes and children in place, keeping the element itself; used for body swaps and history restores
)

// Extension swap styles. Each requires its htmx 4 extension, and both are bundled into
// htmax.js.
const (
	Download Strategy = "download" // Stream the response to the browser as a file download (download extension)
	Upsert   Strategy = "upsert"   // Update or insert response elements by id (upsert extension; supports sort, key: and prepend modifiers via swap.Custom)
)

// Custom creates a swap strategy string with modifiers after the style. The modifiers htmx
// reads are transition, settle, swapEmpty, strip, ignoreTitle, focusScroll, show and
// showTarget, scroll and scrollTarget, target (for out-of-band swaps), and the upsert
// extension's key, sort and prepend. The show and scroll modifiers take separate keys for
// position and target, e.g. swap.Custom("innerHTML show:top showTarget:#other").
func Custom(strategy string) Strategy {
	return Strategy(strategy)
}
