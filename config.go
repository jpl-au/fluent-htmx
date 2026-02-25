package htmx

import (
	"encoding/json"
	"fmt"

	"github.com/jpl-au/fluent-htmx/swap"
)

// config represents HTMX configuration options.
// Use htmx.Config() to create a new builder and chain methods to set options.
type config struct {
	settings map[string]any
}

// Config creates a new HTMX configuration builder.
func Config() *config {
	return &config{
		settings: make(map[string]any),
	}
}

// HistoryEnabled controls whether HTMX records page snapshots in localStorage for back/forward navigation.
// Defaults to true. Disable during testing to avoid localStorage side effects.
func (c *config) HistoryEnabled(enabled bool) *config {
	c.settings["historyEnabled"] = enabled

	return c
}

// HistoryCacheSize sets the maximum number of pages stored in the history cache.
// Defaults to 10. Increase for applications with deep navigation stacks.
func (c *config) HistoryCacheSize(size int) *config {
	c.settings["historyCacheSize"] = size

	return c
}

// RefreshOnHistoryMiss controls whether a full page refresh is issued when a history entry
// is not found in the local cache. Defaults to false (issues an AJAX request instead).
func (c *config) RefreshOnHistoryMiss(refresh bool) *config {
	c.settings["refreshOnHistoryMiss"] = refresh

	return c
}

// DefaultSwapStyle sets the swap strategy used when no hx-swap attribute is specified.
// Defaults to "innerHTML". Set to swap.OuterHTML for component-based architectures where
// the response replaces the entire target element.
func (c *config) DefaultSwapStyle(style swap.Strategy) *config {
	c.settings["defaultSwapStyle"] = string(style)

	return c
}

// DefaultSwapDelay sets the delay in milliseconds before swapping content.
// Defaults to 0. Useful for adding transition effects before the swap occurs.
func (c *config) DefaultSwapDelay(delay int) *config {
	c.settings["defaultSwapDelay"] = delay

	return c
}

// DefaultSettleDelay sets the delay in milliseconds before settling after a swap.
// Defaults to 20. The settling phase applies attribute changes to new content.
func (c *config) DefaultSettleDelay(delay int) *config {
	c.settings["defaultSettleDelay"] = delay

	return c
}

// IncludeIndicatorStyles controls whether HTMX injects its default CSS for loading indicators.
// Defaults to true. Disable when providing custom indicator styles to avoid conflicts.
func (c *config) IncludeIndicatorStyles(include bool) *config {
	c.settings["includeIndicatorStyles"] = include

	return c
}

// IndicatorClass sets the CSS class name used to identify loading indicator elements.
// Defaults to "htmx-indicator".
func (c *config) IndicatorClass(class string) *config {
	c.settings["indicatorClass"] = class

	return c
}

// RequestClass sets the CSS class name applied to elements (or their indicators) during active requests.
// Defaults to "htmx-request".
func (c *config) RequestClass(class string) *config {
	c.settings["requestClass"] = class

	return c
}

// AddedClass sets the CSS class name applied to newly added content before it settles.
// Defaults to "htmx-added". Useful for entry animations.
func (c *config) AddedClass(class string) *config {
	c.settings["addedClass"] = class

	return c
}

// SettlingClass sets the CSS class name applied during the settling phase.
// Defaults to "htmx-settling". Useful for transition effects on attribute changes.
func (c *config) SettlingClass(class string) *config {
	c.settings["settlingClass"] = class

	return c
}

// SwappingClass sets the CSS class name applied during the swapping phase.
// Defaults to "htmx-swapping". Useful for exit animations before content is replaced.
func (c *config) SwappingClass(class string) *config {
	c.settings["swappingClass"] = class

	return c
}

// AllowEval controls whether HTMX can use eval() for features like hx-vars and hx-on handlers.
// Defaults to true. Disable to comply with Content Security Policy (CSP) restrictions
// that forbid eval — note that hx-on:: (inline event handlers) work without eval.
func (c *config) AllowEval(allow bool) *config {
	c.settings["allowEval"] = allow

	return c
}

// AllowScriptTags controls whether script tags in swapped content are executed.
// Defaults to true. Disable to prevent swapped HTML from running arbitrary JavaScript,
// reducing the XSS attack surface when swapping untrusted content.
func (c *config) AllowScriptTags(allow bool) *config {
	c.settings["allowScriptTags"] = allow

	return c
}

// InlineScriptNonce sets a CSP nonce added to inline scripts injected by HTMX.
// Defaults to empty string (no nonce). Required when the page's Content-Security-Policy
// uses nonce-based script restrictions.
func (c *config) InlineScriptNonce(nonce string) *config {
	c.settings["inlineScriptNonce"] = nonce

	return c
}

// InlineStyleNonce sets a CSP nonce added to inline styles injected by HTMX.
// Defaults to empty string (no nonce). Required when the page's Content-Security-Policy
// uses nonce-based style restrictions.
func (c *config) InlineStyleNonce(nonce string) *config {
	c.settings["inlineStyleNonce"] = nonce

	return c
}

// AttributesToSettle specifies which attributes are updated during the settling phase.
// By default HTMX settles "class" and "style". Override to settle additional or fewer attributes.
func (c *config) AttributesToSettle(attrs []string) *config {
	c.settings["attributesToSettle"] = attrs

	return c
}

// WsReconnectDelay sets the reconnection strategy for WebSocket connections.
// Defaults to "full-jitter". Accepts a fixed delay like "1000" (ms) or a backoff
// strategy like "full-jitter" which adds randomised exponential backoff.
func (c *config) WsReconnectDelay(delay string) *config {
	c.settings["wsReconnectDelay"] = delay

	return c
}

// WsBinaryType sets the binary data type for WebSocket messages.
// Defaults to "blob". Set to "arraybuffer" if the server sends binary frames.
func (c *config) WsBinaryType(binaryType string) *config {
	c.settings["wsBinaryType"] = binaryType

	return c
}

// DisableSelector sets the CSS selector that marks elements where HTMX should not process.
// Defaults to "[hx-disable], [data-hx-disable]".
func (c *config) DisableSelector(selector string) *config {
	c.settings["disableSelector"] = selector

	return c
}

// DisableInheritance globally prevents HTMX attributes from being inherited by child elements.
// Defaults to false. When enabled, each element must explicitly declare its own HTMX attributes.
func (c *config) DisableInheritance(disable bool) *config {
	c.settings["disableInheritance"] = disable

	return c
}

// WithCredentials controls whether cross-origin AJAX requests include cookies and auth headers.
// Defaults to false. Enable for cross-domain APIs that require authentication.
func (c *config) WithCredentials(withCreds bool) *config {
	c.settings["withCredentials"] = withCreds

	return c
}

// Timeout sets the request timeout in milliseconds.
// Defaults to 0 (no timeout). Set a value to prevent requests from hanging indefinitely.
func (c *config) Timeout(ms int) *config {
	c.settings["timeout"] = ms

	return c
}

// ScrollBehaviour sets how the page scrolls when using the "show" modifier on hx-swap.
// Allowed values: "instant" (default), "smooth", "auto".
func (c *config) ScrollBehaviour(behaviour string) *config {
	c.settings["scrollBehavior"] = behaviour

	return c
}

// ScrollBehavior is the American spelling alias for ScrollBehaviour.
func (c *config) ScrollBehavior(behavior string) *config {
	return c.ScrollBehaviour(behavior)
}

// DefaultFocusScroll controls whether a focused element scrolls into view after a swap.
// Defaults to false. Enable to ensure focused inputs are always visible after content updates.
func (c *config) DefaultFocusScroll(scroll bool) *config {
	c.settings["defaultFocusScroll"] = scroll

	return c
}

// GetCacheBusterParam appends a cache-busting query parameter to GET requests.
// Defaults to false. Enable to prevent stale cached responses from intermediary proxies.
func (c *config) GetCacheBusterParam(enable bool) *config {
	c.settings["getCacheBusterParam"] = enable

	return c
}

// GlobalViewTransitions enables the View Transition API for all swaps.
// Defaults to false. When enabled, HTMX wraps swaps in document.startViewTransition()
// for smooth animated transitions between content states.
func (c *config) GlobalViewTransitions(enable bool) *config {
	c.settings["globalViewTransitions"] = enable

	return c
}

// MethodsThatUseURLParams sets which HTTP methods encode parameters in the URL query string.
// Defaults to ["get"]. Other methods send parameters in the request body.
func (c *config) MethodsThatUseURLParams(methods []string) *config {
	c.settings["methodsThatUseUrlParams"] = methods

	return c
}

// SelfRequestsOnly restricts HTMX AJAX requests to the same domain as the page.
// Defaults to true. Disable to allow cross-origin requests (also requires WithCredentials
// for authenticated cross-origin requests).
func (c *config) SelfRequestsOnly(selfOnly bool) *config {
	c.settings["selfRequestsOnly"] = selfOnly

	return c
}

// IgnoreTitle controls whether title tags in swapped content update the page title.
// Defaults to false. Enable to prevent partial swaps from changing the document title.
func (c *config) IgnoreTitle(ignore bool) *config {
	c.settings["ignoreTitle"] = ignore

	return c
}

// ScrollIntoViewOnBoost controls whether boosted elements scroll the target into view.
// Defaults to true. Disable to prevent automatic scrolling on boosted link navigation.
func (c *config) ScrollIntoViewOnBoost(scroll bool) *config {
	c.settings["scrollIntoViewOnBoost"] = scroll

	return c
}

// TriggerSpecsCache provides a pre-populated cache of parsed trigger specifications.
// Expects a map[string]interface{} matching HTMX's internal trigger spec format.
// This is an advanced option for optimising trigger parsing on pages with many elements.
func (c *config) TriggerSpecsCache(cache any) *config {
	c.settings["triggerSpecsCache"] = cache

	return c
}

// ResponseHandling configures how HTMX processes responses based on HTTP status codes.
// Expects a []map[string]interface{} where each entry specifies a code pattern, swap behaviour,
// and error flag. See the HTMX documentation for the default response handling rules.
func (c *config) ResponseHandling(handling any) *config {
	c.settings["responseHandling"] = handling

	return c
}

// AllowNestedOobSwaps controls whether out-of-band swaps are processed on elements
// nested inside other OOB elements. Defaults to true. Disable to prevent unintended
// cascading swaps in deeply nested response fragments.
func (c *config) AllowNestedOobSwaps(allow bool) *config {
	c.settings["allowNestedOobSwaps"] = allow

	return c
}

// HistoryRestoreAsHxRequest controls whether history cache miss reloads are sent with the
// HX-Request header. Defaults to true. This allows the server to distinguish history
// restoration requests from regular page loads.
func (c *config) HistoryRestoreAsHxRequest(asHxRequest bool) *config {
	c.settings["historyRestoreAsHxRequest"] = asHxRequest

	return c
}

// ReportValidityOfForms controls whether HTMX calls reportValidity() on forms before
// submitting. Defaults to false. Enable to show browser-native validation messages
// for invalid inputs before the request is sent.
func (c *config) ReportValidityOfForms(report bool) *config {
	c.settings["reportValidityOfForms"] = report

	return c
}

// ToMetaTag renders the configuration as an HTML meta tag.
// Returns: <meta name="htmx-config" content='{"key":"value"}'>.
func (c *config) ToMetaTag() (string, error) {
	if len(c.settings) == 0 {
		return "", nil
	}

	jsonBytes, err := json.Marshal(c.settings)
	if err != nil {
		return "", fmt.Errorf("failed to marshal htmx config: %w", err)
	}

	return fmt.Sprintf(`<meta name="htmx-config" content='%s'>`, string(jsonBytes)), nil
}

// ToJSON returns the configuration as a JSON string.
func (c *config) ToJSON() (string, error) {
	jsonBytes, err := json.Marshal(c.settings)
	if err != nil {
		return "", err
	}

	return string(jsonBytes), nil
}
