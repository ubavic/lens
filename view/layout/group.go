package layout

import (
	"strconv"

	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

type Group struct {
	uid         string
	Name        string
	Description string
	Collapsible bool
	Columns     []Column
}

func (c *Group) Node() gomponents.Node {
	return gomponents.El("lens-group",
		html.Header(gomponents.Text(c.Name)),
		html.Div(
			gomponents.Map(c.Columns, func(f Column) gomponents.Node { return f.Node() }),
		),
	)
}

func (c *Group) ResolveIds(id string) {
	c.uid = id

	for i, fi := range c.Columns {
		fi.ResolveIds(id + "-" + strconv.Itoa(i))
	}
}
