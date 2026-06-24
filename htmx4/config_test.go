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
	cfg.HistoryEnabled(false).DefaultTimeout(5000)

	tag, err = cfg.ToMetaTag()
	if err != nil {
		t.Fatalf("ToMetaTag() returned error: %v", err)
	}

	if !strings.HasPrefix(tag, "<meta name=\"htmx-config\" content='") {
		t.Errorf("Invalid meta tag format: %s", tag)
	}

	if !strings.Contains(tag, `"history":false`) {
		t.Errorf("Meta tag missing history key: %s", tag)
	}

	if !strings.Contains(tag, `"defaultTimeout":5000`) {
		t.Errorf("Meta tag missing defaultTimeout key: %s", tag)
	}
}

func TestConfigToJSON(t *testing.T) {
	cfg := Config().DefaultSwapStyle("outerHTML").Transitions(true)

	jsonStr, err := cfg.ToJSON()
	if err != nil {
		t.Fatalf("ToJSON() returned error: %v", err)
	}

	if !strings.Contains(jsonStr, `"defaultSwap":"outerHTML"`) {
		t.Errorf("JSON missing defaultSwap: %s", jsonStr)
	}

	if !strings.Contains(jsonStr, `"transitions":true`) {
		t.Errorf("JSON missing transitions: %s", jsonStr)
	}
}

func TestConfigChaining(t *testing.T) {
	// All config methods should support chaining.
	cfg := Config().
		HistoryEnabled(true).
		DefaultSwapStyle("outerHTML").
		DefaultTimeout(3000).
		Mode("same-origin").
		NoSwap([]string{"204", "304", "4xx", "5xx"}).
		MorphSkip("#preserve")

	jsonStr, err := cfg.ToJSON()
	if err != nil {
		t.Fatalf("ToJSON() returned error: %v", err)
	}

	for _, expected := range []string{"history", "defaultSwap", "defaultTimeout", "mode", "noSwap", "morphSkip"} {
		if !strings.Contains(jsonStr, expected) {
			t.Errorf("Chained config missing %s: %s", expected, jsonStr)
		}
	}
}
