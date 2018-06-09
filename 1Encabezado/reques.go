package main
import (
	"fmt"
	"net/http"
	"log"
)
func main() {

    http.HandleFunc("/params", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Header)
		access_token := r.Header.Get("access_token")
		if len(access_token)!= 0{
			fmt.Println(access_token)
		}
		///para agregar encabezadosss
		r.Header.Set("nombre","valor")
		fmt.Println(r.Header)

		fmt.Fprintf(w, "hola mundo")
    })
    http.HandleFunc("/otro", func(w http.ResponseWriter, r *http.Request) {
    })
    log.Fatal(http.ListenAndServe(":8000", nil))
}

