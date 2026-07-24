package event

import (
	"strings"
	"testing"
)

// all maps each exported event constant to its value so the invariants below cover the whole
// package. The rule mirrors htmx2/event: every constant is the full dispatched name under the
// "htmx:" namespace.
var all = map[string]string{
	"BeforeInit": BeforeInit, "AfterInit": AfterInit, "BeforeOnInit": BeforeOnInit,
	"BeforeProcess": BeforeProcess, "AfterProcess": AfterProcess, "BeforeRequest": BeforeRequest,
	"AfterRequest": AfterRequest, "FinallyRequest": FinallyRequest, "ConfigRequest": ConfigRequest,
	"BeforeResponse": BeforeResponse, "BeforeSwap": BeforeSwap, "AfterSwap": AfterSwap,
	"FinallySwap": FinallySwap, "BeforeSettle": BeforeSettle, "AfterSettle": AfterSettle,
	"BeforeCleanup": BeforeCleanup, "AfterCleanup": AfterCleanup,
	"AfterImplicitInheritance": AfterImplicitInheritance,
	"BeforeMorphAttr":          BeforeMorphAttr, "BeforeMorphNode": BeforeMorphNode,
	"BeforeHistoryUpdate": BeforeHistoryUpdate, "AfterHistoryUpdate": AfterHistoryUpdate,
	"BeforeHistoryRestore": BeforeHistoryRestore, "AfterHistoryPush": AfterHistoryPush,
	"AfterHistoryReplace":  AfterHistoryReplace,
	"BeforeViewTransition": BeforeViewTransition, "AfterViewTransition": AfterViewTransition,
	"ResponseError": ResponseError, "Error": Error, "Confirm": Confirm, "Abort": Abort,
	"DownloadStart": DownloadStart, "DownloadProgress": DownloadProgress, "DownloadComplete": DownloadComplete,
	"BeforeSSEConnection": BeforeSSEConnection, "AfterSSEConnection": AfterSSEConnection,
	"BeforeSSEMessage": BeforeSSEMessage, "AfterSSEMessage": AfterSSEMessage,
	"SSEClose": SSEClose, "SSEError": SSEError,
	"BeforeWSConnection": BeforeWSConnection, "AfterWSConnection": AfterWSConnection,
	"BeforeWSMessage": BeforeWSMessage, "AfterWSMessage": AfterWSMessage,
	"BeforeWSRequest": BeforeWSRequest, "AfterWSRequest": AfterWSRequest,
	"WSClose": WSClose, "WSError": WSError,
	"HistoryCacheBeforeSave": HistoryCacheBeforeSave, "HistoryCacheAfterSave": HistoryCacheAfterSave,
	"HistoryCacheBeforeRestore": HistoryCacheBeforeRestore, "HistoryCacheAfterRestore": HistoryCacheAfterRestore,
	"HistoryCacheHit": HistoryCacheHit, "HistoryCacheMiss": HistoryCacheMiss,
	"BeforeHeadMerge": BeforeHeadMerge, "AfterHeadMerge": AfterHeadMerge,
	"BeforeHeadAdd": BeforeHeadAdd, "BeforeHeadRemove": BeforeHeadRemove,
	"SecurityStrip": SecurityStrip, "SecurityViolation": SecurityViolation, "Prompt": Prompt,
	"MultipartBeforeConnection": MultipartBeforeConnection, "MultipartAfterConnection": MultipartAfterConnection,
	"MultipartBeforePart": MultipartBeforePart, "MultipartAfterPart": MultipartAfterPart,
	"MultipartError": MultipartError, "MultipartClose": MultipartClose,
}

// hxOnUnsafe lists the events htmx 4 dispatches in camelCase. Because the HTML parser lowercases
// hx-on attribute names, these cannot be bound through HxOn and are addEventListener-only; htmx 4,
// unlike htmx 2, dispatches no lowercase duplicate. Kept as an explicit set so the test below can
// hold every other event to the hx-on-safe (lowercase) rule.
var hxOnUnsafe = map[string]bool{
	"BeforeViewTransition":     true,
	"AfterViewTransition":      true,
	"AfterImplicitInheritance": true,
}

// TestEventNamesCarryHtmxPrefix asserts every constant is the full dispatched name. A bare suffix
// matches no event htmx dispatches, so it is useless for a JS listener. Mirrors htmx2/event.
func TestEventNamesCarryHtmxPrefix(t *testing.T) {
	for name, value := range all {
		if !strings.HasPrefix(value, "htmx:") {
			t.Errorf("%s = %q: must be the full dispatched name carrying the htmx: prefix", name, value)
		}
	}
}

// TestEventNamesFireThroughHxOn asserts every event is lowercase (so it survives the lowercasing
// an hx-on attribute undergoes), except the documented camelCase events htmx 4 dispatches, which
// are addEventListener-only. It also fails if a listed exception has become lowercase, which would
// mean the exception is stale and should be dropped.
func TestEventNamesFireThroughHxOn(t *testing.T) {
	for name, value := range all {
		isLower := value == strings.ToLower(value)

		if hxOnUnsafe[name] {
			if isLower {
				t.Errorf("%s = %q is lowercase and hx-on-safe; remove it from hxOnUnsafe", name, value)
			}
			continue
		}

		if !isLower {
			t.Errorf("%s = %q has uppercase but is not a known addEventListener-only event; it will not fire through HxOn", name, value)
		}
	}
}
