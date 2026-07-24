package htmx

import (
	"strings"
	"testing"

	"github.com/jpl-au/fluent-htmx/htmx2/swap"
	"github.com/jpl-au/fluent-htmx/htmx2/sync"
	"github.com/jpl-au/fluent/html5/div"
)

func TestNew(t *testing.T) {
	d := div.New()
	w := New(d)

	if w == nil {
		t.Fatal("New() returned nil")
	}

	if w.element != d {
		t.Error("Wrapper does not wrap the provided element")
	}
}

func TestElementInterfaceDelegation(t *testing.T) {
	d := div.New().ID("test")
	w := New(d)

	html := string(w.RenderBytes())
	if !strings.Contains(html, `id="test"`) {
		t.Errorf("Render() delegation failed, got: %s", html)
	}

	nodes := w.Nodes()
	_ = nodes // Nodes() delegation successful if no panic
}

func TestMethodChaining(t *testing.T) {
	d := div.New()
	w := New(d)

	// All methods must return *Wrapper to support chaining.
	result := w.HxGet("/api/users").HxTarget("#result").HxSwap(swap.InnerHTML).HxTrigger("click")

	if result == nil {
		t.Fatal("method chaining returned nil")
	}

	html := string(d.RenderBytes())
	for _, attr := range []string{`hx-get="/api/users"`, `hx-target="#result"`, `hx-swap="innerHTML"`, `hx-trigger="click"`} {
		if !strings.Contains(html, attr) {
			t.Errorf("chained call missing %s, got: %s", attr, html)
		}
	}
}

// setterCase asserts a single client setter renders the expected attribute fragment.
type setterCase struct {
	name string
	set  func(*Wrapper)
	want string
}

func TestClientSetters(t *testing.T) {
	cases := []setterCase{
		{"HxGet", func(w *Wrapper) { w.HxGet("/api/users") }, `hx-get="/api/users"`},
		{"HxPost", func(w *Wrapper) { w.HxPost("/api/users") }, `hx-post="/api/users"`},
		{"HxPut", func(w *Wrapper) { w.HxPut("/api/users/1") }, `hx-put="/api/users/1"`},
		{"HxPatch", func(w *Wrapper) { w.HxPatch("/api/users/1") }, `hx-patch="/api/users/1"`},
		{"HxDelete", func(w *Wrapper) { w.HxDelete("/api/users/1") }, `hx-delete="/api/users/1"`},
		{"HxSwap", func(w *Wrapper) { w.HxSwap(swap.OuterHTML) }, `hx-swap="outerHTML"`},
		{"HxSwapCustom", func(w *Wrapper) { w.HxSwap(swap.Custom("innerHTML swap:1s")) }, `hx-swap="innerHTML swap:1s"`},
		{"HxTarget", func(w *Wrapper) { w.HxTarget("#result") }, `hx-target="#result"`},
		{"HxTrigger", func(w *Wrapper) { w.HxTrigger("keyup changed delay:500ms") }, `hx-trigger="keyup changed delay:500ms"`},
		{"HxBoost", func(w *Wrapper) { w.HxBoost(true) }, `hx-boost="true"`},
		{"HxConfirm", func(w *Wrapper) { w.HxConfirm("Are you sure?") }, `hx-confirm="Are you sure?"`},
		{"HxIndicator", func(w *Wrapper) { w.HxIndicator("#spinner") }, `hx-indicator="#spinner"`},
		{"HxPushURLTrue", func(w *Wrapper) { w.HxPushURL("true") }, `hx-push-url="true"`},
		{"HxPushURLPath", func(w *Wrapper) { w.HxPushURL("/custom/path") }, `hx-push-url="/custom/path"`},
		{"HxPushURLFalse", func(w *Wrapper) { w.HxPushURL("false") }, `hx-push-url="false"`},
		{"HxExt", func(w *Wrapper) { w.HxExt("ws") }, `hx-ext="ws"`},
		{"HxSelect", func(w *Wrapper) { w.HxSelect("#content") }, `hx-select="#content"`},
		{"HxSelectOOB", func(w *Wrapper) { w.HxSelectOOB("#sidebar") }, `hx-select-oob="#sidebar"`},
		{"HxSwapOOB", func(w *Wrapper) { w.HxSwapOOB("true") }, `hx-swap-oob="true"`},
		{"HxReplaceURL", func(w *Wrapper) { w.HxReplaceURL("/new-path") }, `hx-replace-url="/new-path"`},
		{"HxParams", func(w *Wrapper) { w.HxParams("not secret") }, `hx-params="not secret"`},
		{"HxPrompt", func(w *Wrapper) { w.HxPrompt("Enter a value") }, `hx-prompt="Enter a value"`},
		{"HxEncoding", func(w *Wrapper) { w.HxEncoding("multipart/form-data") }, `hx-encoding="multipart/form-data"`},
		{"HxPreserve", func(w *Wrapper) { w.HxPreserve() }, `hx-preserve="true"`},
		{"HxHistory", func(w *Wrapper) { w.HxHistory(false) }, `hx-history="false"`},
		{"HxHistoryElt", func(w *Wrapper) { w.HxHistoryElt() }, `hx-history-elt="true"`},
		{"HxDisable", func(w *Wrapper) { w.HxDisable() }, `hx-disable="true"`},
		{"HxDisabledElt", func(w *Wrapper) { w.HxDisabledElt("#submit-btn") }, `hx-disabled-elt="#submit-btn"`},
		{"HxDisinherit", func(w *Wrapper) { w.HxDisinherit("hx-target") }, `hx-disinherit="hx-target"`},
		{"HxInherit", func(w *Wrapper) { w.HxInherit("hx-target") }, `hx-inherit="hx-target"`},
		{"HxValidate", func(w *Wrapper) { w.HxValidate(true) }, `hx-validate="true"`},
		{"HxVars", func(w *Wrapper) { w.HxVars("myVar:computeValue()") }, `hx-vars="myVar:computeValue()"`},
		// Extensions.
		{"WSConnect", func(w *Wrapper) { w.WSConnect("/ws/chat") }, `ws-connect="/ws/chat"`},
		{"WSSend", func(w *Wrapper) { w.WSSend() }, `ws-send`},
		{"SSEConnect", func(w *Wrapper) { w.SSEConnect("/sse/events") }, `sse-connect="/sse/events"`},
		{"SSESwap", func(w *Wrapper) { w.SSESwap("newMessage") }, `sse-swap="newMessage"`},
		{"SSEClose", func(w *Wrapper) { w.SSEClose("streamEnd") }, `sse-close="streamEnd"`},
		{"Preload", func(w *Wrapper) { w.Preload("mouseover") }, `preload="mouseover"`},
		{"PreloadImagesOn", func(w *Wrapper) { w.PreloadImages(true) }, `preload-images="true"`},
		{"PreloadImagesOff", func(w *Wrapper) { w.PreloadImages(false) }, `preload-images="false"`},
		{"HxTargetError", func(w *Wrapper) { w.HxTargetError("#error-container") }, `hx-target-error="#error-container"`},
		{"HxTargetCode", func(w *Wrapper) { w.HxTargetCode(404, "#not-found") }, `hx-target-404="#not-found"`},
		{"HxTargetCodePattern", func(w *Wrapper) { w.HxTargetCodePattern("5*", "#server-error") }, `hx-target-5*="#server-error"`},
		{"HxHead", func(w *Wrapper) { w.HxHead("merge") }, `hx-head="merge"`},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			d := div.New()
			w := New(d)
			tc.set(w)

			html := string(d.RenderBytes())
			if !strings.Contains(html, tc.want) {
				t.Errorf("%s: want %q in %s", tc.name, tc.want, html)
			}
		})
	}
}

func TestHxVals(t *testing.T) {
	d := div.New()
	New(d).HxVals(`{"key":"value"}`)

	html := string(d.RenderBytes())
	if !strings.Contains(html, "hx-vals") {
		t.Errorf("HxVals() did not set attribute, got: %s", html)
	}

	// The attribute value may be HTML-escaped by the renderer.
	if !strings.Contains(html, "key") || !strings.Contains(html, "value") {
		t.Errorf("HxVals() value not preserved, got: %s", html)
	}
}

func TestHxHeaders(t *testing.T) {
	d := div.New()
	New(d).HxHeaders(`{"X-Custom":"value"}`)

	html := string(d.RenderBytes())
	if !strings.Contains(html, "hx-headers") {
		t.Errorf("HxHeaders() did not set attribute, got: %s", html)
	}

	if !strings.Contains(html, "X-Custom") {
		t.Errorf("HxHeaders() value not preserved, got: %s", html)
	}
}

func TestHxRequest(t *testing.T) {
	d := div.New()
	New(d).HxRequest(`{"timeout":5000}`)

	html := string(d.RenderBytes())
	if !strings.Contains(html, "hx-request") {
		t.Errorf("HxRequest() did not set attribute, got: %s", html)
	}

	if !strings.Contains(html, "timeout") {
		t.Errorf("HxRequest() value not preserved, got: %s", html)
	}
}

func TestHxInclude(t *testing.T) {
	d := div.New()
	New(d).HxInclude("[name='email']")

	html := string(d.RenderBytes())
	if !strings.Contains(html, "hx-include") {
		t.Errorf("HxInclude() did not set attribute, got: %s", html)
	}
}

func TestHxSync(t *testing.T) {
	tests := []struct {
		name     string
		strategy sync.Strategy
		want     string
	}{
		{"Drop", sync.Drop, `hx-sync="drop"`},
		{"Abort", sync.Abort, `hx-sync="abort"`},
		{"Replace", sync.Replace, `hx-sync="replace"`},
		{"QueueFirst", sync.QueueFirst, `hx-sync="queue first"`},
		{"QueueLast", sync.QueueLast, `hx-sync="queue last"`},
		{"QueueAll", sync.QueueAll, `hx-sync="queue all"`},
		{"Custom", sync.Custom("closest form:abort"), `hx-sync="closest form:abort"`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := div.New()
			New(d).HxSync(tt.strategy)
			html := string(d.RenderBytes())
			if !strings.Contains(html, tt.want) {
				t.Errorf("HxSync(%s) want %s in %s", tt.name, tt.want, html)
			}
		})
	}
}

func TestHxOn(t *testing.T) {
	d := div.New()
	New(d).HxOn("after-swap", "console.log('swapped')")

	html := string(d.RenderBytes())
	// htmx 2 event attributes use the double colon: hx-on::EVENT.
	if !strings.Contains(html, "hx-on::after-swap") {
		t.Errorf("HxOn() did not set attribute, got: %s", html)
	}

	if !strings.Contains(html, "console.log") {
		t.Errorf("HxOn() handler not preserved, got: %s", html)
	}
}

// The wrapper delegates both custom-attribute paths to the wrapped element:
// SetAttribute escapes its value at set-time, SetAttributeRaw stores it verbatim.
func TestSetAttributeRaw(t *testing.T) {
	d := div.New()
	w := New(d)

	w.SetAttribute("data-escaped", `"><script>`)
	w.SetAttributeRaw("data-raw", "a&amp;b")

	html := string(d.RenderBytes())
	if !strings.Contains(html, `data-escaped="&#34;&gt;&lt;script&gt;"`) {
		t.Errorf("SetAttribute did not escape, got: %s", html)
	}
	if !strings.Contains(html, `data-raw="a&amp;b"`) {
		t.Errorf("SetAttributeRaw did not pass the value through verbatim, got: %s", html)
	}
}
