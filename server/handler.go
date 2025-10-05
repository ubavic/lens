package server

import (
	"io"
	"log"
	"net/http"

	"github.com/ubavic/lens/view"
)

func pageHandler(page view.Page) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL.Path)
		page.Node().Render(w)
	}
}

func postHandler(page view.Page) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		trigger := r.Header.Get("HX-Trigger-Name")

		defer r.Body.Close()
		bytes, _ := io.ReadAll(r.Body)
		log.Println(r.Method, r.URL.Path, trigger, string(bytes))

		page.Node().Render(w)
	}
}
