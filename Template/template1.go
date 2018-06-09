package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// este es para mandar un template	///template, err := template.New("Hola").Parse("hola este es un template")// de este forma se manda a la pagina html
		template, err := template.ParseFiles("./index.html") // este es si le pasas un arch Html
		// metodo ParseFiles para pasar un archivo HTML
		if err != nil {
			fmt.Println("error")
			panic(err)
		}
		template.Execute(w, nil)
	})

	fmt.Println("el servidos escucha en el puerto :8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
