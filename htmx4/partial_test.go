package htmx

import (
	"testing"

	"github.com/jpl-au/fluent-htmx/htmx4/swap"
	"github.com/jpl-au/fluent/html5/span"
)

func TestPartial(t *testing.T) {
	got := string(Partial("#count", span.Text("3")).HxSwap(swap.OuterHTML).RenderBytes())
	want := `<template hx="" type="partial" hx-target="#count" hx-swap="outerHTML"><span>3</span></template>`
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
