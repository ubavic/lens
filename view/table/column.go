package table

import (
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

type Column struct {
	uid         string
	Name        string
	Description string
	SortedBy    int8
	Filter      string
}

func (tc *Column) Node() gomponents.Node {
	return html.Td(gomponents.Text(tc.Name),
		gomponents.If(tc.SortedBy < 0, gomponents.Raw("&#9650;")),
		gomponents.If(tc.SortedBy > 0, gomponents.Raw("&#9660;")),
		gomponents.If(tc.Filter != "", gomponents.Raw("&#9887;")),
	)
}

func (tc *Column) ResolveIds(id string) {
	tc.uid = id
}
