package htmx

// Server-Sent Events extension. The attributes use the hx-sse:* namespace. The extension is
// bundled in htmax.js; with the core htmx.js build, include the SSE extension script yourself.
// Unnamed SSE messages are swapped automatically, and named events are dispatched as DOM
// events (listen for them with HxTrigger).

// SSEConnect establishes a Server-Sent Events connection to the given URL.
// The connection remains open and automatically reconnects on failure.
func (h *Wrapper) SSEConnect(url string) *Wrapper {
	h.element.SetAttribute("hx-sse:connect", url)

	return h
}

// SSEClose closes the SSE connection when the named event is received.
// Useful for finite streams where the server signals completion.
func (h *Wrapper) SSEClose(eventName string) *Wrapper {
	h.element.SetAttribute("hx-sse:close", eventName)

	return h
}
