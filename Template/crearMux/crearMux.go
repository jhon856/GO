package crearMux

import (
	"fmt"
	"log"
	"net/http"
)

type CustomerHandler func(http.ResponseWriter, *http.Request)

//MuxFacilito s
type MuxFacilito struct {
	rutasfacilitas map[string]CustomerHandler
}

//AddMux funcion
func (this *MuxFacilito) AddMux(ruta string, fn CustomerHandler) {
	this.rutasfacilitas[ruta] = fn
}
func (this *MuxFacilito) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fn := this.rutasfacilitas[r.URL.Path]
	fn(w, r)
}

func main() {
	newMap := make(map[string]CustomerHandler)
	mux := &MuxFacilito{rutasfacilitas: newMap}
	mux.AddMux("/hola", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hola mundo Server desde una funcion")
	})
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}
