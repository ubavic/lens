package view

import (
	"strconv"

	"github.com/ubavic/lens/view/component"
	"github.com/ubavic/lens/view/navigation"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

type Page struct {
	Id         string
	Name       string
	Components []component.Component
}

func (p Page) Node(sessionId string) gomponents.Node {
	nav := navigation.Side{}

	return html.Doctype(
		html.HTML(html.Lang("en"),
			html.Head(
				html.Meta(html.Charset("utf-8")),
				html.Meta(html.Name("viewport"), html.Content("width=device-width, initial-scale=1")),
				html.TitleEl(gomponents.Text(p.Name)),
				html.Link(gomponents.Attr("rel", "stylesheet"), gomponents.Attr("href", "/static/main.css")),
				html.Style(`@import url('https://fonts.googleapis.com/css2?family=Roboto:ital,wght@0,100..900;1,100..900&display=swap');`),
				html.Script(gomponents.Attr("src", "/static/ws.js")),
			),
			html.Body(
				gomponents.Attr("lens-session-id", sessionId),
				nav.Node(),
				html.Main(
					html.Header(gomponents.Text(p.Name)),
					gomponents.Map(p.Components, func(f component.Component) gomponents.Node { return f.Node() }),
				)),
		),
	)
}

func (p Page) Post(resource string) gomponents.Node {
	return p.Node("")
}

func (p Page) Resolve(hm component.HandlerMap) {
	for i, c := range p.Components {
		c.Resolve(component.Uid(strconv.Itoa(i)), hm)
	}
}
