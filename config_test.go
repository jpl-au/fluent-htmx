package htmx

import (
	"strings"
	"testing"
)

func TestConfigToMetaTag(t *testing.T) {
	// Empty config
	cfg := Config()
	tag, err := cfg.ToMetaTag()
	if err != nil {
		t.Fatalf("ToMetaTag() returned error for empty config: %v", err)
	}
	if tag != "" {
		t.Errorf("Expected empty string for empty config, got: %s", tag)
	}

	// Populated config
	cfg.HistoryEnabled(false).Timeout(5000)
	tag, err = cfg.ToMetaTag()
	if err != nil {
		t.Fatalf("ToMetaTag() returned error: %v", err)
	}

	if !strings.HasPrefix(tag, "<meta name=\"htmx-config\" content='") {
		t.Errorf("Invalid meta tag format: %s", tag)
	}
	if !strings.Contains(tag, `"historyEnabled":false`) {
		t.Errorf("Meta tag missing historyEnabled: %s", tag)
	}
	if !strings.Contains(tag, `"timeout":5000`) {
		t.Errorf("Meta tag missing timeout: %s", tag)
	}
}
