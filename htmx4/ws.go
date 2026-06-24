package htmx

// WebSocket extension. The attributes use the hx-ws:* namespace. The extension is bundled in
// htmax.js; with the core htmx.js build, include the WebSocket extension script yourself.

// WsConnect establishes a WebSocket connection.
// The URL can be absolute or relative. Supports optional ws:// or wss:// prefixes.
func (h *Wrapper) WsConnect(url string) *Wrapper {
	h.element.SetAttribute("hx-ws:connect", url)

	return h
}

// WsSend marks the element to transmit data to the WebSocket server.
// Form values are serialised and sent to the nearest WebSocket connection.
func (h *Wrapper) WsSend() *Wrapper {
	h.element.SetAttribute("hx-ws:send", "")

	return h
}
