package htmx

import (
	"encoding/json"
	"fmt"
	"strings"

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

// HistoryReload keeps history tracking on but makes a back or forward navigation reload
// the page from the server instead of swapping the restored content in. It writes the
// "reload" value of the history setting; HistoryEnabled writes the true and false values.
func (c *config) HistoryReload() *config {
	c.settings["history"] = "reload"

	return c
}

// SafeEval makes the hx-csp extension run hx-on, hx-vals js: and hx-confirm js: code
// through nonce-carrying script injection instead of the Function constructor, so a page
// can drop unsafe-eval from its policy. Defaults to false. Read by the hx-csp extension.
func (c *config) SafeEval(safe bool) *config {
	c.settings["safeEval"] = safe

	return c
}

// BoostBrowserIndicator shows the browser's native loading indicator for every boosted
// request, without an hx-browser-indicator attribute on each element. Defaults to false.
// Read by the browser-indicator extension.
func (c *config) BoostBrowserIndicator(on bool) *config {
	c.settings["boostBrowserIndicator"] = on

	return c
}

// LogAll turns on htmx's console logging of every event it dispatches. Defaults to false.
func (c *config) LogAll(log bool) *config {
	c.settings["logAll"] = log

	return c
}

// Prefix sets an extra attribute prefix that htmx reads beside hx-, so data-hx-get works
// as hx-get. Defaults to "data-hx-". The typed setters in this package always write the
// hx- form.
func (c *config) Prefix(prefix string) *config {
	c.settings["prefix"] = prefix

	return c
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
// swapped into the DOM. The list replaces htmx's default of 204 and 304 rather than adding
// to it, so repeat those when keeping, for example, error responses from replacing content:
// NoSwap([]string{"204", "304", "4xx", "5xx"}).
func (c *config) NoSwap(codes []string) *config {
	c.settings["noSwap"] = codes

	return c
}

// Extensions sets the comma-separated allow list of extensions permitted to activate.
// Extensions are loaded by including their script; this list restricts which of the loaded
// extensions htmx will actually run, so an included script cannot activate unless named here.
// Names are the registration names, which differ from the attribute and script names for
// some: sse, ws, preload, download, upsert, ptag, browser-indicator, history-cache,
// alpine-compat, compat, hx-pending, hx-targets, hx-live, hx-head, hx-csp, hx-prompt and
// hx-multipart.
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

// AllowEmptySwapAfterOOB controls what happens to the main swap target when a response holds
// only out-of-band content. By default the main swap is skipped, so the target keeps what it
// had. Pass true to swap the empty remainder in, which clears the target. Defaults to false.
func (c *config) AllowEmptySwapAfterOOB(allow bool) *config {
	c.settings["allowEmptySwapAfterOOB"] = allow

	return c
}

// MorphIgnore lists attribute name prefixes that morph swaps preserve on existing elements
// rather than overwrite from the response, so external libraries can keep their own
// attributes across a morph. Each entry matches every attribute whose name starts with it,
// so "data-my-" covers data-my-state and data-my-other. Defaults to ["data-htmx-powered"].
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

// MorphSkip sets the CSS selector matching elements that morph swaps leave untouched,
// preserving them and their subtrees verbatim across a morph. Defaults to
// "[hx-morph-skip]", which is what HxMorphSkip writes; a new selector replaces that, so
// include it in the new value to keep the attribute working.
func (c *config) MorphSkip(selector string) *config {
	c.settings["morphSkip"] = selector

	return c
}

// MorphSkipChildren sets the CSS selector matching elements whose children morph swaps leave
// untouched, while still morphing the matched element itself. Defaults to
// "[hx-morph-skip-children]", which is what HxMorphSkipChildren writes; a new selector
// replaces that, so include it in the new value to keep the attribute working.
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

	// The JSON sits inside a single-quoted attribute. json.Marshal escapes <, > and &
	// inside strings, but a single quote passes through and would end the attribute.
	// The browser decodes the entity before htmx reads the content.
	content := strings.ReplaceAll(string(jsonBytes), "'", "&#39;")

	return fmt.Sprintf(`<meta name="htmx-config" content='%s'>`, content), nil
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
