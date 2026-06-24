// Package event defines htmx event names for use with HxOn() handlers and JavaScript
// listeners.
//
// Events follow the htmx:phase:action[:sub-action] scheme; most errors are consolidated
// into htmx:error. The constants below hold the full dispatched names.
package event

// Lifecycle events.
const (
	BeforeInit     = "htmx:before:init"     // Before htmx initialises a node
	AfterInit      = "htmx:after:init"      // After htmx initialises a node
	BeforeProcess  = "htmx:before:process"  // Before htmx processes a node
	AfterProcess   = "htmx:after:process"   // After htmx processes a node
	BeforeRequest  = "htmx:before:request"  // Before an AJAX request is sent
	AfterRequest   = "htmx:after:request"   // After an AJAX request completes
	FinallyRequest = "htmx:finally:request" // Always fires after a request, success or failure
	ConfigRequest  = "htmx:config:request"  // Before a request, to customise params/headers
	BeforeResponse = "htmx:before:response" // Before the response body is read (cancellable)
	BeforeSwap     = "htmx:before:swap"     // Before content is swapped in
	AfterSwap      = "htmx:after:swap"      // After content is swapped in
	BeforeSettle   = "htmx:before:settle"   // Before the settle phase
	AfterSettle    = "htmx:after:settle"    // After the settle phase
	BeforeCleanup  = "htmx:before:cleanup"  // Before an element is cleaned up
	AfterCleanup   = "htmx:after:cleanup"   // After an element is cleaned up
)

// History events.
const (
	BeforeHistoryUpdate  = "htmx:before:history:update"  // Before history is updated
	AfterHistoryUpdate   = "htmx:after:history:update"   // After history is updated
	BeforeHistoryRestore = "htmx:before:history:restore" // Before history is restored
	AfterHistoryPush     = "htmx:after:history:push"     // After a URL is pushed into history
	AfterHistoryReplace  = "htmx:after:history:replace"  // After a URL is replaced in history
)

// View transition events.
const (
	BeforeViewTransition = "htmx:before:viewTransition" // Before a view transition starts (cancellable)
	AfterViewTransition  = "htmx:after:viewTransition"  // After a view transition completes
)

// Error events.
const (
	ResponseError = "htmx:response:error" // HTTP error response
	Error         = "htmx:error"          // Consolidated error event
)

// Request control events.
const (
	Confirm = "htmx:confirm" // Fired to confirm a request; cancellable
	Abort   = "htmx:abort"   // Sent to an element to abort its in-flight request
)
