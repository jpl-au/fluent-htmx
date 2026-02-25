// Embedded htmx JavaScript assets.
// The dist/ directory contains htmx core and official extensions, vendored at a
// specific version so that consumers of this package have zero external JS
// dependencies — go get is all that's needed.

package htmx

import (
	"embed"
	"io/fs"
)

// Version is the embedded htmx release.
const Version = "2.0.8"

//go:embed dist
var assets embed.FS

// Assets returns a filesystem rooted at dist/ containing htmx.min.js and the
// ext/ subdirectory. Use this when you need to serve the files through your own
// middleware or at a custom path.
func Assets() fs.FS {
	sub, _ := fs.Sub(assets, "dist")
	return sub
}
