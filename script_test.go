package htmx

import "testing"

func TestScript(t *testing.T) {
	n := Script("/_htmx/")
	got := string(n.Render())
	want := `<script src="/_htmx/htmx.min.js"></script>`
	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestExtScript(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"ws", `<script src="/_htmx/ext/ws.js"></script>`},
		{"sse", `<script src="/_htmx/ext/sse.js"></script>`},
		{"preload", `<script src="/_htmx/ext/preload.js"></script>`},
		{"response-targets", `<script src="/_htmx/ext/response-targets.js"></script>`},
		{"head-support", `<script src="/_htmx/ext/head-support.js"></script>`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := string(ExtScript("/_htmx/", tt.name).Render())
			if got != tt.want {
				t.Fatalf("got %q, want %q", got, tt.want)
			}
		})
	}
}

func TestScriptWritesToWriter(t *testing.T) {
	n := Script("/_htmx/")
	var buf []byte
	w := &byteWriter{buf: &buf}
	result := n.Render(w)
	if result != nil {
		t.Fatal("expected nil return when writer provided")
	}
	got := string(*w.buf)
	want := `<script src="/_htmx/htmx.min.js"></script>`
	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

// byteWriter is a minimal io.Writer for testing.
type byteWriter struct {
	buf *[]byte
}

func (w *byteWriter) Write(p []byte) (int, error) {
	*w.buf = append(*w.buf, p...)
	return len(p), nil
}
