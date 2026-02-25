package htmx

// Boolean string values for HTMX attributes.
const (
	boolTrue  = "true"
	boolFalse = "false"
)

// SwapStrategy defines the strategy used when swapping content into the DOM.
type SwapStrategy string

const (
	SwapInnerHTML   SwapStrategy = "innerHTML"   // Replace the inner html of the target element
	SwapOuterHTML   SwapStrategy = "outerHTML"   // Replace the entire target element with the response
	SwapBeforeBegin SwapStrategy = "beforebegin" // Insert the response before the target element
	SwapAfterBegin  SwapStrategy = "afterbegin"  // Insert the response before the first child of the target element
	SwapBeforeEnd   SwapStrategy = "beforeend"   // Insert the response after the last child of the target element
	SwapAfterEnd    SwapStrategy = "afterend"    // Insert the response after the target element
	SwapDelete      SwapStrategy = "delete"      // Deletes the target element regardless of the response
	SwapNone        SwapStrategy = "none"        // Does not target any part of the DOM
)

// CustomSwap creates a custom swap strategy string, allowing for modifiers
// like "innerHTML swap:1s".
func CustomSwap(strategy string) SwapStrategy {
	return SwapStrategy(strategy)
}

// HTMX request headers sent by the client.
const (
	HXRequestHeader               = "HX-Request"
	HXBoostedHeader               = "HX-Boosted"
	HXCurrentURLHeader            = "HX-Current-URL"
	HXHistoryRestoreRequestHeader = "HX-History-Restore-Request"
	HXPromptHeader                = "HX-Prompt"
	HXTargetHeader                = "HX-Target"
	HXTriggerNameHeader           = "HX-Trigger-Name"
	HXTriggerHeader               = "HX-Trigger"
)

// HTMX response headers sent by the server.
const (
	HXLocationHeader           = "HX-Location"
	HXPushURLHeader            = "HX-Push-Url"
	HXRedirectHeader           = "HX-Redirect"
	HXRefreshHeader            = "HX-Refresh"
	HXReplaceURLHeader         = "HX-Replace-Url"
	HXReswapHeader             = "HX-Reswap"
	HXRetargetHeader           = "HX-Retarget"
	HXReselectHeader           = "HX-Reselect"
	HXTriggerAfterSettleHeader = "HX-Trigger-After-Settle"
	HXTriggerAfterSwapHeader   = "HX-Trigger-After-Swap"
)

// CSS classes applied by HTMX during the request lifecycle.
const (
	HXClassAdded     = "htmx-added"     // Applied to new content before swap, removed after settled
	HXClassIndicator = "htmx-indicator" // Toggles visible (opacity:1) when htmx-request is present
	HXClassRequest   = "htmx-request"   // Applied during requests to element or hx-indicator target
	HXClassSettling  = "htmx-settling"  // Applied to target after content swap, removed after settled
	HXClassSwapping  = "htmx-swapping"  // Applied to target before swap, removed after swapped
)

// HTMX events that can be used with HxOn() for event handling.
const (
	// Lifecycle events.
	EventAbort             = "abort"                // Send to abort a request
	EventAfterOnLoad       = "afterOnLoad"          // After AJAX response processing
	EventAfterProcessNode  = "afterProcessNode"     // After htmx initializes a node
	EventAfterRequest      = "afterRequest"         // After AJAX request completes
	EventAfterSettle       = "afterSettle"          // After DOM has settled
	EventAfterSwap         = "afterSwap"            // After new content swapped in
	EventBeforeCleanupElt  = "beforeCleanupElement" // Before element disabled or removed
	EventBeforeOnLoad      = "beforeOnLoad"         // Before response processing
	EventBeforeProcessNode = "beforeProcessNode"    // Before htmx initializes a node
	EventBeforeRequest     = "beforeRequest"        // Before AJAX request
	EventBeforeSwap        = "beforeSwap"           // Before swap, allows config
	EventBeforeSend        = "beforeSend"           // Just before request sent
	EventBeforeTransition  = "beforeTransition"     // Before View Transition swap
	EventConfigRequest     = "configRequest"        // Before request, customize params/headers
	EventConfirm           = "confirm"              // After trigger, can cancel request

	// History events.
	EventHistoryCacheError       = "historyCacheError"         // Error during cache writing
	EventHistoryCacheHit         = "historyCacheHit"           // Cache hit in history
	EventHistoryCacheMiss        = "historyCacheMiss"          // Cache miss in history
	EventHistoryCacheMissLoad    = "historyCacheMissLoad"      // Successful remote retrieval
	EventHistoryCacheMissLoadErr = "historyCacheMissLoadError" // Failed remote retrieval
	EventHistoryRestore          = "historyRestore"            // History restoration action
	EventBeforeHistorySave       = "beforeHistorySave"         // Before content saved to cache
	EventPushedIntoHistory       = "pushedIntoHistory"         // URL pushed into history
	EventReplacedInHistory       = "replacedInHistory"         // URL replaced in history

	// Content events.
	EventLoad = "load" // New content added to DOM

	// Out-of-band events.
	EventOOBAfterSwap     = "oobAfterSwap"     // After OOB element swapped
	EventOOBBeforeSwap    = "oobBeforeSwap"    // Before OOB element swap
	EventOOBErrorNoTarget = "oobErrorNoTarget" // OOB element has no matching ID

	// Error events.
	EventOnLoadError   = "onLoadError"   // Exception during onLoad handling
	EventResponseError = "responseError" // HTTP error response (non-200/300)
	EventSendAbort     = "sendAbort"     // Request aborted
	EventSendError     = "sendError"     // Network error prevents request
	EventSwapError     = "swapError"     // Error during swap phase
	EventTargetError   = "targetError"   // Invalid target specified
	EventTimeout       = "timeout"       // Request timeout

	// SSE events.
	EventNoSSESourceError = "noSSESourceError" // Element refers to SSE event but no source
	EventSSEError         = "sseError"         // Error with SSE source
	EventSSEOpen          = "sseOpen"          // SSE source opened

	// Validation events.
	EventValidationValidate = "validation:validate" // Before element validated
	EventValidationFailed   = "validation:failed"   // Element fails validation
	EventValidationHalted   = "validation:halted"   // Request halted due to validation

	// XHR events.
	EventXHRAbort     = "xhr:abort"     // AJAX request aborts
	EventXHRLoadEnd   = "xhr:loadend"   // AJAX request ends
	EventXHRLoadStart = "xhr:loadstart" // AJAX request starts
	EventXHRProgress  = "xhr:progress"  // AJAX request progress

	// Prompt event.
	EventPrompt = "prompt" // After prompt shown
)
