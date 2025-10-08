package navigation

import (
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

type Entry struct {
	Name string
	Icon string
	Path string
}

func (e Entry) Node() gomponents.Node {
	return html.A(html.Href(e.Path), gomponents.Text(e.Name))
}
