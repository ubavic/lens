package server

import (
	"log"
	"net/http"

	"github.com/ubavic/lens/view"
	"github.com/ubavic/lens/view/navigation"
)

type ServerConfig struct {
	Port  string
	Pages []view.Page
}

func StartServer(config ServerConfig) error {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("GET /static/", http.StripPrefix("/static/", fs))

	prepareNavigation(config)

	indexRoutes = make([]string, 0, len(config.Pages))
	for _, p := range config.Pages {
		p.ResolveIds()
		http.HandleFunc("POST /"+p.Id+"/", postHandler(p))
		http.HandleFunc("/"+p.Id+"/", pageHandler(p))
		indexRoutes = append(indexRoutes, p.Id)
	}

	http.HandleFunc("/", indexHandler)

	log.Printf("Server listening on port http://localhost:%s...\n", config.Port)

	return http.ListenAndServe(":"+config.Port, nil)
}

func prepareNavigation(config ServerConfig) {

	routes := make([]navigation.Entry, 0, len(config.Pages))

	for _, p := range config.Pages {
		routes = append(routes, navigation.Entry{
			Name: p.Name,
			Path: "/" + p.Id + "/",
		})
	}

	navigation.SetEntries(routes)
}
