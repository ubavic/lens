package main

import (
	"fmt"
	"net/http"

	"github.com/ubavic/lens/view"
	"github.com/ubavic/lens/view/form"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	form := form.Form{
		Fields: []form.Field{
			{
				Name:        "Name",
				Type:        form.FieldText,
				Required:    true,
				Description: "Name of the field",
			},
			{
				Name: "Email",
				Type: form.FieldText,
			},
			{
				Name:  "Password",
				Type:  form.FieldText,
				Error: "Error message",
			},
		},
	}

	view.RenderPage(&form, w)
}

func main() {
	port := "8081"

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", handler)

	fmt.Printf("Server listening on port http://localhost:%s...\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
