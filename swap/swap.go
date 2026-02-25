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

// Custom creates a custom swap strategy string, allowing for modifiers
// like "innerHTML swap:1s".
func Custom(strategy string) Strategy {
	return Strategy(strategy)
}
