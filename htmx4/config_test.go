package htmx

import (
	"strings"
	"testing"

	"github.com/jpl-au/fluent-htmx/htmx4/swap"
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

// Extension settings nest under the extension's own key, which is how each extension reads
// them from htmx.config.
func TestConfigExtensionGroups(t *testing.T) {
	cfg := Config().
		SSEReleaseOn(SSEReleaseFirst).SSEReconnectDelay(250).
		WSReconnectCodes([]int{1006}).WSProtocols([]string{"chat"}).
		LiveInputDebounce(50).
		PreloadBoostTimeout(2000).
		HistoryCacheEnabled(true).HistoryCacheSwapStyle(swap.OuterSync).
		MultipartReconnect(false).MultipartReconnectDelay(750).
		HistoryReload().SafeEval(true).BoostBrowserIndicator(true).
		CompatUseExplicitInheritance(true).
		LogAll(true).Prefix("data-hx-")

	got, err := cfg.ToJSON()
	if err != nil {
		t.Fatalf("ToJSON() returned error: %v", err)
	}

	for _, frag := range []string{
		`"sse":{"reconnectDelay":250,"releaseOn":"first"}`,
		`"ws":{"protocols":["chat"],"reconnectCodes":[1006]}`,
		`"live":{"inputDebounce":50}`,
		`"preload":{"boostTimeout":2000}`,
		`"historyCache":{"disable":false,"swapStyle":"outerSync"}`,
		`"multipart":{"reconnect":false,"reconnectDelay":750}`,
		`"history":"reload"`,
		`"safeEval":true`,
		`"boostBrowserIndicator":true`,
		`"compat":{"useExplicitInheritace":true}`,
		`"logAll":true`,
		`"prefix":"data-hx-"`,
	} {
		if !strings.Contains(got, frag) {
			t.Errorf("ToJSON() missing %s in %s", frag, got)
		}
	}
}

// A single quote inside a value would end the content attribute unless escaped.
func TestConfigToMetaTagEscapesQuotes(t *testing.T) {
	tag, err := Config().MorphSkip("[data-x='y'] & more").ToMetaTag()
	if err != nil {
		t.Fatalf("ToMetaTag() returned error: %v", err)
	}

	want := `<meta name="htmx-config" content='{"morphSkip":"[data-x=&#39;y&#39;] \u0026 more"}'>`
	if tag != want {
		t.Errorf("tag = %s, want %s", tag, want)
	}
}
