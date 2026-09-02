package htmx

// Multipart streaming extension. The attributes use the hx-multipart:* namespace. It is a
// separate script (not bundled in htmax.js) - include the multipart extension script yourself.
//
// The extension streams a multipart/mixed or multipart/parallel response and swaps each part
// independently: with multipart/mixed the parts swap in the order they arrive, with
// multipart/parallel each swaps as soon as it arrives. Every part carries its own HX-* headers
// (target, swap, select, trigger), so one response can update several places on the page. Build
// the response server-side with [MultipartWriter], choosing the mode with a multipart.Type.
//
// A one-shot multipart response needs no attribute: the extension acts on any response whose
// Content-Type is a multipart.Type value. The hx-multipart:connect attribute is only for the
// long-lived case, opening a reconnecting stream for server-push updates.

// MultipartConnect opens a long-lived multipart stream to url, reconnecting on failure. The
// connection is established on the element's hx-trigger, or on load if no trigger is set. Use it
// for server-push updates; a one-shot multipart response needs no attribute, since the extension
// acts on any response whose Content-Type is a multipart type.
func (h *Wrapper) MultipartConnect(url string) *Wrapper {
	h.element.SetAttribute("hx-multipart:connect", url)

	return h
}

// MultipartClose closes the multipart connection when the given trigger fires. It takes effect
// only on an element that also has hx-multipart:connect. The value is an hx-trigger spec; the
// usual value is the name of an event a part fires through its HX-Trigger header, such as
// "done", and that part still swaps before the connection closes.
func (h *Wrapper) MultipartClose(trigger string) *Wrapper {
	h.element.SetAttribute("hx-multipart:close", trigger)

	return h
}
