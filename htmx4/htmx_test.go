package htmx

import (
	"strings"
	"testing"

	"github.com/jpl-au/fluent-htmx/htmx4/swap"
	"github.com/jpl-au/fluent-htmx/htmx4/sync"
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
		{"HxAction", func(w *Wrapper) { w.HxAction("/api/items") }, `hx-action="/api/items"`},
		{"HxMethod", func(w *Wrapper) { w.HxMethod("post") }, `hx-method="post"`},
		{"HxSwap", func(w *Wrapper) { w.HxSwap(swap.OuterHTML) }, `hx-swap="outerHTML"`},
		{"HxSwapMorph", func(w *Wrapper) { w.HxSwap(swap.OuterMorph) }, `hx-swap="outerMorph"`},
		{"HxSwapCustom", func(w *Wrapper) { w.HxSwap(swap.Custom("innerHTML show:top")) }, `hx-swap="innerHTML show:top"`},
		{"HxTarget", func(w *Wrapper) { w.HxTarget("#result") }, `hx-target="#result"`},
		{"HxTrigger", func(w *Wrapper) { w.HxTrigger("keyup delay:500ms") }, `hx-trigger="keyup delay:500ms"`},
		{"HxBoost", func(w *Wrapper) { w.HxBoost(true) }, `hx-boost="true"`},
		{"HxConfirm", func(w *Wrapper) { w.HxConfirm("Sure?") }, `hx-confirm="Sure?"`},
		{"HxIndicator", func(w *Wrapper) { w.HxIndicator("#spinner") }, `hx-indicator="#spinner"`},
		{"HxPushURL", func(w *Wrapper) { w.HxPushURL("/custom/path") }, `hx-push-url="/custom/path"`},
		{"HxSelect", func(w *Wrapper) { w.HxSelect("#content") }, `hx-select="#content"`},
		{"HxSelectOOB", func(w *Wrapper) { w.HxSelectOOB("#sidebar") }, `hx-select-oob="#sidebar"`},
		{"HxSwapOOB", func(w *Wrapper) { w.HxSwapOOB("true") }, `hx-swap-oob="true"`},
		{"HxReplaceURL", func(w *Wrapper) { w.HxReplaceURL("/new") }, `hx-replace-url="/new"`},
		{"HxEncoding", func(w *Wrapper) { w.HxEncoding("multipart/form-data") }, `hx-encoding="multipart/form-data"`},
		{"HxPreserve", func(w *Wrapper) { w.HxPreserve() }, `hx-preserve="true"`},
		{"HxHistoryElt", func(w *Wrapper) { w.HxHistoryElt() }, `hx-history-elt="true"`},
		{"HxValidate", func(w *Wrapper) { w.HxValidate(true) }, `hx-validate="true"`},
		// Additional attributes.
		{"HxIgnore", func(w *Wrapper) { w.HxIgnore() }, `hx-ignore="true"`},
		{"HxDisable", func(w *Wrapper) { w.HxDisable("#submit-btn") }, `hx-disable="#submit-btn"`},
		{"HxConfig", func(w *Wrapper) { w.HxConfig("timeout:5000") }, `hx-config="timeout:5000"`},
		{"HxStatus", func(w *Wrapper) { w.HxStatus("422", "swap:none") }, `hx-status:422="swap:none"`},
		// Extensions.
		{"WsConnect", func(w *Wrapper) { w.WsConnect("/ws/chat") }, `hx-ws:connect="/ws/chat"`},
		{"WsSend", func(w *Wrapper) { w.WsSend() }, `hx-ws:send`},
		{"SSEConnect", func(w *Wrapper) { w.SSEConnect("/sse/events") }, `hx-sse:connect="/sse/events"`},
		{"SSEClose", func(w *Wrapper) { w.SSEClose("streamEnd") }, `hx-sse:close="streamEnd"`},
		{"Preload", func(w *Wrapper) { w.Preload("mouseover") }, `hx-preload="mouseover"`},
		{"HxHead", func(w *Wrapper) { w.HxHead("merge") }, `hx-head="merge"`},
		{"HxOptimistic", func(w *Wrapper) { w.HxOptimistic("#pending") }, `hx-optimistic="#pending"`},
		{"HxTargets", func(w *Wrapper) { w.HxTargets(".card") }, `hx-targets=".card"`},
		{"HxLive", func(w *Wrapper) { w.HxLive("count + 1") }, `hx-live="count + 1"`},
		{"HxBrowserIndicator", func(w *Wrapper) { w.HxBrowserIndicator(true) }, `hx-browser-indicator="true"`},
		{"HxPtag", func(w *Wrapper) { w.HxPtag("v42") }, `hx-ptag="v42"`},
		{"HxHistory", func(w *Wrapper) { w.HxHistory(false) }, `hx-history="false"`},
		{"HxNonce", func(w *Wrapper) { w.HxNonce("abc123") }, `hx-nonce="abc123"`},
		{"HxPrompt", func(w *Wrapper) { w.HxPrompt("Enter name") }, `hx-prompt="Enter name"`},
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
	if !strings.Contains(html, "hx-vals") || !strings.Contains(html, "key") {
		t.Errorf("HxVals() not rendered correctly, got: %s", html)
	}
}

func TestHxHeaders(t *testing.T) {
	d := div.New()
	New(d).HxHeaders(`{"X-Custom":"value"}`)

	html := string(d.RenderBytes())
	if !strings.Contains(html, "hx-headers") || !strings.Contains(html, "X-Custom") {
		t.Errorf("HxHeaders() not rendered correctly, got: %s", html)
	}
}

func TestHxInclude(t *testing.T) {
	d := div.New()
	New(d).HxInclude("closest form")

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
		{"QueueAll", sync.QueueAll, `hx-sync="queue all"`},
		{"Custom", sync.Custom("this:queue all"), `hx-sync="this:queue all"`},
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
	New(d).HxOn("click", "alert('hi')")

	html := string(d.RenderBytes())
	// hx-on uses a single colon: hx-on:EVENT.
	if !strings.Contains(html, "hx-on:click") {
		t.Errorf("HxOn() did not set hx-on:click, got: %s", html)
	}

	if !strings.Contains(html, "alert") {
		t.Errorf("HxOn() handler not preserved, got: %s", html)
	}
}

func TestInheritModifier(t *testing.T) {
	cases := []setterCase{
		{"Inherited", func(w *Wrapper) { w.HxConfirm("Sure?", Inherited) }, `hx-confirm:inherited="Sure?"`},
		{"InheritedAppend", func(w *Wrapper) { w.HxInclude("#extra", InheritedAppend) }, `hx-include:inherited:append="#extra"`},
		{"Bool setter", func(w *Wrapper) { w.HxBoost(true, Inherited) }, `hx-boost:inherited="true"`},
		{"Typed setter", func(w *Wrapper) { w.HxSwap(swap.OuterHTML, Inherited) }, `hx-swap:inherited="outerHTML"`},
		{"No modifier is unchanged", func(w *Wrapper) { w.HxTarget("#main") }, `hx-target="#main"`},
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
