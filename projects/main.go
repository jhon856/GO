package main

import (
	"fmt"
	"log"
	"net/http"

	"./mux"
)

func main() {

	mux := mux.CreateMux()
	mux.AddMux("/hola", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hola mundo Server desde una funcion")
	})
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}
