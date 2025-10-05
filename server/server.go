package server

import (
	"log"
	"net/http"

	"github.com/ubavic/lens/view"
)

type ServerConfig struct {
	Port  string
	Pages []view.Page
}

func StartServer(config ServerConfig) error {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("GET /static/", http.StripPrefix("/static/", fs))

	for _, p := range config.Pages {
		p.ResolveIds()
		http.HandleFunc("POST /"+p.Id+"/", postHandler(p))
		http.HandleFunc("/"+p.Id+"/", pageHandler(p))
	}

	log.Printf("Server listening on port http://localhost:%s...\n", config.Port)

	return http.ListenAndServe(":"+config.Port, nil)
}
