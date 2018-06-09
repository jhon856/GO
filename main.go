package main

import (
    "fmt"
    "io"
    "net/http"
)

func main() {

    http.HandleFunc("/holamundo", func(w http.ResponseWriter, r *http.Request) {
        io.WriteString(w, "hola mundosss!!!")
    })
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Println("hay una nueva petición")
    io.WriteString(w, "hola mundo")
}
