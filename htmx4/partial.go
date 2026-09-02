package htmx

import (
	"bytes"
	"io"

	"github.com/jpl-au/fluent/node"
)

// Partial builds an <hx-partial> block for a response: content that htmx routes to its own
// target instead of the element that made the request. htmx rewrites the tag to
// <template hx type="partial"> before parsing the response, so the block is inert to the
// browser and processed by htmx alone. Chain HxSwap on the result to choose the swap style;
// innerHTML applies otherwise. The target may use the extended selector forms such as
// "closest li" or "next .error", resolved relative to the element that made the request.
// Pass no inheritance modifier to HxSwap or HxTarget here: htmx reads the block's
// attributes by exact name, so a modified name is ignored.
// Several partials may sit in one response beside the main content, and a response holding
// only partials leaves the main target as it was.
//
//	htmx.Partial("#count", span.Textf("%d", n)).HxSwap(swap.OuterHTML)
func Partial(target string, nodes ...node.Node) *Wrapper {
	return New(&partial{nodes: nodes}).HxTarget(target)
}

// partial is the <hx-partial> element. fluent has no element for a tag outside HTML, so this
// carries the few pieces of node.Element the block needs: children, attributes, and the open
// and close tags.
type partial struct {
	nodes []node.Node
	attrs []node.Attribute
}

var (
	partialOpen  = []byte("<hx-partial")
	partialClose = []byte("</hx-partial>")
)

func (p *partial) SetAttribute(key, value string) {
	p.SetAttributeRaw(key, node.EscapeAttribute(value))
}

func (p *partial) SetAttributeRaw(key, value string) {
	for i, attr := range p.attrs {
		if attr.Key == key {
			p.attrs[i].Value = value
			return
		}
	}
	p.attrs = append(p.attrs, node.Attribute{Key: key, Value: value})
}

func (p *partial) RenderOpen(buf *bytes.Buffer) {
	buf.Write(partialOpen)
	for _, attr := range p.attrs {
		buf.WriteByte(' ')
		buf.WriteString(attr.Key)
		buf.WriteString(`="`)
		buf.WriteString(attr.Value)
		buf.WriteByte('"')
	}
	buf.WriteByte('>')
}

func (p *partial) RenderClose(buf *bytes.Buffer) { buf.Write(partialClose) }

func (p *partial) RenderBuilder(buf *bytes.Buffer) {
	p.RenderOpen(buf)
	for _, child := range p.nodes {
		if child != nil {
			child.RenderBuilder(buf)
		}
	}
	p.RenderClose(buf)
}

func (p *partial) Render(w io.Writer) { _, _ = p.WriteTo(w) }

func (p *partial) WriteTo(w io.Writer) (int64, error) {
	var buf bytes.Buffer
	p.RenderBuilder(&buf)
	return buf.WriteTo(w)
}

func (p *partial) RenderBytes() []byte {
	var buf bytes.Buffer
	p.RenderBuilder(&buf)
	return buf.Bytes()
}

func (p *partial) Nodes() []node.Node { return p.nodes }
