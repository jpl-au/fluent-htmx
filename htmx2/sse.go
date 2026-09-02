package htmx

// Server-Sent Events extension. Uses the sse-connect / sse-swap attributes. Enable it with
// HxExt("sse") and include the SSE extension script; without it the attributes are silent no-ops.

// SSEConnect establishes a Server-Sent Events connection to the given URL.
// The connection remains open and automatically reconnects on failure.
// Requires the SSE extension to be enabled via HxExt("sse").
func (h *Wrapper) SSEConnect(url string) *Wrapper {
	h.element.SetAttribute("sse-connect", url)

	return h
}

// SSESwap listens for a named SSE event and swaps its data into the element.
// The name must match the server's "event:" field; an event sent without a name
// arrives as "message". Several names may be given separated by commas. The
// element must be the sse-connect element or a descendant of it. An event can
// also trigger a request instead, through HxTrigger("sse:name").
func (h *Wrapper) SSESwap(eventName string) *Wrapper {
	h.element.SetAttribute("sse-swap", eventName)

	return h
}

// SSEClose closes the SSE connection when the specified event is received.
// Useful for finite streams where the server signals completion.
func (h *Wrapper) SSEClose(eventName string) *Wrapper {
	h.element.SetAttribute("sse-close", eventName)

	return h
}
