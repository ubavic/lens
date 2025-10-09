package server

import (
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/ubavic/lens/view"
)

func pageHandler(page view.Page) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u, _ := uuid.NewUUID()
		uid := u.String()

		newSession(uid)

		page.Resolve(sessions[uid].handlers)

		log.Println(r.Method, r.URL.Path)
		page.Node(uid).Render(w)
	}
}
