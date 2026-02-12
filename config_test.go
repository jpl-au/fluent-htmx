package htmx

import (
	"strings"
	"testing"
)

func TestConfigToMetaTag(t *testing.T) {
	// Empty config should return empty string.
	cfg := Config()

	tag, err := cfg.ToMetaTag()
	if err != nil {
		t.Fatalf("ToMetaTag() returned error for empty config: %v", err)
	}

	if tag != "" {
		t.Errorf("Expected empty string for empty config, got: %s", tag)
	}

	// Populated config should render as a meta tag.
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

func TestConfigToJSON(t *testing.T) {
	cfg := Config().DefaultSwapStyle("outerHTML").GlobalViewTransitions(true)

	jsonStr, err := cfg.ToJSON()
	if err != nil {
		t.Fatalf("ToJSON() returned error: %v", err)
	}

	if !strings.Contains(jsonStr, `"defaultSwapStyle":"outerHTML"`) {
		t.Errorf("JSON missing defaultSwapStyle: %s", jsonStr)
	}

	if !strings.Contains(jsonStr, `"globalViewTransitions":true`) {
		t.Errorf("JSON missing globalViewTransitions: %s", jsonStr)
	}
}

func TestConfigChaining(t *testing.T) {
	// All config methods should support chaining.
	cfg := Config().
		HistoryEnabled(true).
		HistoryCacheSize(20).
		DefaultSwapStyle("outerHTML").
		Timeout(3000).
		SelfRequestsOnly(false)

	jsonStr, err := cfg.ToJSON()
	if err != nil {
		t.Fatalf("ToJSON() returned error: %v", err)
	}

	for _, expected := range []string{"historyEnabled", "historyCacheSize", "defaultSwapStyle", "timeout", "selfRequestsOnly"} {
		if !strings.Contains(jsonStr, expected) {
			t.Errorf("Chained config missing %s: %s", expected, jsonStr)
		}
	}
}
