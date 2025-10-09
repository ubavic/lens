package table

import (
	"github.com/ubavic/lens/view/component"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

type Row struct {
	data []string
}

func (t *Row) Node() gomponents.Node {
	return html.Tr(
		gomponents.Map(t.data, func(s string) gomponents.Node { return html.Td(gomponents.Text(s)) }),
	)
}

func (t *Row) Resolve(id component.Uid, hm component.HandlerMap) {
	// TODO
}
