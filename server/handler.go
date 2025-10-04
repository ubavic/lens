package server

import (
	"net/http"

	"github.com/ubavic/lens/view"

	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	n := view.RenderHomePage()
	page := page(n)
	page.Render(w)
}

func page(body gomponents.Node) gomponents.Node {
	return html.Doctype(
		html.HTML(html.Lang("en"),
			html.Head(
				html.Meta(html.Charset("utf-8")),
				html.Meta(html.Name("viewport"), html.Content("width=device-width, initial-scale=1")),
				html.TitleEl(gomponents.Text("LENS")),
				html.Link(gomponents.Attr("rel", "stylesheet"), gomponents.Attr("href", "static/main.css")),
				html.Style(`@import url('https://fonts.googleapis.com/css2?family=Roboto:ital,wght@0,100..900;1,100..900&display=swap');`),
			),
			html.Body(html.Main(body)),
		),
	)
}
