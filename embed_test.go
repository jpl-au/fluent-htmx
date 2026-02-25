package htmx

import (
	"io/fs"
	"testing"
)

func TestAssetsContainsCore(t *testing.T) {
	f, err := fs.ReadFile(Assets(), "htmx.min.js")
	if err != nil {
		t.Fatalf("failed to read htmx.min.js: %v", err)
	}
	if len(f) == 0 {
		t.Fatal("htmx.min.js is empty")
	}
}

func TestAssetsContainsExtensions(t *testing.T) {
	exts := []string{"ws", "sse", "preload", "response-targets", "head-support"}
	for _, ext := range exts {
		t.Run(ext, func(t *testing.T) {
			f, err := fs.ReadFile(Assets(), "ext/"+ext+".js")
			if err != nil {
				t.Fatalf("failed to read ext/%s.js: %v", ext, err)
			}
			if len(f) == 0 {
				t.Fatalf("ext/%s.js is empty", ext)
			}
		})
	}
}

func TestVersionIsSet(t *testing.T) {
	if Version == "" {
		t.Fatal("Version constant is empty")
	}
}
