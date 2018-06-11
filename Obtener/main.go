package main

import (
	"fmt"
	"log"
	"net/http"
)

//get, post, put, delete
func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("El Metodo es +" + r.Method)

		switch r.Method {
		case "GET":
			fmt.Fprintf(w, "hola mundo con el metodo Get##")
		case "POST":
			fmt.Fprintf(w, "hola mundo con el metodo Post##")
		case "PUT":
			fmt.Fprintf(w, "hola mundo con el metodo Put##")
		case "DELETE":
			fmt.Fprintf(w, "hola mundo con el metodo Delete##")
		default:
			http.Error(w, "Este es un Error", http.StatusBadRequest)
		}

		fmt.Fprintf(w, "hola mundo")
	})
	log.Fatal(http.ListenAndServe(":8000", nil))
}
