// Fluent node helpers for htmx script tags.

package htmx

import (
	"bytes"
	"io"

	"github.com/jpl-au/fluent/node"
)

// scriptNode renders a <script src="..."></script> tag. It implements
// node.Node so it can be placed directly in a fluent tree.
type scriptNode struct {
	src string
}

// Script returns a node that renders <script src="{prefix}htmx.min.js"></script>.
// The prefix should match the path where Handler is mounted.
//
//	htmx.Script("/_htmx/") // <script src="/_htmx/htmx.min.js"></script>
func Script(prefix string) node.Node {
	return &scriptNode{src: prefix + "htmx.min.js"}
}

// ExtScript returns a node that renders <script src="{prefix}ext/{name}.js"></script>.
// The prefix should match the path where Handler is mounted.
//
//	htmx.ExtScript("/_htmx/", "ws") // <script src="/_htmx/ext/ws.js"></script>
func ExtScript(prefix string, name string) node.Node {
	return &scriptNode{src: prefix + "ext/" + name + ".js"}
}

func (s *scriptNode) Render(w ...io.Writer) []byte {
	var buf bytes.Buffer
	s.RenderBuilder(&buf)
	if len(w) > 0 && w[0] != nil {
		_, _ = buf.WriteTo(w[0])
		return nil
	}
	return buf.Bytes()
}

func (s *scriptNode) RenderBuilder(buf *bytes.Buffer) {
	buf.WriteString(`<script src="`)
	buf.WriteString(s.src)
	buf.WriteString(`"></script>`)
}

func (s *scriptNode) Nodes() []node.Node { return nil }
