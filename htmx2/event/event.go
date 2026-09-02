// Package event defines htmx 2 event names for use with HxOn() handlers and JavaScript listeners.
// It covers core events and the events emitted by the extensions.
//
// The values are the kebab-case dispatched names (for example "htmx:after-request"). htmx 2
// dispatches every event under both its camelCase name and this fully lowercase kebab duplicate,
// and only the kebab form survives the lowercasing an hx-on attribute undergoes, so it is the
// form that works through both HxOn and addEventListener.
package event

// Lifecycle events.
const (
	Load                 = "htmx:load"                   // New content added to the DOM
	BeforeOnLoad         = "htmx:before-on-load"         // Before response processing
	AfterOnLoad          = "htmx:after-on-load"          // After AJAX response processing
	BeforeProcessNode    = "htmx:before-process-node"    // Before htmx initialises a node
	AfterProcessNode     = "htmx:after-process-node"     // After htmx initialises a node
	BeforeRequest        = "htmx:before-request"         // Before an AJAX request is issued
	BeforeSend           = "htmx:before-send"            // Just before a request is sent
	ConfigRequest        = "htmx:config-request"         // Before a request, to customise params/headers
	AfterRequest         = "htmx:after-request"          // After an AJAX request completes
	BeforeSwap           = "htmx:before-swap"            // Before content is swapped in; cancellable
	AfterSwap            = "htmx:after-swap"             // After content is swapped in
	BeforeTransition     = "htmx:before-transition"      // Before a View Transition swap
	AfterSettle          = "htmx:after-settle"           // After the DOM has settled
	BeforeCleanupElement = "htmx:before-cleanup-element" // Before an element is disabled or removed
	Confirm              = "htmx:confirm"                // Fired to confirm a request; cancellable
	Prompt               = "htmx:prompt"                 // After the prompt dialog is shown
	Trigger              = "htmx:trigger"                // Fired when an element is triggered
	Abort                = "htmx:abort"                  // Sent to an element to abort its request
)

// History events.
const (
	BeforeHistorySave         = "htmx:before-history-save"           // Before content is saved to the history cache
	BeforeHistoryUpdate       = "htmx:before-history-update"         // Before the history is updated
	PushedIntoHistory         = "htmx:pushed-into-history"           // A URL was pushed into history
	ReplacedInHistory         = "htmx:replaced-in-history"           // A URL was replaced in history
	HistoryItemCreated        = "htmx:history-item-created"          // A history cache entry was created
	HistoryRestore            = "htmx:history-restore"               // A history restoration occurred
	HistoryCacheHit           = "htmx:history-cache-hit"             // A history navigation hit the cache
	HistoryCacheMiss          = "htmx:history-cache-miss"            // A history navigation missed the cache
	HistoryCacheMissLoad      = "htmx:history-cache-miss-load"       // A cache miss was served by a remote load
	HistoryCacheMissLoadError = "htmx:history-cache-miss-load-error" // A cache-miss remote load failed
	HistoryCacheError         = "htmx:history-cache-error"           // An error writing the history cache
)

// Out-of-band swap events.
const (
	OOBBeforeSwap    = "htmx:oob-before-swap"     // Before an out-of-band element is swapped
	OOBAfterSwap     = "htmx:oob-after-swap"      // After an out-of-band element is swapped
	OOBErrorNoTarget = "htmx:oob-error-no-target" // An out-of-band element has no matching target
)

// Error events.
const (
	Error         = "htmx:error"          // Consolidated error event
	ResponseError = "htmx:response-error" // An HTTP error response (non-2xx/3xx)
	SendError     = "htmx:send-error"     // A network error prevented the request
	SendAbort     = "htmx:send-abort"     // A request was aborted
	SwapError     = "htmx:swap-error"     // An error during the swap phase
	TargetError   = "htmx:target-error"   // An invalid target was specified
	Timeout       = "htmx:timeout"        // A request timed out
	OnLoadError   = "htmx:on-load-error"  // An exception during onLoad handling
)

// Validation events.
const (
	ValidationValidate = "htmx:validation:validate" // Before an element is validated
	ValidationFailed   = "htmx:validation:failed"   // An element failed validation
	ValidationHalted   = "htmx:validation:halted"   // A request was halted by validation
)

// XHR progress events.
const (
	XHRAbort     = "htmx:xhr:abort"     // The underlying XHR aborted
	XHRLoadStart = "htmx:xhr:loadstart" // The underlying XHR started
	XHRLoadEnd   = "htmx:xhr:loadend"   // The underlying XHR ended
	XHRProgress  = "htmx:xhr:progress"  // The underlying XHR reported progress
)

// WebSocket extension events (ws-connect).
const (
	WSConnecting    = "htmx:ws-connecting"     // A WebSocket connection is being established
	WSOpen          = "htmx:ws-open"           // A WebSocket connection opened
	WSClose         = "htmx:ws-close"          // A WebSocket connection closed
	WSError         = "htmx:ws-error"          // A WebSocket error
	WSBeforeMessage = "htmx:ws-before-message" // Before a received WebSocket message is handled
	WSAfterMessage  = "htmx:ws-after-message"  // After a received WebSocket message is handled
	WSConfigSend    = "htmx:ws-config-send"    // Before a message is sent, to configure it
	WSBeforeSend    = "htmx:ws-before-send"    // Before a message is sent
	WSAfterSend     = "htmx:ws-after-send"     // After a message is sent
)

// Server-Sent Events extension events (sse-connect).
const (
	SSEOpen          = "htmx:sse-open"           // An SSE connection opened
	SSEError         = "htmx:sse-error"          // An SSE connection error
	SSEMessage       = "htmx:sse-message"        // An SSE message was received
	SSEClose         = "htmx:sse-close"          // An SSE connection closed; detail.type is nodeMissing, nodeReplaced or message
	SSEBeforeMessage = "htmx:sse-before-message" // Before an SSE message is handled
	NoSSESourceError = "htmx:no-ssesource-error" // An element references SSE with no source
)

// Head-support extension events (hx-head).
const (
	BeforeHeadMerge     = "htmx:before-head-merge"     // Before the document head is merged
	AfterHeadMerge      = "htmx:after-head-merge"      // After the document head is merged
	AddingHeadElement   = "htmx:adding-head-element"   // Before an element is added to the head
	RemovingHeadElement = "htmx:removing-head-element" // Before an element is removed from the head
)
