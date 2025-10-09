package layout

import (
	"strconv"

	"github.com/ubavic/lens/view/component"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

type Group struct {
	uid         component.Uid
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

func (c *Group) Resolve(id component.Uid, hm component.HandlerMap) {
	c.uid = id

	for i, fi := range c.Columns {
		fi.Resolve(id+"-"+strconv.Itoa(i), hm)
	}
}
