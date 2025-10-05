package table

import (
	"strconv"

	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

type Table struct {
	uid         string
	Caption     string
	HideCaption bool
	HideHeader  bool
	HideFilters bool
	Columns     []Column
}

func (t *Table) Node() gomponents.Node {
	data := t.Data()

	return gomponents.Group{
		gomponents.If(!t.HideCaption, gomponents.Text(t.Caption)),
		html.Table(
			gomponents.If(
				!t.HideHeader,
				html.THead(gomponents.Map(t.Columns, func(tc Column) gomponents.Node { return tc.Node() })),
			),
			html.TBody(
				gomponents.Map(data, func(ss []string) gomponents.Node { r := Row{data: ss}; return r.Node() }),
			),
		),
	}
}

// TODO
func (t *Table) Data() [][]string {
	return [][]string{
		{"sda1", "sda2"},
		{"sdb1", "sdb2"},
		{"sdb1", "sdb2"},
	}
}

func (t *Table) ResolveIds(id string) {
	t.uid = id

	for i, col := range t.Columns {
		col.ResolveIds(id + "-" + strconv.Itoa(i))
	}
}
