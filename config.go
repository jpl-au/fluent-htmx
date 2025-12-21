package htmx

import (
	"encoding/json"
	"fmt"
)

// config represents HTMX configuration options.
// Use htmx.Config() to create a new configuration and chain methods to set options.
type config struct {
	config map[string]interface{}
}

// Config creates a new HTMX configuration builder.
func Config() *config {
	return &config{
		config: make(map[string]interface{}),
	}
}

// HistoryEnabled sets whether history is enabled (defaults to true, useful for testing).
func (c *config) HistoryEnabled(enabled bool) *config {
	c.config["historyEnabled"] = enabled
	return c
}

// HistoryCacheSize sets the history cache size (defaults to 10).
func (c *config) HistoryCacheSize(size int) *config {
	c.config["historyCacheSize"] = size
	return c
}

// RefreshOnHistoryMiss sets whether to issue a full page refresh on history misses (defaults to false).
func (c *config) RefreshOnHistoryMiss(refresh bool) *config {
	c.config["refreshOnHistoryMiss"] = refresh
	return c
}

// DefaultSwapStyle sets the default swap style (defaults to "innerHTML").
func (c *config) DefaultSwapStyle(style string) *config {
	c.config["defaultSwapStyle"] = style
	return c
}

// DefaultSwapDelay sets the default swap delay in milliseconds (defaults to 0).
func (c *config) DefaultSwapDelay(delay int) *config {
	c.config["defaultSwapDelay"] = delay
	return c
}

// DefaultSettleDelay sets the default settle delay in milliseconds (defaults to 20).
func (c *config) DefaultSettleDelay(delay int) *config {
	c.config["defaultSettleDelay"] = delay
	return c
}

// IncludeIndicatorStyles sets whether to include indicator styles (defaults to true).
func (c *config) IncludeIndicatorStyles(include bool) *config {
	c.config["includeIndicatorStyles"] = include
	return c
}

// IndicatorClass sets the indicator class name (defaults to "htmx-indicator").
func (c *config) IndicatorClass(class string) *config {
	c.config["indicatorClass"] = class
	return c
}

// RequestClass sets the request class name (defaults to "htmx-request").
func (c *config) RequestClass(class string) *config {
	c.config["requestClass"] = class
	return c
}

// AddedClass sets the added class name (defaults to "htmx-added").
func (c *config) AddedClass(class string) *config {
	c.config["addedClass"] = class
	return c
}

// SettlingClass sets the settling class name (defaults to "htmx-settling").
func (c *config) SettlingClass(class string) *config {
	c.config["settlingClass"] = class
	return c
}

// SwappingClass sets the swapping class name (defaults to "htmx-swapping").
func (c *config) SwappingClass(class string) *config {
	c.config["swappingClass"] = class
	return c
}

// AllowEval sets whether to allow eval for certain features (defaults to true).
func (c *config) AllowEval(allow bool) *config {
	c.config["allowEval"] = allow
	return c
}

// AllowScriptTags sets whether to process script tags in new content (defaults to true).
func (c *config) AllowScriptTags(allow bool) *config {
	c.config["allowScriptTags"] = allow
	return c
}

// InlineScriptNonce sets the nonce to be added to inline scripts (defaults to ”).
func (c *config) InlineScriptNonce(nonce string) *config {
	c.config["inlineScriptNonce"] = nonce
	return c
}

// InlineStyleNonce sets the nonce to be added to inline styles (defaults to ”).
func (c *config) InlineStyleNonce(nonce string) *config {
	c.config["inlineStyleNonce"] = nonce
	return c
}

// AttributesToSettle sets the attributes to settle during the settling phase.
func (c *config) AttributesToSettle(attrs []string) *config {
	c.config["attributesToSettle"] = attrs
	return c
}

// WsReconnectDelay sets the WebSocket reconnect delay strategy (defaults to "full-jitter").
func (c *config) WsReconnectDelay(delay string) *config {
	c.config["wsReconnectDelay"] = delay
	return c
}

// WsBinaryType sets the type of binary data received over WebSocket (defaults to "blob").
func (c *config) WsBinaryType(binaryType string) *config {
	c.config["wsBinaryType"] = binaryType
	return c
}

// DisableSelector sets the selector for disabled elements (defaults to "[hx-disable], [data-hx-disable]").
func (c *config) DisableSelector(selector string) *config {
	c.config["disableSelector"] = selector
	return c
}

// DisableInheritance sets whether to disable attribute inheritance (defaults to false).
func (c *config) DisableInheritance(disable bool) *config {
	c.config["disableInheritance"] = disable
	return c
}

// WithCredentials sets whether to allow cross-site requests with credentials (defaults to false).
func (c *config) WithCredentials(withCreds bool) *config {
	c.config["withCredentials"] = withCreds
	return c
}

// Timeout sets the request timeout in milliseconds (defaults to 0).
func (c *config) Timeout(ms int) *config {
	c.config["timeout"] = ms
	return c
}

// ScrollBehaviour sets the scroll behaviour for the show modifier (defaults to "instant").
// Allowed values: "instant", "smooth", "auto".
func (c *config) ScrollBehaviour(behaviour string) *config {
	c.config["scrollBehavior"] = behaviour
	return c
}

// ScrollBehavior is the American spelling alias for ScrollBehaviour.
func (c *config) ScrollBehavior(behavior string) *config {
	return c.ScrollBehaviour(behavior)
}

// DefaultFocusScroll sets whether focused elements scroll into view (defaults to false).
func (c *config) DefaultFocusScroll(scroll bool) *config {
	c.config["defaultFocusScroll"] = scroll
	return c
}

// GetCacheBusterParam sets whether to append cache buster to GET requests (defaults to false).
func (c *config) GetCacheBusterParam(enable bool) *config {
	c.config["getCacheBusterParam"] = enable
	return c
}

// GlobalViewTransitions sets whether to use View Transition API when swapping (defaults to false).
func (c *config) GlobalViewTransitions(enable bool) *config {
	c.config["globalViewTransitions"] = enable
	return c
}

// MethodsThatUseUrlParams sets which HTTP methods encode parameters in the URL.
func (c *config) MethodsThatUseUrlParams(methods []string) *config {
	c.config["methodsThatUseUrlParams"] = methods
	return c
}

// SelfRequestsOnly sets whether to only allow AJAX requests to the same domain (defaults to true).
func (c *config) SelfRequestsOnly(selfOnly bool) *config {
	c.config["selfRequestsOnly"] = selfOnly
	return c
}

// IgnoreTitle sets whether to ignore title tags in new content (defaults to false).
func (c *config) IgnoreTitle(ignore bool) *config {
	c.config["ignoreTitle"] = ignore
	return c
}

// ScrollIntoViewOnBoost sets whether boosted elements scroll into view (defaults to true).
func (c *config) ScrollIntoViewOnBoost(scroll bool) *config {
	c.config["scrollIntoViewOnBoost"] = scroll
	return c
}

// TriggerSpecsCache sets the cache for evaluated trigger specifications.
func (c *config) TriggerSpecsCache(cache interface{}) *config {
	c.config["triggerSpecsCache"] = cache
	return c
}

// ResponseHandling sets the default response handling behaviour.
func (c *config) ResponseHandling(handling interface{}) *config {
	c.config["responseHandling"] = handling
	return c
}

// AllowNestedOobSwaps sets whether to process OOB swaps on nested elements (defaults to true).
func (c *config) AllowNestedOobSwaps(allow bool) *config {
	c.config["allowNestedOobSwaps"] = allow
	return c
}

// HistoryRestoreAsHxRequest sets whether to treat history cache miss reloads as HX-Request (defaults to true).
func (c *config) HistoryRestoreAsHxRequest(asHxRequest bool) *config {
	c.config["historyRestoreAsHxRequest"] = asHxRequest
	return c
}

// ReportValidityOfForms sets whether to report input validation errors (defaults to false).
func (c *config) ReportValidityOfForms(report bool) *config {
	c.config["reportValidityOfForms"] = report
	return c
}

// ToMetaTag renders the configuration as an HTML meta tag.
// Returns: <meta name="htmx-config" content='{"key":"value"}'>
func (c *config) ToMetaTag() (string, error) {
	if len(c.config) == 0 {
		return "", nil
	}
	jsonBytes, err := json.Marshal(c.config)
	if err != nil {
		return "", fmt.Errorf("failed to marshal htmx config: %w", err)
	}
	return fmt.Sprintf(`<meta name="htmx-config" content='%s'>`, string(jsonBytes)), nil
}

// ToJSON returns the configuration as a JSON string.
func (c *config) ToJSON() (string, error) {
	jsonBytes, err := json.Marshal(c.config)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}
