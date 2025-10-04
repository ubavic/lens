package server

import (
	"fmt"
	"net/http"
)

func StartServer(port string) error {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", homeHandler)

	fmt.Printf("Server listening on port http://localhost:%s...\n", port)

	return http.ListenAndServe(":"+port, nil)
}
