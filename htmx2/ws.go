package htmx

// WebSocket extension. Uses the ws-connect / ws-send attributes. Enable it with HxExt("ws") and
// include the WebSocket extension script; without it the attributes are silent no-ops.

// WSConnect sets the ws-connect attribute to establish a WebSocket connection.
// The URL can be absolute or relative. Supports optional ws:// or wss:// prefixes.
// Defaults to the page's scheme, host, and port if not specified.
func (h *Wrapper) WSConnect(url string) *Wrapper {
	h.element.SetAttribute("ws-connect", url)

	return h
}

// WSSend marks the element to transmit data to the WebSocket server.
// Form values are automatically serialised as JSON and sent to the nearest WebSocket connection.
// Includes a HEADERS field with standard htmx request headers.
func (h *Wrapper) WSSend() *Wrapper {
	h.element.SetAttribute("ws-send", "")

	return h
}
