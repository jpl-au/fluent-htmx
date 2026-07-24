package event

import (
	"strings"
	"testing"
)

// all maps each exported event constant to its value so the invariants below cover the whole
// package. The rules encode what makes an htmx 2 event name actually usable: every event is
// dispatched under the "htmx:" namespace, and htmx 2 also dispatches a fully lowercase kebab
// duplicate, which is the only form that survives the lowercasing an hx-on attribute undergoes
// (so it is the form that works through both HxOn and addEventListener).
var all = map[string]string{
	"Load": Load, "BeforeOnLoad": BeforeOnLoad, "AfterOnLoad": AfterOnLoad,
	"BeforeProcessNode": BeforeProcessNode, "AfterProcessNode": AfterProcessNode,
	"BeforeRequest": BeforeRequest, "BeforeSend": BeforeSend, "ConfigRequest": ConfigRequest,
	"AfterRequest": AfterRequest, "BeforeSwap": BeforeSwap, "AfterSwap": AfterSwap,
	"BeforeTransition": BeforeTransition, "AfterSettle": AfterSettle,
	"BeforeCleanupElement": BeforeCleanupElement, "Confirm": Confirm, "Prompt": Prompt,
	"Trigger": Trigger, "Abort": Abort,
	"BeforeHistorySave": BeforeHistorySave, "BeforeHistoryUpdate": BeforeHistoryUpdate,
	"PushedIntoHistory": PushedIntoHistory, "ReplacedInHistory": ReplacedInHistory,
	"HistoryItemCreated": HistoryItemCreated, "HistoryRestore": HistoryRestore,
	"HistoryCacheHit": HistoryCacheHit, "HistoryCacheMiss": HistoryCacheMiss,
	"HistoryCacheMissLoad": HistoryCacheMissLoad, "HistoryCacheMissLoadError": HistoryCacheMissLoadError,
	"HistoryCacheError": HistoryCacheError,
	"OOBBeforeSwap":     OOBBeforeSwap, "OOBAfterSwap": OOBAfterSwap, "OOBErrorNoTarget": OOBErrorNoTarget,
	"Error": Error, "ResponseError": ResponseError, "SendError": SendError, "SendAbort": SendAbort,
	"SwapError": SwapError, "TargetError": TargetError, "Timeout": Timeout, "OnLoadError": OnLoadError,
	"ValidationValidate": ValidationValidate, "ValidationFailed": ValidationFailed,
	"ValidationHalted": ValidationHalted,
	"XHRAbort":         XHRAbort, "XHRLoadStart": XHRLoadStart, "XHRLoadEnd": XHRLoadEnd, "XHRProgress": XHRProgress,
	"WSConnecting": WSConnecting, "WSOpen": WSOpen, "WSClose": WSClose, "WSError": WSError,
	"WSBeforeMessage": WSBeforeMessage, "WSAfterMessage": WSAfterMessage, "WSConfigSend": WSConfigSend,
	"WSBeforeSend": WSBeforeSend, "WSAfterSend": WSAfterSend,
	"SSEOpen": SSEOpen, "SSEError": SSEError, "SSEMessage": SSEMessage,
	"SSEBeforeMessage": SSEBeforeMessage, "NoSSESourceError": NoSSESourceError,
	"BeforeHeadMerge": BeforeHeadMerge, "AfterHeadMerge": AfterHeadMerge,
	"AddingHeadElement": AddingHeadElement, "RemovingHeadElement": RemovingHeadElement,
}

// TestEventNamesCarryHtmxPrefix asserts every constant is the full dispatched name. A bare
// suffix like "afterSwap" matches no event htmx dispatches, so it is useless for a JS listener.
func TestEventNamesCarryHtmxPrefix(t *testing.T) {
	for name, value := range all {
		if !strings.HasPrefix(value, "htmx:") {
			t.Errorf("%s = %q: must be the full dispatched name carrying the htmx: prefix", name, value)
		}
	}
}

// TestEventNamesFireThroughHxOn asserts every constant is lowercase. HxOn writes the name into
// an hx-on attribute, which the HTML parser lowercases, so any uppercase letter means the handler
// binds to an event htmx never fires. htmx 2's kebab duplicate is the lowercase form to use.
func TestEventNamesFireThroughHxOn(t *testing.T) {
	for name, value := range all {
		if value != strings.ToLower(value) {
			t.Errorf("%s = %q: has uppercase, so it will not match once the hx-on attribute is lowercased; use the kebab-case dispatched form", name, value)
		}
	}
}
