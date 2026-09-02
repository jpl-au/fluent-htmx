package htmx

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jpl-au/fluent-htmx/htmx4/swap"
)

func TestResumeHeaderReaders(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	r.Header.Set(HXPTagHeader, "v7")
	r.Header.Set(LastEventIDHeader, "42")
	r.Header.Set(HXLastPartIDHeader, "3")

	if got := HxPTag(r); got != "v7" {
		t.Errorf("HxPTag = %q, want %q", got, "v7")
	}
	if got := HxLastEventID(r); got != "42" {
		t.Errorf("HxLastEventID = %q, want %q", got, "42")
	}
	if got := HxLastPartID(r); got != "3" {
		t.Errorf("HxLastPartID = %q, want %q", got, "3")
	}
}

func TestHxPTagUnchanged(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	r.Header.Set(HXPTagHeader, "v7")

	w := httptest.NewRecorder()
	if !HxPTagUnchanged(w, r, "v7") {
		t.Error("HxPTagUnchanged = false for a matching tag, want true")
	}
	if w.Code != http.StatusNotModified {
		t.Errorf("status = %d, want %d", w.Code, http.StatusNotModified)
	}

	w = httptest.NewRecorder()
	if HxPTagUnchanged(w, r, "v8") {
		t.Error("HxPTagUnchanged = true for a new tag, want false")
	}
	if got := w.Header().Get(HXPTagHeader); got != "v8" {
		t.Errorf("HX-PTag = %q, want %q", got, "v8")
	}
}

func TestHxLocationWith(t *testing.T) {
	w := httptest.NewRecorder()
	err := HxLocationWith(w, Location{Path: "/orders", Target: "#main", Swap: swap.InnerHTML, Replace: "true"})
	if err != nil {
		t.Fatalf("HxLocationWith() returned error: %v", err)
	}

	want := `{"path":"/orders","replace":"true","swap":"innerHTML","target":"#main"}`
	if got := w.Header().Get(HXLocationHeader); got != want {
		t.Errorf("HX-Location = %q, want %q", got, want)
	}

	w = httptest.NewRecorder()
	if err := HxLocationWith(w, Location{Path: "/orders", Push: "false"}); err != nil {
		t.Fatalf("HxLocationWith() returned error: %v", err)
	}
	if got, want := w.Header().Get(HXLocationHeader), `{"path":"/orders","push":"false"}`; got != want {
		t.Errorf("HX-Location = %q, want %q", got, want)
	}
}

// The prompt extension sends encodeURI(answer), so the reader must decode it.
func TestHxPromptDecodes(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	r.Header.Set(HXPromptHeader, "caf%C3%A9%20au%20lait+1")

	if got, want := HxPrompt(r), "café au lait+1"; got != want {
		t.Errorf("HxPrompt = %q, want %q", got, want)
	}

	r.Header.Set(HXPromptHeader, "100%")
	if got := HxPrompt(r); got != "100%" {
		t.Errorf("HxPrompt on bad encoding = %q, want the raw value", got)
	}
}
