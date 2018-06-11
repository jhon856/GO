package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	http.HandleFunc("/holamundo", func(s http.ResponseWriter, c *http.Request) {
		io.WriteString(s, "hola mundosssssssssssssssssssssssss")
	})
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hay una nueva peticion")
	io.WriteString(w, "hola mundo")
}
