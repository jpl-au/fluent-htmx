package htmx

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/jpl-au/fluent/html5/div"
	"github.com/jpl-au/fluent/text"
)

func TestServerHxRequest(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	if HxRequest(r) {
		t.Error("HxRequest() should return false without HX-Request header")
	}

	r.Header.Set("HX-Request", "true")
	if !HxRequest(r) {
		t.Error("HxRequest() should return true with HX-Request: true header")
	}
}

func TestHandle(t *testing.T) {
	called := false
	r := httptest.NewRequest(http.MethodGet, "/", nil)

	if Handle(r, func() { called = true }) {
		t.Error("Handle() should return false for non-HTMX request")
	}

	if called {
		t.Error("Handle() should not call fn for non-HTMX request")
	}

	r.Header.Set("HX-Request", "true")
	if !Handle(r, func() { called = true }) {
		t.Error("Handle() should return true for HTMX request")
	}

	if !called {
		t.Error("Handle() should call fn for HTMX request")
	}
}

func TestServerHxBoosted(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	if HxBoosted(r) {
		t.Error("HxBoosted() should return false without header")
	}

	r.Header.Set("HX-Boosted", "true")
	if !HxBoosted(r) {
		t.Error("HxBoosted() should return true with header")
	}
}

func TestServerHxCurrentURL(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	r.Header.Set("HX-Current-URL", "https://example.com/page")

	if got := HxCurrentURL(r); got != "https://example.com/page" {
		t.Errorf("HxCurrentURL() = %q, want %q", got, "https://example.com/page")
	}
}

func TestServerHxHistoryRestoreRequest(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	if HxHistoryRestoreRequest(r) {
		t.Error("HxHistoryRestoreRequest() should return false without header")
	}

	r.Header.Set("HX-History-Restore-Request", "true")
	if !HxHistoryRestoreRequest(r) {
		t.Error("HxHistoryRestoreRequest() should return true with header")
	}
}

func TestServerHxPrompt(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	if got := HxPrompt(r); got != "" {
		t.Errorf("HxPrompt() without header = %q, want empty string", got)
	}

	r.Header.Set("HX-Prompt", "user input here")
	if got := HxPrompt(r); got != "user input here" {
		t.Errorf("HxPrompt() = %q, want %q", got, "user input here")
	}
}

func TestServerHxTarget(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	r.Header.Set("HX-Target", "result-div")

	if got := HxTarget(r); got != "result-div" {
		t.Errorf("HxTarget() = %q, want %q", got, "result-div")
	}
}

func TestServerHxTriggerName(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	r.Header.Set("HX-Trigger-Name", "search-form")

	if got := HxTriggerName(r); got != "search-form" {
		t.Errorf("HxTriggerName() = %q, want %q", got, "search-form")
	}
}

func TestServerHxTrigger(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	r.Header.Set("HX-Trigger", "btn-submit")

	if got := HxTrigger(r); got != "btn-submit" {
		t.Errorf("HxTrigger() = %q, want %q", got, "btn-submit")
	}
}

func TestHxRedirectHTMX(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	r.Header.Set("HX-Request", "true")

	HxRedirect(w, r, "/dashboard", http.StatusSeeOther)

	if got := w.Header().Get("HX-Redirect"); got != "/dashboard" {
		t.Errorf("HX-Redirect header = %q, want %q", got, "/dashboard")
	}

	if w.Code != http.StatusOK {
		t.Errorf("status code = %d, want %d (HTMX processes redirects client-side)", w.Code, http.StatusOK)
	}
}

func TestHxRedirectStandard(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)

	HxRedirect(w, r, "/dashboard", http.StatusSeeOther)

	if w.Code != http.StatusSeeOther {
		t.Errorf("status code = %d, want %d", w.Code, http.StatusSeeOther)
	}
}

func TestServerHxPushURL(t *testing.T) {
	w := httptest.NewRecorder()
	HxPushURL(w, "/new-url")

	if got := w.Header().Get("HX-Push-Url"); got != "/new-url" {
		t.Errorf("HX-Push-Url header = %q, want %q", got, "/new-url")
	}
}

func TestHxLocationHeader(t *testing.T) {
	w := httptest.NewRecorder()
	HxLocation(w, "/new-page")

	if got := w.Header().Get("HX-Location"); got != "/new-page" {
		t.Errorf("HX-Location header = %q, want %q", got, "/new-page")
	}
}

func TestServerHxReplaceURL(t *testing.T) {
	w := httptest.NewRecorder()
	HxReplaceURL(w, "/replaced")

	if got := w.Header().Get("HX-Replace-Url"); got != "/replaced" {
		t.Errorf("HX-Replace-Url header = %q, want %q", got, "/replaced")
	}
}

func TestHxRefreshHeader(t *testing.T) {
	w := httptest.NewRecorder()
	HxRefresh(w)

	if got := w.Header().Get("HX-Refresh"); got != "true" {
		t.Errorf("HX-Refresh header = %q, want %q", got, "true")
	}
}

func TestHxRetargetHeader(t *testing.T) {
	w := httptest.NewRecorder()
	HxRetarget(w, "#error-div")

	if got := w.Header().Get("HX-Retarget"); got != "#error-div" {
		t.Errorf("HX-Retarget header = %q, want %q", got, "#error-div")
	}
}

func TestHxReswapHeader(t *testing.T) {
	w := httptest.NewRecorder()
	HxReswap(w, SwapOuterHTML)

	if got := w.Header().Get("HX-Reswap"); got != "outerHTML" {
		t.Errorf("HX-Reswap header = %q, want %q", got, "outerHTML")
	}
}

func TestHxReselectHeader(t *testing.T) {
	w := httptest.NewRecorder()
	HxReselect(w, "#main-content")

	if got := w.Header().Get("HX-Reselect"); got != "#main-content" {
		t.Errorf("HX-Reselect header = %q, want %q", got, "#main-content")
	}
}

func TestTriggerBuilderSimpleEvents(t *testing.T) {
	w := httptest.NewRecorder()
	tb := NewTrigger(w)

	err := tb.AddTrigger("showMessage", nil).
		AddTrigger("refreshList", nil).
		Write(text.RawText("<div>ok</div>"), http.StatusOK)

	if err != nil {
		t.Fatalf("Write() returned error: %v", err)
	}

	got := w.Header().Get("HX-Trigger")
	if got != "showMessage,refreshList" {
		t.Errorf("HX-Trigger header = %q, want %q", got, "showMessage,refreshList")
	}

	if w.Code != http.StatusOK {
		t.Errorf("status code = %d, want %d", w.Code, http.StatusOK)
	}

	if w.Body.String() != "<div>ok</div>" {
		t.Errorf("body = %q, want %q", w.Body.String(), "<div>ok</div>")
	}
}

func TestTriggerBuilderWriteNode(t *testing.T) {
	w := httptest.NewRecorder()
	tb := NewTrigger(w)

	n := div.New().Text("fluent node")
	err := tb.AddTrigger("event", nil).Write(n, http.StatusOK)

	if err != nil {
		t.Fatalf("Write() returned error: %v", err)
	}

	if !strings.Contains(w.Body.String(), "fluent node") {
		t.Errorf("body = %q, want it to contain %q", w.Body.String(), "fluent node")
	}
}

func TestTriggerBuilderDetailedEvents(t *testing.T) {
	w := httptest.NewRecorder()
	tb := NewTrigger(w)

	err := tb.AddTrigger("showMessage", map[string]string{"level": "info", "message": "Saved"}).
		Write(text.RawText(""), http.StatusOK)

	if err != nil {
		t.Fatalf("Write() returned error: %v", err)
	}

	got := w.Header().Get("HX-Trigger")
	if got == "" {
		t.Fatal("HX-Trigger header is empty")
	}

	if !strings.Contains(got, "showMessage") {
		t.Errorf("HX-Trigger header %q missing event name", got)
	}

	if !strings.Contains(got, "info") {
		t.Errorf("HX-Trigger header %q missing detail value", got)
	}
}

func TestTriggerBuilderAfterSwap(t *testing.T) {
	w := httptest.NewRecorder()
	tb := NewTrigger(w)

	err := tb.AddTriggerAfterSwap("scrollTo", nil).
		Write(text.RawText(""), http.StatusOK)

	if err != nil {
		t.Fatalf("Write() returned error: %v", err)
	}

	if got := w.Header().Get("HX-Trigger-After-Swap"); got != "scrollTo" {
		t.Errorf("HX-Trigger-After-Swap header = %q, want %q", got, "scrollTo")
	}
}

func TestTriggerBuilderAfterSettle(t *testing.T) {
	w := httptest.NewRecorder()
	tb := NewTrigger(w)

	err := tb.AddTriggerAfterSettle("initTooltips", nil).
		Write(text.RawText(""), http.StatusOK)

	if err != nil {
		t.Fatalf("Write() returned error: %v", err)
	}

	if got := w.Header().Get("HX-Trigger-After-Settle"); got != "initTooltips" {
		t.Errorf("HX-Trigger-After-Settle header = %q, want %q", got, "initTooltips")
	}
}

func TestTriggerBuilderSetsContentType(t *testing.T) {
	w := httptest.NewRecorder()
	tb := NewTrigger(w)

	err := tb.AddTrigger("event", nil).Write(text.RawText("<p>test</p>"), http.StatusOK)
	if err != nil {
		t.Fatalf("Write() returned error: %v", err)
	}

	if ct := w.Header().Get("Content-Type"); ct != "text/html; charset=utf-8" {
		t.Errorf("Content-Type = %q, want %q", ct, "text/html; charset=utf-8")
	}
}

func TestResponseHelper(t *testing.T) {
	w := httptest.NewRecorder()
	err := Response(w, text.RawText("<p>Hello</p>"), http.StatusOK)
	if err != nil {
		t.Fatalf("Response() returned error: %v", err)
	}

	if ct := w.Header().Get("Content-Type"); ct != "text/html; charset=utf-8" {
		t.Errorf("Content-Type = %q, want %q", ct, "text/html; charset=utf-8")
	}

	if w.Code != http.StatusOK {
		t.Errorf("status code = %d, want %d", w.Code, http.StatusOK)
	}

	if w.Body.String() != "<p>Hello</p>" {
		t.Errorf("body = %q, want %q", w.Body.String(), "<p>Hello</p>")
	}
}

func TestResponseNodeHelper(t *testing.T) {
	w := httptest.NewRecorder()
	n := div.New().Text("hello")
	err := Response(w, n, http.StatusOK)
	if err != nil {
		t.Fatalf("Response() returned error: %v", err)
	}

	if !strings.Contains(w.Body.String(), "hello") {
		t.Errorf("body = %q, want it to contain %q", w.Body.String(), "hello")
	}
}
