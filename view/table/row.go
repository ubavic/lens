package table

import (
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

func (t *Row) ResolveIds(id string) {
	// TODO
}
