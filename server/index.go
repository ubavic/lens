package server

import "net/http"

var indexRoutes []string

func indexHandler(w http.ResponseWriter, r *http.Request) {
	home := `<html><body><ol>`

	for _, r := range indexRoutes {
		home += `<li><a href='/` + r + `'>` + r + `</a></li>`
	}

	home += `</ol></body></html>`

	w.Write([]byte(home))
}
