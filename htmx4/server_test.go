package htmx

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/jpl-au/fluent-htmx/htmx4/swap"
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

func TestServerHxTarget(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	r.Header.Set("HX-Target", "div#result")

	if got := HxTarget(r); got != "div#result" {
		t.Errorf("HxTarget() = %q, want %q", got, "div#result")
	}
}

func TestServerHxSource(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	r.Header.Set("HX-Source", "button#submit")

	if got := HxSource(r); got != "button#submit" {
		t.Errorf("HxSource() = %q, want %q", got, "button#submit")
	}
}

func TestServerHxRequestType(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	if got := HxRequestType(r); got != "" {
		t.Errorf("HxRequestType() without header = %q, want empty", got)
	}

	r.Header.Set("HX-Request-Type", "partial")
	if got := HxRequestType(r); got != "partial" {
		t.Errorf("HxRequestType() = %q, want %q", got, "partial")
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

func TestServerResponseHeaders(t *testing.T) {
	tests := []struct {
		name   string
		set    func(http.ResponseWriter)
		header string
		want   string
	}{
		{"HxPushURL", func(w http.ResponseWriter) { HxPushURL(w, "/new-url") }, "HX-Push-Url", "/new-url"},
		{"HxLocation", func(w http.ResponseWriter) { HxLocation(w, "/new-page") }, "HX-Location", "/new-page"},
		{"HxReplaceURL", func(w http.ResponseWriter) { HxReplaceURL(w, "/replaced") }, "HX-Replace-Url", "/replaced"},
		{"HxRefresh", func(w http.ResponseWriter) { HxRefresh(w) }, "HX-Refresh", "true"},
		{"HxRetarget", func(w http.ResponseWriter) { HxRetarget(w, "#error-div") }, "HX-Retarget", "#error-div"},
		{"HxReswap", func(w http.ResponseWriter) { HxReswap(w, swap.OuterHTML) }, "HX-Reswap", "outerHTML"},
		{"HxReselect", func(w http.ResponseWriter) { HxReselect(w, "#main") }, "HX-Reselect", "#main"},
		{"HxDownload", func(w http.ResponseWriter) { HxDownload(w, "/files/report.pdf") }, "HX-Download", "/files/report.pdf"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			tt.set(w)
			if got := w.Header().Get(tt.header); got != tt.want {
				t.Errorf("%s header = %q, want %q", tt.header, got, tt.want)
			}
		})
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
	if !strings.Contains(got, "showMessage") || !strings.Contains(got, "info") {
		t.Errorf("HX-Trigger header %q missing event name or detail", got)
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
