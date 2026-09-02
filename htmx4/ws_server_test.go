package htmx

import (
	"testing"

	"github.com/jpl-au/fluent-htmx/htmx4/swap"
	"github.com/jpl-au/fluent/html5/div"
)

func TestParseWSMessage(t *testing.T) {
	data := []byte(`{"name":"Ann","tags":["a","b"],"count":2,"headers":{"HX-Request":"true","HX-Target":"chat"}}`)
	m, err := ParseWSMessage(data)
	if err != nil {
		t.Fatalf("ParseWSMessage() returned error: %v", err)
	}

	if got := m.Headers["HX-Target"]; got != "chat" {
		t.Errorf("Headers[HX-Target] = %q, want %q", got, "chat")
	}
	if got := m.Values["name"]; got != "Ann" {
		t.Errorf("Values[name] = %v, want %q", got, "Ann")
	}
	if got := m.Values["count"]; got != 2.0 {
		t.Errorf("Values[count] = %v, want 2", got)
	}
	if tags, ok := m.Values["tags"].([]any); !ok || len(tags) != 2 {
		t.Errorf("Values[tags] = %v, want two entries", m.Values["tags"])
	}
	if _, ok := m.Values["headers"]; ok {
		t.Error("Values still holds the headers key")
	}
}

func TestParseWSMessageInvalid(t *testing.T) {
	if _, err := ParseWSMessage([]byte("<p>not json</p>")); err == nil {
		t.Error("ParseWSMessage() returned no error for non-JSON input")
	}
}

func TestWSResponseJSON(t *testing.T) {
	got, err := WSResponse{Content: div.Text("hi"), Target: "#chat", Swap: swap.BeforeEnd}.JSON()
	if err != nil {
		t.Fatalf("JSON() returned error: %v", err)
	}

	want := `{"content":"<div>hi</div>","swap":"beforeend","target":"#chat"}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
