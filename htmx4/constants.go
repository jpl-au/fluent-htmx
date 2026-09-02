package htmx

// Boolean string values for HTMX attributes.
const (
	boolTrue  = "true"
	boolFalse = "false"
)

// HTMX request headers sent by the client.
//
// The triggering element is reported in HX-Source (format "tagName#id") and the request
// type in HX-Request-Type ("full" or "partial"). HX-Prompt is not sent by htmx 4 core; it is
// sent by the hx-prompt extension (restored in beta5), which pairs with HxPrompt.
const (
	HXRequestHeader               = "HX-Request"
	HXBoostedHeader               = "HX-Boosted"
	HXCurrentURLHeader            = "HX-Current-URL"
	HXHistoryRestoreRequestHeader = "HX-History-Restore-Request"
	HXPromptHeader                = "HX-Prompt"
	HXTargetHeader                = "HX-Target"
	HXSourceHeader                = "HX-Source"
	HXRequestTypeHeader           = "HX-Request-Type"
	HXPTagHeader                  = "HX-PTag" // Sent by the ptag extension with the stored tag; the server answers with the same header, see HxPTag
)

// Resume headers sent by the client when a stream reconnects. The SSE extension sends the
// id of the last event it handled as Last-Event-ID, and the multipart extension sends the
// HX-Part-ID of the last part it swapped as HX-Last-Part-ID. Read them to resume a stream
// from where the browser left off; see [SSEWriter.SendEvent] and [PartID].
const (
	LastEventIDHeader  = "Last-Event-ID"
	HXLastPartIDHeader = "HX-Last-Part-ID"
)

// HTMX response headers sent by the server.
//
// Events are triggered via the HX-Trigger header (or client-side JavaScript). HX-Download is
// read by the bundled download extension and set with HxDownload.
const (
	HXTriggerHeader    = "HX-Trigger"
	HXLocationHeader   = "HX-Location"
	HXPushURLHeader    = "HX-Push-Url"
	HXRedirectHeader   = "HX-Redirect"
	HXRefreshHeader    = "HX-Refresh"
	HXReplaceURLHeader = "HX-Replace-Url"
	HXReswapHeader     = "HX-Reswap"
	HXRetargetHeader   = "HX-Retarget"
	HXReselectHeader   = "HX-Reselect"
	HXDownloadHeader   = "HX-Download"
	HXPartIDHeader     = "HX-Part-ID" // On a multipart part; set with PartID
)
