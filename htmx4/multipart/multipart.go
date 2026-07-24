// Package multipart defines the response content types that drive htmx's multipart streaming
// extension.
package multipart

// Type is the response Content-Type that drives the multipart extension; the named constants are
// its valid values. htmx.NewMultipart takes a Type rather than a free string, so the streaming
// mode is checked at compile time instead of failing silently at runtime.
type Type string

const (
	Mixed    Type = "multipart/mixed"    // Swaps parts in the order they arrive
	Parallel Type = "multipart/parallel" // Swaps each part as soon as it arrives
)
