package htmx

import "github.com/jpl-au/fluent-htmx/htmx4/swap"

// Extension configuration. Each extension reads its own object under htmx.config, such as
// htmx.config.sse, htmx.config.ws or htmx.config.multipart, and these setters write into
// those objects. The generated JSON nests them the same way, so the meta tag carries them
// without any extra wiring. Durations are milliseconds, as the extensions read them.

// group returns the nested settings object for an extension, creating it on first use.
func (c *config) group(name string) map[string]any {
	g, ok := c.settings[name].(map[string]any)
	if !ok {
		g = make(map[string]any)
		c.settings[name] = g
	}

	return g
}

// SSEReconnect controls whether the SSE extension reconnects after the stream ends or
// fails. Defaults to true for an hx-sse:connect element and false for a one-shot response.
func (c *config) SSEReconnect(reconnect bool) *config {
	c.group("sse")["reconnect"] = reconnect

	return c
}

// SSEReconnectDelay sets the first reconnect delay in milliseconds. Each failed attempt
// doubles it up to SSEReconnectMaxDelay. Defaults to 500.
func (c *config) SSEReconnectDelay(ms int) *config {
	c.group("sse")["reconnectDelay"] = ms

	return c
}

// SSEReconnectMaxDelay caps the reconnect delay in milliseconds. Defaults to 60000.
func (c *config) SSEReconnectMaxDelay(ms int) *config {
	c.group("sse")["reconnectMaxDelay"] = ms

	return c
}

// SSEReconnectMaxAttempts caps how many times the SSE extension reconnects before giving up.
// Defaults to unlimited.
func (c *config) SSEReconnectMaxAttempts(attempts int) *config {
	c.group("sse")["reconnectMaxAttempts"] = attempts

	return c
}

// SSEReconnectJitter sets the random fraction added to or taken from each reconnect delay,
// so many clients do not reconnect at the same instant. Defaults to 0.3.
func (c *config) SSEReconnectJitter(fraction float64) *config {
	c.group("sse")["reconnectJitter"] = fraction

	return c
}

// SSEPauseOnBackground controls whether the SSE extension drops the connection while the
// tab is hidden and reconnects when it is shown again. Defaults to true for an
// hx-sse:connect element.
func (c *config) SSEPauseOnBackground(pause bool) *config {
	c.group("sse")["pauseOnBackground"] = pause

	return c
}

// SSEReleaseOn sets when the request that opened an SSE stream is released, so the
// element's request cycle completes while the stream stays open. Defaults to
// SSEReleaseImmediate for an hx-sse:connect element and SSEReleaseEnd otherwise.
func (c *config) SSEReleaseOn(when SSERelease) *config {
	c.group("sse")["releaseOn"] = string(when)

	return c
}

// WSReconnect controls whether the WebSocket extension reconnects after the socket closes
// with one of the WSReconnectCodes. Defaults to true.
func (c *config) WSReconnect(reconnect bool) *config {
	c.group("ws")["reconnect"] = reconnect

	return c
}

// WSReconnectCodes lists the WebSocket close codes that trigger a reconnect. Defaults to
// 1001, 1005, 1006, 1011, 1012, 1013 and 1014; a normal close (1000) does not reconnect.
func (c *config) WSReconnectCodes(codes []int) *config {
	c.group("ws")["reconnectCodes"] = codes

	return c
}

// WSReconnectDelay sets the first reconnect delay in milliseconds. Each failed attempt
// doubles it up to WSReconnectMaxDelay. Defaults to 500.
func (c *config) WSReconnectDelay(ms int) *config {
	c.group("ws")["reconnectDelay"] = ms

	return c
}

// WSReconnectMaxDelay caps the reconnect delay in milliseconds. Defaults to 60000.
func (c *config) WSReconnectMaxDelay(ms int) *config {
	c.group("ws")["reconnectMaxDelay"] = ms

	return c
}

// WSReconnectMaxAttempts caps how many times the WebSocket extension reconnects before
// giving up. Defaults to unlimited.
func (c *config) WSReconnectMaxAttempts(attempts int) *config {
	c.group("ws")["reconnectMaxAttempts"] = attempts

	return c
}

// WSReconnectJitter sets the random fraction added to or taken from each reconnect delay.
// Defaults to 0.3.
func (c *config) WSReconnectJitter(fraction float64) *config {
	c.group("ws")["reconnectJitter"] = fraction

	return c
}

// WSPauseOnBackground controls whether the WebSocket extension closes the socket while the
// tab is hidden and reopens it when the tab is shown again. Defaults to true.
func (c *config) WSPauseOnBackground(pause bool) *config {
	c.group("ws")["pauseOnBackground"] = pause

	return c
}

// WSMaxOutgoingMessagesQueueSize caps how many messages sent while the socket is not open
// are held for delivery once it opens. A message past the cap raises htmx:ws:error and is
// dropped. Defaults to 100; 0 turns queuing off.
func (c *config) WSMaxOutgoingMessagesQueueSize(size int) *config {
	c.group("ws")["maxOutgoingMessagesQueueSize"] = size

	return c
}

// WSProtocols lists the sub-protocols offered when a WebSocket is opened, as the second
// argument to the WebSocket constructor. Defaults to none.
func (c *config) WSProtocols(protocols []string) *config {
	c.group("ws")["protocols"] = protocols

	return c
}

// MultipartReconnect controls whether the multipart extension reconnects after a stream ends
// or fails. Defaults to true for an hx-multipart:connect element and false for a one-shot
// response.
func (c *config) MultipartReconnect(reconnect bool) *config {
	c.group("multipart")["reconnect"] = reconnect

	return c
}

// MultipartReconnectDelay sets the first reconnect delay in milliseconds. Each failed
// attempt doubles it up to MultipartReconnectMaxDelay. Defaults to 500.
func (c *config) MultipartReconnectDelay(ms int) *config {
	c.group("multipart")["reconnectDelay"] = ms

	return c
}

// MultipartReconnectMaxDelay caps the reconnect delay in milliseconds. Defaults to 60000.
func (c *config) MultipartReconnectMaxDelay(ms int) *config {
	c.group("multipart")["reconnectMaxDelay"] = ms

	return c
}

// MultipartReconnectMaxAttempts caps how many times the multipart extension reconnects
// before giving up. Defaults to unlimited.
func (c *config) MultipartReconnectMaxAttempts(attempts int) *config {
	c.group("multipart")["reconnectMaxAttempts"] = attempts

	return c
}

// MultipartReconnectJitter sets the random fraction added to or taken from each reconnect
// delay. Defaults to 0.3.
func (c *config) MultipartReconnectJitter(fraction float64) *config {
	c.group("multipart")["reconnectJitter"] = fraction

	return c
}

// MultipartPauseOnBackground controls whether the multipart extension drops the connection
// while the tab is hidden and reconnects when it is shown again. Defaults to true for an
// hx-multipart:connect element.
func (c *config) MultipartPauseOnBackground(pause bool) *config {
	c.group("multipart")["pauseOnBackground"] = pause

	return c
}

// LiveInputDebounce sets how long in milliseconds the live extension waits after an input
// event before re-running expressions. Defaults to 100.
func (c *config) LiveInputDebounce(ms int) *config {
	c.group("live")["inputDebounce"] = ms

	return c
}

// LiveBindPrefix sets an extra attribute prefix that the live extension treats as a
// binding, beside hx-live:. Defaults to ":", so :text="expr" works as hx-live:text. Set
// it to "" when Alpine.js is on the page, because Alpine claims the ":" prefix.
func (c *config) LiveBindPrefix(prefix string) *config {
	c.group("live")["bindPrefix"] = prefix

	return c
}

// LiveUseDollar exposes the live extension's q() selector as $ inside expressions as well.
// Defaults to false.
func (c *config) LiveUseDollar(use bool) *config {
	c.group("live")["useDollar"] = use

	return c
}

// PreloadAutoBoost controls whether boosted links preload on mousedown without an
// hx-preload attribute. Defaults to true.
func (c *config) PreloadAutoBoost(auto bool) *config {
	c.group("preload")["autoBoost"] = auto

	return c
}

// PreloadBoostEvent sets the event that starts a preload on a boosted link or an hx-get
// element without an hx-preload attribute. Defaults to "mousedown".
func (c *config) PreloadBoostEvent(event string) *config {
	c.group("preload")["boostEvent"] = event

	return c
}

// PreloadBoostTimeout sets how long in milliseconds a preloaded response stays usable, for
// a boosted link or an hx-get element without an hx-preload attribute. Defaults to 5000; a
// value of 0 is ignored and the default stays.
func (c *config) PreloadBoostTimeout(ms int) *config {
	c.group("preload")["boostTimeout"] = ms

	return c
}

// HistoryCacheEnabled turns the history-cache extension on or off. The extension is in the
// htmax.js bundle but starts disabled there, so a bundle user calls this with true to opt
// in. Under the bundle, writing any HistoryCache* setting also enables the cache, so a
// bundle user who sets another option and wants the cache off must call this with false.
// With the core build and the separate script, it is on unless this is set to false.
func (c *config) HistoryCacheEnabled(enabled bool) *config {
	c.group("historyCache")["disable"] = !enabled

	return c
}

// HistoryCacheSize sets how many pages the history cache keeps in session storage, with
// the oldest evicted first. Defaults to 10; 0 turns caching off.
func (c *config) HistoryCacheSize(pages int) *config {
	c.group("historyCache")["size"] = pages

	return c
}

// HistoryCacheRefreshOnMiss controls whether a history navigation that misses the cache
// reloads the page rather than fetching it with a request. Defaults to false.
func (c *config) HistoryCacheRefreshOnMiss(refresh bool) *config {
	c.group("historyCache")["refreshOnMiss"] = refresh

	return c
}

// HistoryCacheSwapStyle sets the swap style used when a page is restored from the cache.
// Defaults to swap.OuterSync.
func (c *config) HistoryCacheSwapStyle(style swap.Strategy) *config {
	c.group("historyCache")["swapStyle"] = string(style)

	return c
}

// CompatDoNotTriggerOldEvents stops the htmx-2-compat extension dispatching htmx 2 event
// names alongside the htmx 4 ones. Defaults to false.
func (c *config) CompatDoNotTriggerOldEvents(stop bool) *config {
	c.group("compat")["doNotTriggerOldEvents"] = stop

	return c
}

// CompatUseExplicitInheritance keeps htmx 4's explicit attribute inheritance under the
// htmx-2-compat extension, which otherwise switches implicit inheritance on. Defaults to
// false. The key is written as htmx spells it, useExplicitInheritace.
func (c *config) CompatUseExplicitInheritance(explicit bool) *config {
	c.group("compat")["useExplicitInheritace"] = explicit

	return c
}

// CompatSwapErrorResponseCodes keeps htmx 4's behaviour of swapping error responses under
// the htmx-2-compat extension, which otherwise restores the htmx 2 default of not swapping
// them. Defaults to false.
func (c *config) CompatSwapErrorResponseCodes(swapErrors bool) *config {
	c.group("compat")["swapErrorResponseCodes"] = swapErrors

	return c
}

// CompatSuppressInheritanceLogs silences the htmx-2-compat extension's console notes about
// attributes that htmx 2 inherited and htmx 4 does not, and the htmxImplicitInheritace
// event it dispatches with them. Defaults to false.
func (c *config) CompatSuppressInheritanceLogs(suppress bool) *config {
	c.group("compat")["suppressInheritanceLogs"] = suppress

	return c
}
