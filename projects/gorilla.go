package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func YourHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	nombre := vars["nombre"]
	id := vars["id"]
	fmt.Fprintf(w, "los parametros son "+nombre+" "+id+"\n")
	w.Write([]byte("Gorilla!\n"))
}

func main() {
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/usuarios/{nombre}/{id:[0-9]+}", YourHandler).Methods("PUT", "DELETE")

	/*
	   en la linea 22 se puede configurar para la aceptacion de metodos GET PUT POST DELETE
	*/

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))
}
