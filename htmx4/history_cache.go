package htmx

// History-cache extension (hx-history). Not bundled in htmax.js - load it as a separate script
// (dist/ext/hx-history-cache.js). Global cache options live on htmx.config.historyCache.

// HxHistory controls whether the element's page is stored in the history cache. Pass false to
// keep a sensitive page out of the cache (emitting hx-history="false"); the default is to
// cache. Distinct from HxHistoryElt, which names the snapshot element for restoration.
func (h *Wrapper) HxHistory(enabled bool) *Wrapper {
	value := boolFalse
	if enabled {
		value = boolTrue
	}

	h.element.SetAttribute("hx-history", value)

	return h
}
