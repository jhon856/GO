package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hola mundo Server")
	})
	// %%%%%%%%%%%%%%%%%%%%%%%%	mux := http.ServeMux()
	// %%%%%%%%%%%%%%%%%%%%%%%	mux.Handle("path","handleralgo asi    &user{name:jhon}")
	server := &http.Server{
		Addr:    "localhost:8000",
		Handler: nil, // si es nil usan DefaulServerMUX####
		//	Handler: mux, // serverMux   %%%%%%%%%%%%%%%%%%%%%%%
	}

	log.Fatal(server.ListenAndServe())

}
