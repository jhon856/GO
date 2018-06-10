package main

import (
	"log"
	"net/http"
	"text/template"
)

var templates = template.Must(template.New("T").ParseGlob("ayuda/*.html"))

// para algunos archivos este ParseFiles.. para muchos ParseGlob
// ParseFiles("./index.html", "./footer.html", "./registro.html"))

func RenderTemplate(w http.ResponseWriter, name string, data interface{}) {
	w.Header().Set("Content-Type", "text/html")
	err := templates.ExecuteTemplate(w, name, data)
	// este  si le pasas un arch Html

	if err != nil {
		http.Error(w, "no es posible cargar los templates", http.StatusInternalServerError)

	}
}
func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		RenderTemplate(w, "registrox", nil)
	})

	log.Fatal("el servidos escucha en el puerto :8000")
	log.Fatal(http.ListenAndServe(":8000", nil))

}
