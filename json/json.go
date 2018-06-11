package main

import (
	"encoding/json"
	"net/http"
)

//Go no exporta variables en minuscula, si se desea eso >>>>  string ´json:titulo´
type Curso struct {
	Title        string
	NumeroVideos int
}
type Cursos []Curso

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cursos := Cursos{
			Curso{"Curso de Go", 30},
			Curso{"Curso de Go Web", 12},
			Curso{"Curso JAva", 44},
		}
		json.NewEncoder(w).Encode(cursos)
	})
	http.ListenAndServe(":8000", nil)
}
