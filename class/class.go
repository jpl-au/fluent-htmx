// Package class defines CSS class names applied by HTMX during the request lifecycle.
package class

const (
	Added     = "htmx-added"     // Applied to new content before swap, removed after settled
	Indicator = "htmx-indicator" // Toggles visible (opacity:1) when htmx-request is present
	Request   = "htmx-request"   // Applied during requests to element or hx-indicator target
	Settling  = "htmx-settling"  // Applied to target after content swap, removed after settled
	Swapping  = "htmx-swapping"  // Applied to target before swap, removed after swapped
)
