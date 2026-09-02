package htmx

// History-cache extension, registered as history-cache. It is bundled in htmax.js but starts
// disabled there; Config().HistoryCacheEnabled(true) turns it on. With the core htmx.js build,
// include dist/ext/hx-history-cache.js. Cache options live under Config().HistoryCache*.

// HxHistoryExclude keeps the element's page out of the history cache, by writing
// hx-history="false". The extension reads only that value, and every page is cached unless
// one element on it carries the attribute. Use it on a page holding sensitive content.
// Distinct from HxHistoryElt, which names the snapshot element for restoration.
func (h *Wrapper) HxHistoryExclude() *Wrapper {
	h.element.SetAttribute("hx-history", boolFalse)

	return h
}
