package main

import (
	"html/template"
	"log"
	"net/http"
)

var templates = template.Must(template.New("T").ParseGlob("templates_p_error/*.html"))

//var error = template.Must(template.ParseFiles("templates_p_error/error.html"))

func handlerError(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
	//error.Execute(w, nil)
}

func renderTemplate(w http.ResponseWriter, name string, data interface{}) {

	w.Header().Set("Content-Type", "text/html")
	err := templates.ExecuteTemplate(w, name, data)
	if err != nil {
		//http.Error(w, "No es posible retornar el template", http.StatusInternalServerError)
		handlerError(w, http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "index", nil)
	})

	log.Println("El servidor a la escucha en el puerto :8000 ")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
