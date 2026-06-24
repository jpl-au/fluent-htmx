package htmx

// SSEConnect establishes a Server-Sent Events connection to the given URL.
// The connection remains open and automatically reconnects on failure.
// Requires the SSE extension to be enabled via HxExt("sse").
func (h *Wrapper) SSEConnect(url string) *Wrapper {
	h.element.SetAttribute("sse-connect", url)

	return h
}

// SSESwap listens for a named SSE event and swaps its data into the element.
// The event name must match what the server sends in the SSE "event:" field.
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
