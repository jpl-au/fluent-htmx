// SSE extension methods for HTMX.

package htmx

// SSESwap sets the sse-swap attribute to listen for a specific SSE event and swap content.
func (h *Wrapper) SSESwap(eventName string) *Wrapper {
	h.node.SetAttribute("sse-swap", eventName)

	return h
}

// SSEConnect sets the sse-connect attribute to establish an SSE connection.
func (h *Wrapper) SSEConnect(url string) *Wrapper {
	h.node.SetAttribute("sse-connect", url)

	return h
}

// SSEClose sets the sse-close attribute to specify when to close the SSE connection.
func (h *Wrapper) SSEClose(eventName string) *Wrapper {
	h.node.SetAttribute("sse-close", eventName)

	return h
}
