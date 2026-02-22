package htmx

import (
	"strings"
	"testing"

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

	html := string(w.Render())
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
	result := w.HxGet("/api/users").HxTarget("#result").HxSwap("innerHTML").HxTrigger("click")

	if result == nil {
		t.Fatal("method chaining returned nil")
	}

	html := string(d.Render())
	for _, attr := range []string{`hx-get="/api/users"`, `hx-target="#result"`, `hx-swap="innerHTML"`, `hx-trigger="click"`} {
		if !strings.Contains(html, attr) {
			t.Errorf("chained call missing %s, got: %s", attr, html)
		}
	}
}

func TestHxGet(t *testing.T) {
	d := div.New()
	w := New(d)

	w.HxGet("/api/users")

	html := string(d.Render())
	if !strings.Contains(html, `hx-get="/api/users"`) {
		t.Errorf("HxGet() did not set attribute correctly, got: %s", html)
	}
}

func TestHxPost(t *testing.T) {
	d := div.New()
	w := New(d)

	w.HxPost("/api/users")

	html := string(d.Render())
	if !strings.Contains(html, `hx-post="/api/users"`) {
		t.Errorf("HxPost() did not set attribute correctly, got: %s", html)
	}
}

func TestHxPut(t *testing.T) {
	d := div.New()
	w := New(d)

	w.HxPut("/api/users/1")

	html := string(d.Render())
	if !strings.Contains(html, `hx-put="/api/users/1"`) {
		t.Errorf("HxPut() did not set attribute correctly, got: %s", html)
	}
}

func TestHxPatch(t *testing.T) {
	d := div.New()
	w := New(d)

	w.HxPatch("/api/users/1")

	html := string(d.Render())
	if !strings.Contains(html, `hx-patch="/api/users/1"`) {
		t.Errorf("HxPatch() did not set attribute correctly, got: %s", html)
	}
}

func TestHxDelete(t *testing.T) {
	d := div.New()
	w := New(d)

	w.HxDelete("/api/users/1")

	html := string(d.Render())
	if !strings.Contains(html, `hx-delete="/api/users/1"`) {
		t.Errorf("HxDelete() did not set attribute correctly, got: %s", html)
	}
}

func TestHxSwap(t *testing.T) {
	d := div.New()
	w := New(d)

	w.HxSwap("outerHTML")

	html := string(d.Render())
	if !strings.Contains(html, `hx-swap="outerHTML"`) {
		t.Errorf("HxSwap() did not set attribute correctly, got: %s", html)
	}
}

func TestHxTarget(t *testing.T) {
	d := div.New()
	w := New(d)

	w.HxTarget("#result")

	html := string(d.Render())
	if !strings.Contains(html, `hx-target="#result"`) {
		t.Errorf("HxTarget() did not set attribute correctly, got: %s", html)
	}
}

func TestHxTrigger(t *testing.T) {
	d := div.New()
	w := New(d)

	w.HxTrigger("keyup changed delay:500ms")

	html := string(d.Render())
	if !strings.Contains(html, `hx-trigger="keyup changed delay:500ms"`) {
		t.Errorf("HxTrigger() did not set attribute correctly, got: %s", html)
	}
}

func TestHxBoost(t *testing.T) {
	d := div.New()
	w := New(d)

	w.HxBoost(true)

	html := string(d.Render())
	if !strings.Contains(html, `hx-boost="true"`) {
		t.Errorf("HxBoost(true) did not set attribute correctly, got: %s", html)
	}
}

func TestHxConfirm(t *testing.T) {
	d := div.New()
	w := New(d)

	w.HxConfirm("Are you sure?")

	html := string(d.Render())
	if !strings.Contains(html, `hx-confirm="Are you sure?"`) {
		t.Errorf("HxConfirm() did not set attribute correctly, got: %s", html)
	}
}

func TestHxVals(t *testing.T) {
	d := div.New()
	w := New(d)

	w.HxVals(`{"key":"value"}`)

	html := string(d.Render())
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
	w := New(d)

	w.HxHeaders(`{"X-Custom":"value"}`)

	html := string(d.Render())
	if !strings.Contains(html, "hx-headers") {
		t.Errorf("HxHeaders() did not set attribute, got: %s", html)
	}

	if !strings.Contains(html, "X-Custom") {
		t.Errorf("HxHeaders() value not preserved, got: %s", html)
	}
}

func TestHxIndicator(t *testing.T) {
	d := div.New()
	w := New(d)

	w.HxIndicator("#spinner")

	html := string(d.Render())
	if !strings.Contains(html, `hx-indicator="#spinner"`) {
		t.Errorf("HxIndicator() did not set attribute correctly, got: %s", html)
	}
}

func TestHxPushURL(t *testing.T) {
	d := div.New()
	w := New(d)

	w.HxPushURL("true")

	html := string(d.Render())
	if !strings.Contains(html, `hx-push-url="true"`) {
		t.Errorf("HxPushURL(\"true\") did not set attribute correctly, got: %s", html)
	}

	d2 := div.New()
	w2 := New(d2)
	w2.HxPushURL("/custom/path")

	html2 := string(d2.Render())
	if !strings.Contains(html2, `hx-push-url="/custom/path"`) {
		t.Errorf("HxPushURL(\"/custom/path\") did not set attribute correctly, got: %s", html2)
	}

	d3 := div.New()
	w3 := New(d3)
	w3.HxPushURL("false")

	html3 := string(d3.Render())
	if !strings.Contains(html3, `hx-push-url="false"`) {
		t.Errorf("HxPushURL(\"false\") did not set attribute correctly, got: %s", html3)
	}
}

func TestHxExt(t *testing.T) {
	d := div.New()
	w := New(d)

	w.HxExt("ws")

	html := string(d.Render())
	if !strings.Contains(html, `hx-ext="ws"`) {
		t.Errorf("HxExt() did not set attribute correctly, got: %s", html)
	}
}

func TestHxSelect(t *testing.T) {
	d := div.New()
	w := New(d)

	w.HxSelect("#content")

	html := string(d.Render())
	if !strings.Contains(html, `hx-select="#content"`) {
		t.Errorf("HxSelect() did not set attribute correctly, got: %s", html)
	}
}

func TestHxSelectOOB(t *testing.T) {
	d := div.New()
	w := New(d)

	w.HxSelectOOB("#sidebar")

	html := string(d.Render())
	if !strings.Contains(html, `hx-select-oob="#sidebar"`) {
		t.Errorf("HxSelectOOB() did not set attribute correctly, got: %s", html)
	}
}

func TestHxSwapOOB(t *testing.T) {
	d := div.New()
	w := New(d)

	w.HxSwapOOB("true")

	html := string(d.Render())
	if !strings.Contains(html, `hx-swap-oob="true"`) {
		t.Errorf("HxSwapOOB() did not set attribute correctly, got: %s", html)
	}
}

func TestHxReplaceUrl(t *testing.T) {
	d := div.New()
	w := New(d)

	w.HxReplaceURL("/new-path")

	html := string(d.Render())
	if !strings.Contains(html, `hx-replace-url="/new-path"`) {
		t.Errorf("HxReplaceURL() did not set attribute correctly, got: %s", html)
	}
}

func TestHxParams(t *testing.T) {
	d := div.New()
	w := New(d)

	w.HxParams("not secret")

	html := string(d.Render())
	if !strings.Contains(html, `hx-params="not secret"`) {
		t.Errorf("HxParams() did not set attribute correctly, got: %s", html)
	}
}

func TestHxInclude(t *testing.T) {
	d := div.New()
	w := New(d)

	w.HxInclude("[name='email']")

	html := string(d.Render())
	if !strings.Contains(html, "hx-include") {
		t.Errorf("HxInclude() did not set attribute, got: %s", html)
	}
}

func TestHxSync(t *testing.T) {
	d := div.New()
	w := New(d)

	w.HxSync("closest form:abort")

	html := string(d.Render())
	if !strings.Contains(html, `hx-sync="closest form:abort"`) {
		t.Errorf("HxSync() did not set attribute correctly, got: %s", html)
	}
}

func TestHxPrompt(t *testing.T) {
	d := div.New()
	w := New(d)

	w.HxPrompt("Enter a value")

	html := string(d.Render())
	if !strings.Contains(html, `hx-prompt="Enter a value"`) {
		t.Errorf("HxPrompt() did not set attribute correctly, got: %s", html)
	}
}

func TestHxEncoding(t *testing.T) {
	d := div.New()
	w := New(d)

	w.HxEncoding("multipart/form-data")

	html := string(d.Render())
	if !strings.Contains(html, `hx-encoding="multipart/form-data"`) {
		t.Errorf("HxEncoding() did not set attribute correctly, got: %s", html)
	}
}

func TestHxPreserve(t *testing.T) {
	d := div.New()
	w := New(d)

	w.HxPreserve(true)

	html := string(d.Render())
	if !strings.Contains(html, `hx-preserve="true"`) {
		t.Errorf("HxPreserve(true) did not set attribute correctly, got: %s", html)
	}
}

func TestHxHistory(t *testing.T) {
	d := div.New()
	w := New(d)

	w.HxHistory("false")

	html := string(d.Render())
	if !strings.Contains(html, `hx-history="false"`) {
		t.Errorf("HxHistory(\"false\") did not set attribute correctly, got: %s", html)
	}
}

func TestHxHistoryElt(t *testing.T) {
	d := div.New()
	w := New(d)

	w.HxHistoryElt()

	html := string(d.Render())
	if !strings.Contains(html, `hx-history-elt="true"`) {
		t.Errorf("HxHistoryElt() did not set attribute correctly, got: %s", html)
	}
}

func TestHxDisable(t *testing.T) {
	d := div.New()
	w := New(d)

	w.HxDisable()

	html := string(d.Render())
	if !strings.Contains(html, `hx-disable="true"`) {
		t.Errorf("HxDisable() did not set attribute correctly, got: %s", html)
	}
}

func TestHxDisabledElt(t *testing.T) {
	d := div.New()
	w := New(d)

	w.HxDisabledElt("#submit-btn")

	html := string(d.Render())
	if !strings.Contains(html, `hx-disabled-elt="#submit-btn"`) {
		t.Errorf("HxDisabledElt() did not set attribute correctly, got: %s", html)
	}
}

func TestHxDisinherit(t *testing.T) {
	d := div.New()
	w := New(d)

	w.HxDisinherit("hx-target")

	html := string(d.Render())
	if !strings.Contains(html, `hx-disinherit="hx-target"`) {
		t.Errorf("HxDisinherit() did not set attribute correctly, got: %s", html)
	}
}

func TestHxInherit(t *testing.T) {
	d := div.New()
	w := New(d)

	w.HxInherit("hx-target")

	html := string(d.Render())
	if !strings.Contains(html, `hx-inherit="hx-target"`) {
		t.Errorf("HxInherit() did not set attribute correctly, got: %s", html)
	}
}

func TestHxValidate(t *testing.T) {
	d := div.New()
	w := New(d)

	w.HxValidate(true)

	html := string(d.Render())
	if !strings.Contains(html, `hx-validate="true"`) {
		t.Errorf("HxValidate(true) did not set attribute correctly, got: %s", html)
	}
}

func TestHxRequest(t *testing.T) {
	d := div.New()
	w := New(d)

	w.HxRequest(`{"timeout":5000}`)

	html := string(d.Render())
	if !strings.Contains(html, "hx-request") {
		t.Errorf("HxRequest() did not set attribute, got: %s", html)
	}

	if !strings.Contains(html, "timeout") {
		t.Errorf("HxRequest() value not preserved, got: %s", html)
	}
}

func TestHxVars(t *testing.T) {
	d := div.New()
	w := New(d)

	w.HxVars("myVar:computeValue()")

	html := string(d.Render())
	if !strings.Contains(html, `hx-vars="myVar:computeValue()"`) {
		t.Errorf("HxVars() did not set attribute correctly, got: %s", html)
	}
}

func TestHxOn(t *testing.T) {
	d := div.New()
	w := New(d)

	w.HxOn("after-swap", "console.log('swapped')")

	html := string(d.Render())
	if !strings.Contains(html, "hx-on::after-swap") {
		t.Errorf("HxOn() did not set attribute, got: %s", html)
	}

	if !strings.Contains(html, "console.log") {
		t.Errorf("HxOn() handler not preserved, got: %s", html)
	}
}

// Extension tests: WebSocket

func TestWsConnect(t *testing.T) {
	d := div.New()
	w := New(d)

	w.WsConnect("/ws/chat")

	html := string(d.Render())
	if !strings.Contains(html, `ws-connect="/ws/chat"`) {
		t.Errorf("WsConnect() did not set attribute correctly, got: %s", html)
	}
}

func TestWsSend(t *testing.T) {
	d := div.New()
	w := New(d)

	w.WsSend()

	html := string(d.Render())
	if !strings.Contains(html, "ws-send") {
		t.Errorf("WsSend() did not set attribute, got: %s", html)
	}
}

// Extension tests: SSE

func TestSSEConnect(t *testing.T) {
	d := div.New()
	w := New(d)

	w.SSEConnect("/sse/events")

	html := string(d.Render())
	if !strings.Contains(html, `sse-connect="/sse/events"`) {
		t.Errorf("SSEConnect() did not set attribute correctly, got: %s", html)
	}
}

func TestSSESwap(t *testing.T) {
	d := div.New()
	w := New(d)

	w.SSESwap("newMessage")

	html := string(d.Render())
	if !strings.Contains(html, `sse-swap="newMessage"`) {
		t.Errorf("SSESwap() did not set attribute correctly, got: %s", html)
	}
}

func TestSSEClose(t *testing.T) {
	d := div.New()
	w := New(d)

	w.SSEClose("streamEnd")

	html := string(d.Render())
	if !strings.Contains(html, `sse-close="streamEnd"`) {
		t.Errorf("SSEClose() did not set attribute correctly, got: %s", html)
	}
}

// Extension tests: Preload

func TestPreload(t *testing.T) {
	d := div.New()
	w := New(d)

	w.Preload("mouseover")

	html := string(d.Render())
	if !strings.Contains(html, `preload="mouseover"`) {
		t.Errorf("Preload() did not set attribute correctly, got: %s", html)
	}
}

func TestPreloadImages(t *testing.T) {
	d := div.New()
	w := New(d)

	w.PreloadImages(true)

	html := string(d.Render())
	if !strings.Contains(html, `preload-images="true"`) {
		t.Errorf("PreloadImages(true) did not set attribute correctly, got: %s", html)
	}

	d2 := div.New()
	w2 := New(d2)

	w2.PreloadImages(false)

	html2 := string(d2.Render())
	if !strings.Contains(html2, `preload-images="false"`) {
		t.Errorf("PreloadImages(false) did not set attribute correctly, got: %s", html2)
	}
}

// Extension tests: Response Targets

func TestHxTargetError(t *testing.T) {
	d := div.New()
	w := New(d)

	w.HxTargetError("#error-container")

	html := string(d.Render())
	if !strings.Contains(html, `hx-target-error="#error-container"`) {
		t.Errorf("HxTargetError() did not set attribute correctly, got: %s", html)
	}
}

func TestHxTargetCode(t *testing.T) {
	d := div.New()
	w := New(d)

	w.HxTargetCode(404, "#not-found")

	html := string(d.Render())
	if !strings.Contains(html, `hx-target-404="#not-found"`) {
		t.Errorf("HxTargetCode() did not set attribute correctly, got: %s", html)
	}
}

func TestHxTargetCodePattern(t *testing.T) {
	d := div.New()
	w := New(d)

	w.HxTargetCodePattern("5*", "#server-error")

	html := string(d.Render())
	if !strings.Contains(html, `hx-target-5*="#server-error"`) {
		t.Errorf("HxTargetCodePattern() did not set attribute correctly, got: %s", html)
	}
}

// Extension tests: Head Support

func TestHxHead(t *testing.T) {
	d := div.New()
	w := New(d)

	w.HxHead("merge")

	html := string(d.Render())
	if !strings.Contains(html, `hx-head="merge"`) {
		t.Errorf("HxHead() did not set attribute correctly, got: %s", html)
	}
}
