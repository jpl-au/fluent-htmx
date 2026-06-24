// Package sync defines HTMX sync strategies that coordinate requests
// between elements to prevent race conditions.
package sync

// Strategy defines how overlapping requests from multiple elements
// are coordinated.
type Strategy string

const (
	Drop       Strategy = "drop"        // Drop the new request if an existing request is in flight
	Abort      Strategy = "abort"       // Abort the existing request and issue the new one
	Replace    Strategy = "replace"     // Replace the existing request (abort + reissue)
	QueueFirst Strategy = "queue first" // Queue the first request to execute after the current one finishes
	QueueLast  Strategy = "queue last"  // Queue the last request, discarding earlier queued requests
	QueueAll   Strategy = "queue all"   // Queue all requests in order
)

// Custom creates a custom sync strategy string, allowing for element
// selectors like "closest form:abort".
func Custom(strategy string) Strategy {
	return Strategy(strategy)
}
