package htmx

import (
	"github.com/jpl-au/fluent/html5/template"
	"github.com/jpl-au/fluent/node"
)

// Partial builds an <hx-partial> block for a response: content that htmx routes to its own
// target instead of the element that made the request. It renders as the template form the
// browser converts the tag to, <template hx type="partial" hx-target="...">, which is what
// htmx processes, so the two are the same to the client. Chain HxSwap on the result to choose
// the swap style; the default swap applies otherwise. Several partials may sit in one response
// beside the main content, and a response holding only partials leaves the main target as it
// was.
//
//	htmx.Partial("#count", span.Textf("%d", n)).HxSwap(swap.OuterHTML)
func Partial(target string, nodes ...node.Node) *Wrapper {
	t := template.New(nodes...)
	t.SetAttribute("hx", "")
	t.SetAttribute("type", "partial")

	return New(t).HxTarget(target)
}
