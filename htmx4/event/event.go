// Package event defines htmx event names for use with HxOn() handlers and JavaScript listeners.
// It covers htmx 4 core events and the events emitted by the extensions (both the ones bundled
// into htmax.js and the separate scripts).
//
// Events follow the htmx:phase:action[:sub-action] scheme; most errors are consolidated
// into htmx:error. The constants below hold the full dispatched names. The htmx 2 event
// names that the htmx-2-compat extension dispatches alongside these are not listed; the
// htmx2 package holds them. The extension callbacks that never reach the DOM (htmx:scope,
// htmx:process:*, htmx:before:morph:attr, htmx:before:morph:node and
// htmx:after:implicitInheritance) are not listed either.
package event

// Lifecycle events.
const (
	BeforeInit     = "htmx:before:init"     // Before htmx initialises a node
	AfterInit      = "htmx:after:init"      // After htmx initialises a node
	BeforeOnInit   = "htmx:before:on:init"  // Before an hx-on handler is initialised
	BeforeProcess  = "htmx:before:process"  // Before htmx processes a node
	AfterProcess   = "htmx:after:process"   // After htmx processes a node
	BeforeRequest  = "htmx:before:request"  // Before an AJAX request is sent
	AfterRequest   = "htmx:after:request"   // After an AJAX request completes
	FinallyRequest = "htmx:finally:request" // Always fires after a request, success or failure
	ConfigRequest  = "htmx:config:request"  // Before a request, to customise params/headers
	BeforeResponse = "htmx:before:response" // Before the response body is read (cancellable)
	BeforeSwap     = "htmx:before:swap"     // Before content is swapped in
	AfterSwap      = "htmx:after:swap"      // After content is swapped in
	FinallySwap    = "htmx:finally:swap"    // Always fires after a swap, success or failure
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

// Download extension events (hx-download / swap.Download).
const (
	DownloadStart    = "htmx:download:start"    // A file download begins
	DownloadProgress = "htmx:download:progress" // Progress fires as a download streams
	DownloadComplete = "htmx:download:complete" // A file download finishes
)

// Server-Sent Events extension events (hx-sse). The before events are cancellable; the
// before-message event also lets a listener call detail.waitUntil(promise) to delay handling.
const (
	SSEBeforeConnection = "htmx:sse:before:connection" // Before an SSE connection opens or reconnects
	SSEAfterConnection  = "htmx:sse:after:connection"  // After an SSE connection opens
	SSEBeforeMessage    = "htmx:sse:before:message"    // Before an SSE message is handled
	SSEAfterMessage     = "htmx:sse:after:message"     // After an SSE message is handled
	SSEClose            = "htmx:sse:close"             // An SSE connection closes
	SSEError            = "htmx:sse:error"             // An SSE connection error
)

// WebSocket extension events (hx-ws). The before events are cancellable; the two
// before-message events also let a listener call detail.waitUntil(promise) to delay the message.
const (
	WSBeforeConnection      = "htmx:ws:before:connection"       // Before a WebSocket connection opens or reconnects
	WSAfterConnection       = "htmx:ws:after:connection"        // After a WebSocket connection opens
	WSBeforeMessageIncoming = "htmx:ws:before:message:incoming" // Before a received message is handled
	WSAfterMessageIncoming  = "htmx:ws:after:message:incoming"  // After a received message is handled
	WSBeforeMessageOutgoing = "htmx:ws:before:message:outgoing" // Before a message is sent to the server
	WSAfterMessageOutgoing  = "htmx:ws:after:message:outgoing"  // After a message is sent to the server
	WSClose                 = "htmx:ws:close"                   // A WebSocket connection closes
	WSError                 = "htmx:ws:error"                   // A WebSocket connection error
)

// History-cache extension events (hx-history).
const (
	HistoryCacheBeforeSave    = "htmx:history:cache:before:save"    // Before the page is saved to the history cache
	HistoryCacheAfterSave     = "htmx:history:cache:after:save"     // After the page is saved to the history cache
	HistoryCacheBeforeRestore = "htmx:history:cache:before:restore" // Before a page is restored from the history cache
	HistoryCacheAfterRestore  = "htmx:history:cache:after:restore"  // After a page is restored from the history cache
	HistoryCacheHit           = "htmx:history:cache:hit"            // A history navigation was served from the cache
	HistoryCacheMiss          = "htmx:history:cache:miss"           // A history navigation missed the cache
)

// Head-support extension events (hx-head).
const (
	HeadBeforeMerge  = "htmx:head:before:merge"  // Before the document head is merged
	HeadAfterMerge   = "htmx:head:after:merge"   // After the document head is merged
	HeadBeforeAdd    = "htmx:head:before:add"    // Before an element is added to the head
	HeadBeforeRemove = "htmx:head:before:remove" // Before an element is removed from the head
)

// CSP extension events (hx-nonce).
const (
	SecurityStrip     = "htmx:security:strip"     // Every hx-* attribute was stripped from an element with a missing or wrong nonce
	SecurityViolation = "htmx:security:violation" // A content-security-policy violation was detected
)

// Prompt extension events (hx-prompt).
const (
	Prompt = "htmx:prompt" // After the prompt dialog returns an answer and before the request is sent; a listener can cancel the request
)

// Multipart streaming extension events (hx-multipart).
const (
	MultipartBeforeConnection = "htmx:multipart:before:connection" // Before a multipart connection opens; cancellable
	MultipartAfterConnection  = "htmx:multipart:after:connection"  // After a multipart connection opens
	MultipartBeforePart       = "htmx:multipart:before:part"       // Before a part is swapped; cancellable
	MultipartAfterPart        = "htmx:multipart:after:part"        // After a part is swapped
	MultipartError            = "htmx:multipart:error"             // A multipart streaming or reconnect error
	MultipartClose            = "htmx:multipart:close"             // After a multipart connection closes
)
