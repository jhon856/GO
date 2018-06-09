package main

import (
	"net/http"
)

func main() {

	fileServer := http.FileServer(http.Dir("public"))
	http.Handle("/", http.StripPrefix("/", fileServer)) //	http.Handle("/public/", con este codigo haria la busquedaen esa ruta o path
	http.ListenAndServe(":8000", nil)
}

/*
func main() {

	http.HandleFunc("/holamundo", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html") //r.URL.Path[1:] este es para accedes a cualquier directorio
		io.WriteString(w, "hola mundosssssssssssssssssssssssss")
	})
	http.ListenAndServe(":8000", nil)
}

*/
