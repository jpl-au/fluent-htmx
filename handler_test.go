package htmx

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandlerServesHTMXCore(t *testing.T) {
	h := Handler("/_htmx/")
	r := httptest.NewRequest(http.MethodGet, "/_htmx/htmx.min.js", nil)
	w := httptest.NewRecorder()
	h.ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}

	ct := w.Header().Get("Content-Type")
	if !strings.Contains(ct, "javascript") {
		t.Fatalf("expected javascript content-type, got %q", ct)
	}

	body := w.Body.String()
	if !strings.Contains(body, "htmx") {
		t.Fatal("response does not contain htmx code")
	}
}

func TestHandlerServesExtension(t *testing.T) {
	h := Handler("/_htmx/")
	r := httptest.NewRequest(http.MethodGet, "/_htmx/ext/ws.js", nil)
	w := httptest.NewRecorder()
	h.ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}

	body := w.Body.String()
	if !strings.Contains(body, "ws") {
		t.Fatal("response does not contain ws extension code")
	}
}

func TestHandlerSetsCacheHeaders(t *testing.T) {
	h := Handler("/_htmx/")
	r := httptest.NewRequest(http.MethodGet, "/_htmx/htmx.min.js", nil)
	w := httptest.NewRecorder()
	h.ServeHTTP(w, r)

	cc := w.Header().Get("Cache-Control")
	if !strings.Contains(cc, "immutable") {
		t.Fatalf("expected immutable cache-control, got %q", cc)
	}
}

func TestHandler404ForUnknownPath(t *testing.T) {
	h := Handler("/_htmx/")
	r := httptest.NewRequest(http.MethodGet, "/_htmx/nonexistent.js", nil)
	w := httptest.NewRecorder()
	h.ServeHTTP(w, r)

	if w.Code != http.StatusNotFound {
		t.Fatalf("expected 404, got %d", w.Code)
	}
}
