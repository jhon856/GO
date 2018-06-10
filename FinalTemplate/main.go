package main

import (
	"html/template"
	"log"
	"net/http"
)

var templates = template.Must(template.ParseFiles("templates/head.html"))

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := templates.ExecuteTemplate(w, "head.html", nil)

		if err != nil {
			panic(err)
		}
	})

	log.Fatal("el servidos escucha en el puerto :8000")
	log.Fatal(http.ListenAndServe(":8000", nil))

}
