// HTTP handler for serving embedded htmx assets.

package htmx

import (
	"net/http"
)

// Handler returns an http.Handler that serves the embedded htmx JS files.
// The prefix must match the path where the handler is mounted in the router,
// including the trailing slash.
//
// The handler sets long-lived cache headers because the files are versioned by
// the Go module — a new module version means new file contents.
//
// Usage:
//
//	mux.Handle("/_htmx/", htmx.Handler("/_htmx/"))
func Handler(prefix string) http.Handler {
	fs := http.StripPrefix(prefix, http.FileServer(http.FS(Assets())))

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "public, max-age=31536000, immutable")
		fs.ServeHTTP(w, r)
	})
}
