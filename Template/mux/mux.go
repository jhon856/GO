package mux

import (
	"net/http"
)

type customerHandler func(http.ResponseWriter, *http.Request)

//MuxFacilito s
type MuxFacilito struct {
	rutasfacilitas map[string]customerHandler
}

//AddMux funcion
func (this *MuxFacilito) AddMux(ruta string, fn customerHandler) {
	this.rutasfacilitas[ruta] = fn
}
func (this *MuxFacilito) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if fn, ok := this.rutasfacilitas[r.URL.Path]; ok {
		fn(w, r)
	} else {
		http.NotFound(w, r)
	}

}

func CreateMux() *MuxFacilito { //puntero
	newMap := make(map[string]customerHandler)
	return &MuxFacilito{newMap} //valor
}
