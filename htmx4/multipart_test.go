package htmx

import (
	"strings"
	"testing"

	"github.com/jpl-au/fluent/html5/div"
)

func TestMultipartClientSetters(t *testing.T) {
	cases := []setterCase{
		{"MultipartConnect", func(w *Wrapper) { w.MultipartConnect("/stream") }, `hx-multipart:connect="/stream"`},
		{"MultipartClose", func(w *Wrapper) { w.MultipartClose("done") }, `hx-multipart:close="done"`},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			d := div.New()
			c.set(New(d))

			if got := string(d.RenderBytes()); !strings.Contains(got, c.want) {
				t.Errorf("%s rendered %q, want it to contain %q", c.name, got, c.want)
			}
		})
	}
}
