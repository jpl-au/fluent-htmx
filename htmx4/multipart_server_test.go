package htmx

import (
	"io"
	"mime"
	stdmultipart "mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/jpl-au/fluent-htmx/htmx4/multipart"
	"github.com/jpl-au/fluent-htmx/htmx4/swap"
	"github.com/jpl-au/fluent/text"
)

const partContentType = "text/html; charset=utf-8"

func TestNewMultipart(t *testing.T) {
	w := httptest.NewRecorder()

	mw, err := NewMultipart(w, multipart.Mixed)
	if err != nil {
		t.Fatalf("NewMultipart() returned error: %v", err)
	}

	if mw == nil {
		t.Fatal("NewMultipart() returned nil writer")
	}

	mediaType, params, err := mime.ParseMediaType(w.Header().Get("Content-Type"))
	if err != nil {
		t.Fatalf("Content-Type not parseable: %v", err)
	}

	if mediaType != string(multipart.Mixed) {
		t.Errorf("media type = %q, want %q", mediaType, multipart.Mixed)
	}

	if params["boundary"] == "" {
		t.Error("Content-Type carries no boundary")
	}

	if got := w.Header().Get("Cache-Control"); got != "no-cache" {
		t.Errorf("Cache-Control = %q, want %q", got, "no-cache")
	}
}

func TestNewMultipartNoFlusher(t *testing.T) {
	w := &nonFlusher{header: http.Header{}}

	_, err := NewMultipart(w, multipart.Mixed)
	if err == nil {
		t.Error("NewMultipart() should return error when ResponseWriter lacks Flusher")
	}
}

func TestMultipartWriteParts(t *testing.T) {
	w := httptest.NewRecorder()

	mw, err := NewMultipart(w, multipart.Mixed)
	if err != nil {
		t.Fatalf("NewMultipart() returned error: %v", err)
	}

	if err := mw.WritePart(text.RawText("<p>one</p>"), PartTarget("#a"), PartSwap(swap.InnerHTML)); err != nil {
		t.Fatalf("WritePart(one) returned error: %v", err)
	}

	if err := mw.WritePart(text.RawText("<p>two</p>"), PartTarget("#b")); err != nil {
		t.Fatalf("WritePart(two) returned error: %v", err)
	}

	if err := mw.Close(); err != nil {
		t.Fatalf("Close() returned error: %v", err)
	}

	_, params, err := mime.ParseMediaType(w.Header().Get("Content-Type"))
	if err != nil {
		t.Fatalf("Content-Type not parseable: %v", err)
	}

	r := stdmultipart.NewReader(strings.NewReader(w.Body.String()), params["boundary"])

	// First part: target #a, innerHTML swap.
	p1, err := r.NextPart()
	if err != nil {
		t.Fatalf("NextPart(1) returned error: %v", err)
	}

	// Every part must declare text/html, the header the client parser relies on to find a
	// header block. Removing the auto-set header would break the JS parser while still passing
	// Go's tolerant reader, so this invariant is asserted explicitly.
	if got := p1.Header.Get("Content-Type"); got != partContentType {
		t.Errorf("part 1 Content-Type = %q, want %q", got, partContentType)
	}

	if got := p1.Header.Get(HXRetargetHeader); got != "#a" {
		t.Errorf("part 1 HX-Retarget = %q, want %q", got, "#a")
	}

	if got := p1.Header.Get(HXReswapHeader); got != string(swap.InnerHTML) {
		t.Errorf("part 1 HX-Reswap = %q, want %q", got, swap.InnerHTML)
	}

	if body, _ := io.ReadAll(p1); string(body) != "<p>one</p>" {
		t.Errorf("part 1 body = %q, want %q", body, "<p>one</p>")
	}

	// Second part: target #b, no explicit swap style.
	p2, err := r.NextPart()
	if err != nil {
		t.Fatalf("NextPart(2) returned error: %v", err)
	}

	if got := p2.Header.Get("Content-Type"); got != partContentType {
		t.Errorf("part 2 Content-Type = %q, want %q", got, partContentType)
	}

	if got := p2.Header.Get(HXRetargetHeader); got != "#b" {
		t.Errorf("part 2 HX-Retarget = %q, want %q", got, "#b")
	}

	if got := p2.Header.Get(HXReswapHeader); got != "" {
		t.Errorf("part 2 HX-Reswap = %q, want empty", got)
	}

	if body, _ := io.ReadAll(p2); string(body) != "<p>two</p>" {
		t.Errorf("part 2 body = %q, want %q", body, "<p>two</p>")
	}

	// Close must have written the terminating boundary.
	if _, err := r.NextPart(); err != io.EOF {
		t.Errorf("NextPart() after last part = %v, want io.EOF", err)
	}
}

func TestNewMultipartParallel(t *testing.T) {
	w := httptest.NewRecorder()

	if _, err := NewMultipart(w, multipart.Parallel); err != nil {
		t.Fatalf("NewMultipart() returned error: %v", err)
	}

	mediaType, _, err := mime.ParseMediaType(w.Header().Get("Content-Type"))
	if err != nil {
		t.Fatalf("Content-Type not parseable: %v", err)
	}

	if mediaType != string(multipart.Parallel) {
		t.Errorf("media type = %q, want %q", mediaType, multipart.Parallel)
	}
}

func TestMultipartPartOptions(t *testing.T) {
	w := httptest.NewRecorder()

	mw, err := NewMultipart(w, multipart.Mixed)
	if err != nil {
		t.Fatalf("NewMultipart() returned error: %v", err)
	}

	// Every PartOption on a single part, so each HX-* header is exercised.
	if err := mw.WritePart(text.RawText("<p>x</p>"),
		PartTarget("#t"), PartSwap(swap.OuterHTML), PartSelect("#frag"), PartTrigger("saved"), PartID("7"),
		PartRefresh(), PartRedirect("/next"), PartLocation("/there")); err != nil {
		t.Fatalf("WritePart() returned error: %v", err)
	}

	if err := mw.Close(); err != nil {
		t.Fatalf("Close() returned error: %v", err)
	}

	_, params, err := mime.ParseMediaType(w.Header().Get("Content-Type"))
	if err != nil {
		t.Fatalf("Content-Type not parseable: %v", err)
	}

	part, err := stdmultipart.NewReader(strings.NewReader(w.Body.String()), params["boundary"]).NextPart()
	if err != nil {
		t.Fatalf("NextPart() returned error: %v", err)
	}

	for _, c := range []struct{ header, want string }{
		{"Content-Type", partContentType},
		{HXRetargetHeader, "#t"},
		{HXReswapHeader, string(swap.OuterHTML)},
		{HXReselectHeader, "#frag"},
		{HXTriggerHeader, "saved"},
		{HXPartIDHeader, "7"},
		{HXRefreshHeader, "true"},
		{HXRedirectHeader, "/next"},
		{HXLocationHeader, "/there"},
	} {
		if got := part.Header.Get(c.header); got != c.want {
			t.Errorf("%s = %q, want %q", c.header, got, c.want)
		}
	}

	if body, _ := io.ReadAll(part); string(body) != "<p>x</p>" {
		t.Errorf("body = %q, want %q", body, "<p>x</p>")
	}
}

func TestMultipartWritePartNil(t *testing.T) {
	w := httptest.NewRecorder()

	mw, err := NewMultipart(w, multipart.Mixed)
	if err != nil {
		t.Fatalf("NewMultipart() returned error: %v", err)
	}

	// A nil node writes a header-only part. Pairing it with swap.None makes it a true no-op
	// swap that only fires the trigger; without swap.None the extension would swap empty content
	// into the default target.
	if err := mw.WritePart(nil, PartSwap(swap.None), PartTrigger("refresh")); err != nil {
		t.Fatalf("WritePart(nil) returned error: %v", err)
	}

	if err := mw.Close(); err != nil {
		t.Fatalf("Close() returned error: %v", err)
	}

	_, params, err := mime.ParseMediaType(w.Header().Get("Content-Type"))
	if err != nil {
		t.Fatalf("Content-Type not parseable: %v", err)
	}

	part, err := stdmultipart.NewReader(strings.NewReader(w.Body.String()), params["boundary"]).NextPart()
	if err != nil {
		t.Fatalf("NextPart() returned error: %v", err)
	}

	if got := part.Header.Get("Content-Type"); got != partContentType {
		t.Errorf("Content-Type = %q, want %q", got, partContentType)
	}

	if got := part.Header.Get(HXReswapHeader); got != string(swap.None) {
		t.Errorf("HX-Reswap = %q, want %q", got, swap.None)
	}

	if got := part.Header.Get(HXTriggerHeader); got != "refresh" {
		t.Errorf("HX-Trigger = %q, want %q", got, "refresh")
	}

	if body, _ := io.ReadAll(part); len(body) != 0 {
		t.Errorf("body = %q, want empty", body)
	}
}

func TestMultipartWritePartNoOptions(t *testing.T) {
	w := httptest.NewRecorder()

	mw, err := NewMultipart(w, multipart.Mixed)
	if err != nil {
		t.Fatalf("NewMultipart() returned error: %v", err)
	}

	// A part with no options: the auto-set Content-Type is its only header, the exact case the
	// parser invariant protects.
	if err := mw.WritePart(text.RawText("<p>bare</p>")); err != nil {
		t.Fatalf("WritePart() returned error: %v", err)
	}

	if err := mw.Close(); err != nil {
		t.Fatalf("Close() returned error: %v", err)
	}

	_, params, err := mime.ParseMediaType(w.Header().Get("Content-Type"))
	if err != nil {
		t.Fatalf("Content-Type not parseable: %v", err)
	}

	part, err := stdmultipart.NewReader(strings.NewReader(w.Body.String()), params["boundary"]).NextPart()
	if err != nil {
		t.Fatalf("NextPart() returned error: %v", err)
	}

	if got := part.Header.Get("Content-Type"); got != partContentType {
		t.Errorf("Content-Type = %q, want %q", got, partContentType)
	}

	if got := part.Header.Get(HXRetargetHeader); got != "" {
		t.Errorf("HX-Retarget = %q, want empty", got)
	}

	if body, _ := io.ReadAll(part); string(body) != "<p>bare</p>" {
		t.Errorf("body = %q, want %q", body, "<p>bare</p>")
	}
}
