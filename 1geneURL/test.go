package main

import (
    "fmt"
    "net/http"
)

func main() {

    http.HandleFunc("/params", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hola mundosssssssss!!!")
    })
    http.ListenAndServe(":8000", nil)
}


