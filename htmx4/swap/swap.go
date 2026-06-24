// Package swap defines HTMX swap strategies that control how response content
// is inserted into the DOM.
package swap

// Strategy defines the strategy used when swapping content into the DOM.
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
)

// Extension swap styles. Each requires its htmx 4 extension to be loaded. The download
// extension is bundled into htmax.js; the upsert extension is a separate script.
const (
	Download Strategy = "download" // Stream the response to the browser as a file download (download extension)
	Upsert   Strategy = "upsert"   // Update or insert response elements by id (upsert extension; supports sort, key: and prepend modifiers via swap.Custom)
)

// Custom creates a custom swap strategy string, allowing for modifiers.
// The show and scroll modifiers take separate keys for position and target, e.g.
// swap.Custom("innerHTML show:top showTarget:#other").
func Custom(strategy string) Strategy {
	return Strategy(strategy)
}
