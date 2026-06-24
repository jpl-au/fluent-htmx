// Package event defines HTMX event names for use with HxOn() event handlers.
package event

// Lifecycle events.
const (
	Abort             = "abort"                // Send to abort a request
	AfterOnLoad       = "afterOnLoad"          // After AJAX response processing
	AfterProcessNode  = "afterProcessNode"     // After htmx initialises a node
	AfterRequest      = "afterRequest"         // After AJAX request completes
	AfterSettle       = "afterSettle"          // After DOM has settled
	AfterSwap         = "afterSwap"            // After new content swapped in
	BeforeCleanupElt  = "beforeCleanupElement" // Before element disabled or removed
	BeforeOnLoad      = "beforeOnLoad"         // Before response processing
	BeforeProcessNode = "beforeProcessNode"    // Before htmx initialises a node
	BeforeRequest     = "beforeRequest"        // Before AJAX request
	BeforeSwap        = "beforeSwap"           // Before swap, allows config
	BeforeSend        = "beforeSend"           // Just before request sent
	BeforeTransition  = "beforeTransition"     // Before View Transition swap
	ConfigRequest     = "configRequest"        // Before request, customise params/headers
	Confirm           = "confirm"              // After trigger, can cancel request
)

// History events.
const (
	HistoryCacheError       = "historyCacheError"         // Error during cache writing
	HistoryCacheHit         = "historyCacheHit"           // Cache hit in history
	HistoryCacheMiss        = "historyCacheMiss"          // Cache miss in history
	HistoryCacheMissLoad    = "historyCacheMissLoad"      // Successful remote retrieval
	HistoryCacheMissLoadErr = "historyCacheMissLoadError" // Failed remote retrieval
	HistoryRestore          = "historyRestore"            // History restoration action
	BeforeHistorySave       = "beforeHistorySave"         // Before content saved to cache
	PushedIntoHistory       = "pushedIntoHistory"         // URL pushed into history
	ReplacedInHistory       = "replacedInHistory"         // URL replaced in history
)

// Content events.
const (
	Load = "load" // New content added to DOM
)

// Out-of-band events.
const (
	OOBAfterSwap     = "oobAfterSwap"     // After OOB element swapped
	OOBBeforeSwap    = "oobBeforeSwap"    // Before OOB element swap
	OOBErrorNoTarget = "oobErrorNoTarget" // OOB element has no matching ID
)

// Error events.
const (
	OnLoadError   = "onLoadError"   // Exception during onLoad handling
	ResponseError = "responseError" // HTTP error response (non-200/300)
	SendAbort     = "sendAbort"     // Request aborted
	SendError     = "sendError"     // Network error prevents request
	SwapError     = "swapError"     // Error during swap phase
	TargetError   = "targetError"   // Invalid target specified
	Timeout       = "timeout"       // Request timeout
)

// SSE events.
const (
	NoSSESourceError = "noSSESourceError" // Element refers to SSE event but no source
	SSEError         = "sseError"         // Error with SSE source
	SSEOpen          = "sseOpen"          // SSE source opened
)

// Validation events.
const (
	ValidationValidate = "validation:validate" // Before element validated
	ValidationFailed   = "validation:failed"   // Element fails validation
	ValidationHalted   = "validation:halted"   // Request halted due to validation
)

// XHR events.
const (
	XHRAbort     = "xhr:abort"     // AJAX request aborts
	XHRLoadEnd   = "xhr:loadend"   // AJAX request ends
	XHRLoadStart = "xhr:loadstart" // AJAX request starts
	XHRProgress  = "xhr:progress"  // AJAX request progress
)

// Prompt event.
const (
	Prompt = "prompt" // After prompt shown
)
