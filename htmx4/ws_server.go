package htmx

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/jpl-au/fluent-htmx/htmx4/swap"
	"github.com/jpl-au/fluent/node"
)

// WSMessage is what an hx-ws:send element transmits to the server: the form values it
// collected, merged with hx-vals, and the htmx request headers under a headers key. Values
// keeps whatever JSON types the client sent, so a number arrives as float64, a multi-value
// field as []any, and an hx-vals boolean as bool.
type WSMessage struct {
	Headers map[string]string
	Values  map[string]any
}

// ParseWSMessage decodes one message received on a WebSocket from an hx-ws:send element.
func ParseWSMessage(data []byte) (*WSMessage, error) {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, fmt.Errorf("failed to decode websocket message: %w", err)
	}

	m := &WSMessage{Headers: map[string]string{}, Values: map[string]any{}}
	if h, ok := raw["headers"]; ok {
		if err := json.Unmarshal(h, &m.Headers); err != nil {
			return nil, fmt.Errorf("failed to decode websocket headers: %w", err)
		}
		delete(raw, "headers")
	}
	for k, v := range raw {
		var value any
		if err := json.Unmarshal(v, &value); err != nil {
			return nil, fmt.Errorf("failed to decode websocket value %q: %w", k, err)
		}
		m.Values[k] = value
	}

	return m, nil
}

// WSResponse is a message the server sends to an hx-ws:connect element to swap content.
// Content is rendered into the message; Target, Swap and Select override the element's own
// hx-target, hx-swap and hx-select when set. Send the JSON form with JSON. A plain HTML
// string sent on the socket is also swapped, using the element's own attributes, so a
// response with no overrides can skip this type and send the rendered bytes directly.
type WSResponse struct {
	Content node.Node
	Target  string
	Swap    swap.Strategy
	Select  string
}

// JSON encodes the response as the object the WebSocket extension expects, with the
// rendered content under content and any overrides beside it.
func (r WSResponse) JSON() ([]byte, error) {
	var buf bytes.Buffer
	if r.Content != nil {
		r.Content.RenderBuilder(&buf)
	}

	out := map[string]string{"content": buf.String()}
	if r.Target != "" {
		out["target"] = r.Target
	}
	if r.Swap != "" {
		out["swap"] = string(r.Swap)
	}
	if r.Select != "" {
		out["select"] = r.Select
	}

	// A plain encoder, so the markup is not escaped into \u003c sequences that a person
	// reading the socket cannot follow.
	var data bytes.Buffer
	enc := json.NewEncoder(&data)
	enc.SetEscapeHTML(false)
	if err := enc.Encode(out); err != nil {
		return nil, fmt.Errorf("failed to encode websocket response: %w", err)
	}

	return bytes.TrimRight(data.Bytes(), "\n"), nil
}
