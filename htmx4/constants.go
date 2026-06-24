package htmx

// Boolean string values for HTMX attributes.
const (
	boolTrue  = "true"
	boolFalse = "false"
)

// HTMX request headers sent by the client.
//
// The triggering element is reported in HX-Source (format "tagName#id") and the request
// type in HX-Request-Type ("full" or "partial"). HX-Prompt is not sent by htmx 4 (hx-prompt
// was removed and there is no prompt extension); see HxPrompt.
const (
	HXRequestHeader               = "HX-Request"
	HXBoostedHeader               = "HX-Boosted"
	HXCurrentURLHeader            = "HX-Current-URL"
	HXHistoryRestoreRequestHeader = "HX-History-Restore-Request"
	HXPromptHeader                = "HX-Prompt"
	HXTargetHeader                = "HX-Target"
	HXSourceHeader                = "HX-Source"
	HXRequestTypeHeader           = "HX-Request-Type"
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
)
