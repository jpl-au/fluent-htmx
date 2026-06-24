package htmx

// Boolean string values for HTMX attributes.
const (
	boolTrue  = "true"
	boolFalse = "false"
)

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
