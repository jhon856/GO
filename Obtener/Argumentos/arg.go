package main

import (
	"fmt"
	"log"
	"net/http"
)

//get, post, put, delete
func main() {

	http.HandleFunc("/params", func(w http.ResponseWriter, r *http.Request) {
		//	fmt.Println(r.URL)// este muestra solo el path ejemplo "/"
		//fmt.Println(r.URL.RawQuery) // aca se puede usar Split y un ciclo For es la lectura ServletRequest de los clientes JAva Jboss
		fmt.Println(r.URL.Query()) //map
		name := r.URL.Query().Get("name")
		if len(name) != 0 {
			fmt.Println(name)
		}
	})
	log.Fatal(http.ListenAndServe(":8000", nil))
}
