package layout

import (
	"strconv"

	"github.com/ubavic/lens/view/component"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

type Column struct {
	uid      string
	Name     string
	HideName string
	Span     uint
	Elements []component.Component
}

func (c *Column) Node() gomponents.Node {
	return gomponents.El("lens-column",
		html.Header(gomponents.Text(c.Name)),
		gomponents.Map(c.Elements, func(f component.Component) gomponents.Node { return f.Node() }),
	)
}

func (c *Column) ResolveIds(id string) {
	c.uid = id

	for i, fi := range c.Elements {
		fi.ResolveIds(id + "-" + strconv.Itoa(i))
	}
}
