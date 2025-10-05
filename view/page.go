package view

import (
	"strconv"

	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

type Page struct {
	Id         string
	Name       string
	Components []Component
}

func (p Page) Node() gomponents.Node {
	return html.Doctype(
		html.HTML(html.Lang("en"),
			html.Head(
				html.Meta(html.Charset("utf-8")),
				html.Meta(html.Name("viewport"), html.Content("width=device-width, initial-scale=1")),
				html.TitleEl(gomponents.Text("LENS")),
				html.Link(gomponents.Attr("rel", "stylesheet"), gomponents.Attr("href", "/static/main.css")),
				html.Style(`@import url('https://fonts.googleapis.com/css2?family=Roboto:ital,wght@0,100..900;1,100..900&display=swap');`),
				html.Script(gomponents.Attr("src", "/static/htmx.min.js")),
			),
			html.Body(html.Main(gomponents.Map(p.Components, func(f Component) gomponents.Node { return f.Node() }))),
		),
	)
}

func (p Page) Post(resource string) gomponents.Node {
	return p.Node()
}

func (p Page) ResolveIds() {
	for i, c := range p.Components {
		c.ResolveIds(strconv.Itoa(i))
	}
}
