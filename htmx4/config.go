package htmx

import (
	"encoding/json"
	"fmt"

	"github.com/jpl-au/fluent-htmx/htmx4/swap"
)

// config represents HTMX configuration options.
// Use htmx.Config() to create a new builder and chain methods to set options.
type config struct {
	settings map[string]any
}

// Config starts a new htmx configuration builder. Chain setters onto it to override
// individual htmx.config options; any option left unset keeps htmx's own default, so
// you only spell out what you want to change. Emit the finished configuration with
// ToMetaTag to embed it in the page head, or ToJSON for the raw settings object.
func Config() *config {
	return &config{
		settings: make(map[string]any),
	}
}

// HistoryEnabled controls whether htmx tracks history for back/forward navigation.
// When the user navigates back, htmx re-fetches the page and swaps it into the body
// (or into the hx-history-elt element if one is present). Defaults to true; pass false
// to disable history handling so navigation falls back to a normal browser page load.
func (c *config) HistoryEnabled(enabled bool) *config {
	c.settings["history"] = enabled

	return c
}

// DefaultSwapStyle sets the swap strategy used when an element has no hx-swap attribute.
// Defaults to "innerHTML". Set it to swap.OuterHTML for component-style architectures
// where each response is meant to replace the entire target element.
func (c *config) DefaultSwapStyle(style swap.Strategy) *config {
	c.settings["defaultSwap"] = string(style)

	return c
}

// DefaultSettleDelay sets the delay in milliseconds before the settle phase runs after a
// swap. The settle phase applies the new content's attribute changes (such as class and
// style), giving CSS transitions a frame to start. Defaults to 1.
func (c *config) DefaultSettleDelay(delay int) *config {
	c.settings["defaultSettleDelay"] = delay

	return c
}

// DefaultTimeout sets the request timeout in milliseconds; requests that exceed it are
// aborted. Defaults to 60000 (60 seconds). Set 0 to disable the timeout and let requests
// run indefinitely.
func (c *config) DefaultTimeout(ms int) *config {
	c.settings["defaultTimeout"] = ms

	return c
}

// Transitions enables the View Transitions API for all swaps. Defaults to false. When
// enabled, htmx wraps each swap in document.startViewTransition() so the browser can
// animate smoothly between the old and new content.
func (c *config) Transitions(enable bool) *config {
	c.settings["transitions"] = enable

	return c
}

// IncludeIndicatorCSS controls whether htmx injects its default CSS for loading indicators.
// Defaults to true. Disable it when you supply your own indicator styles, so the built-in
// rules do not conflict with yours.
func (c *config) IncludeIndicatorCSS(include bool) *config {
	c.settings["includeIndicatorCSS"] = include

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

// InlineScriptNonce sets a CSP nonce added to inline scripts injected by HTMX.
// Defaults to empty string (no nonce).
func (c *config) InlineScriptNonce(nonce string) *config {
	c.settings["inlineScriptNonce"] = nonce

	return c
}

// DefaultFocusScroll controls whether a focused element scrolls into view after a swap.
// Defaults to false.
func (c *config) DefaultFocusScroll(scroll bool) *config {
	c.settings["defaultFocusScroll"] = scroll

	return c
}

// ImplicitInheritance makes inheritable attributes cascade to descendant elements without
// an explicit modifier. Inheritance is explicit by default - each inheriting attribute
// must carry the :inherited modifier - so enabling this makes every inheritable attribute
// inherit automatically instead. Defaults to false.
func (c *config) ImplicitInheritance(implicit bool) *config {
	c.settings["implicitInheritance"] = implicit

	return c
}

// NoSwap lists the response status codes (and code patterns) whose bodies must not be
// swapped into the DOM. Every response except 204 and 304 is swapped by default, so add
// codes here to keep, for example, error responses from replacing content:
// NoSwap([]string{"4xx", "5xx"}).
func (c *config) NoSwap(codes []string) *config {
	c.settings["noSwap"] = codes

	return c
}

// Extensions sets the comma-separated allow list of extension names permitted to activate.
// Extensions are loaded by including their script; this list restricts which of the loaded
// extensions htmx will actually run, so an included script cannot activate unless named here.
func (c *config) Extensions(extensions string) *config {
	c.settings["extensions"] = extensions

	return c
}

// Mode sets the fetch mode used for requests, which governs cross-origin behaviour.
// Defaults to "same-origin", restricting requests to the page's own origin; widen it
// (for example to "cors") to allow cross-origin requests.
func (c *config) Mode(mode string) *config {
	c.settings["mode"] = mode

	return c
}

// MetaCharacter sets the separator character used in attribute and event names.
// Defaults to ":". Set to "-" for frameworks (e.g. JSX) that disallow ":" in attribute names.
//
// The typed helpers that bake in ":" - the Inherited/InheritedAppend/Append modifiers,
// HxOn, HxStatus and the event/sync constants - are not derived from this setting. If you
// change the meta character, build those attribute names yourself with SetAttribute using
// the configured separator.
func (c *config) MetaCharacter(char string) *config {
	c.settings["metaCharacter"] = char

	return c
}

// MorphIgnore lists full attribute names that morph swaps preserve on existing elements
// rather than overwrite from the response, so external libraries can keep their own
// attributes across a morph. htmx matches each entry by exact name, not by prefix, so list
// every complete attribute name (e.g. "data-my-state", not "data-my-"). Defaults to
// ["data-htmx-powered"].
func (c *config) MorphIgnore(attributeNames []string) *config {
	c.settings["morphIgnore"] = attributeNames

	return c
}

// MorphScanLimit caps how many elements htmx scans when matching nodes during a morph
// swap. Lower it to bound the work on very large subtrees, at the cost of match accuracy.
// Defaults to 10.
func (c *config) MorphScanLimit(limit int) *config {
	c.settings["morphScanLimit"] = limit

	return c
}

// MorphSkip sets a CSS selector matching elements that morph swaps leave untouched,
// preserving them and their subtrees verbatim across a morph.
func (c *config) MorphSkip(selector string) *config {
	c.settings["morphSkip"] = selector

	return c
}

// MorphSkipChildren sets a CSS selector matching elements whose children morph swaps leave
// untouched, while still morphing the matched element itself.
func (c *config) MorphSkipChildren(selector string) *config {
	c.settings["morphSkipChildren"] = selector

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

// ToJSON marshals only the options you set into a JSON object matching the shape of
// htmx.config, for example {"defaultTimeout":5000,"transitions":true}. Reach for it
// when you need the raw settings - to configure htmx.config from your own JavaScript,
// or to embed the object yourself; ToMetaTag wraps this same JSON in an htmx-config
// meta tag ready to drop into the page head.
func (c *config) ToJSON() (string, error) {
	jsonBytes, err := json.Marshal(c.settings)
	if err != nil {
		return "", err
	}

	return string(jsonBytes), nil
}
