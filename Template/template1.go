package main

import (
	"fmt"
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

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// este es para mandar un template	///template, err := template.New("Hola").Parse("hola este es un template")// de este forma se manda a la pagina html
		template, err := template.ParseFiles("./index.html") // este  si le pasas un arch Html
		// metodo ParseFiles para pasar un archivo HTML

		//otro comandoo
		/*
			ParseGlob("templates/*.gohtml")
		*/
		if err != nil {
			fmt.Println("error")
			panic(err)
		}
		tags := []string{"GO", "Java", "Angular"}
		usuario := Usuario{UserName: "Jhonnathan Henriquez",
			Edad:   26,
			Activo: true,
			Admin:  false,
			Tags:   tags}
		template.Execute(w, usuario)
	})

	fmt.Println("el servidos escucha en el puerto :8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
