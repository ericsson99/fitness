package main

import (
	"net/http"
	"text/template"
)

func main() {
	router := http.NewServeMux()

	tmpl, err := template.ParseFiles("templates/layout.html")
	if err != nil {
		panic("template error")
	}

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, struct{}{})
	})

	http.ListenAndServe(":8080", router)
}
