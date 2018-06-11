package main

import (
	"html/template"
	"log"
	"net/http"
)

//Usuario algo
type Usuario struct {
	UserName string
	Edad     int
	Activo   bool
	Admin    bool
	Tags     []string
}

func (this Usuario) TienePermisoAdmin(llave string) bool {
	return this.Activo && this.Admin && llave == "si"
}

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
	staticFiles := http.FileServer(http.Dir("css"))
	mux := http.NewServeMux()
	mux.Handle("/css/", http.StripPrefix("/css/", staticFiles))
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tags := []string{"GO", "Java", "Angular"}
		usuario := Usuario{UserName: "Jhonnathan Henriquez",
			Edad:   26,
			Activo: true,
			Admin:  true,
			Tags:   tags}
		RenderTemplate(w, "registrox", usuario)
	})

	log.Println("el servidos escucha en el puerto :8000")
	log.Fatal(http.ListenAndServe(":8000", mux))

}
