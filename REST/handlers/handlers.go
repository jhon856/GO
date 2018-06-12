package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"../models"
	"github.com/gorilla/mux"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Se obtiene todos los usuarios!")
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{Id: 1, Username: "jhonnathan", Password: "clave123"}
	w.Header().Set("Content-type", "application/json")
	vars := mux.Vars(r)
	userId, _ := strconv.Atoi(vars["id"])
	user, err := models.GetUser(userId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
	// estas serializando con esto marshal
	output, _ := json.Marshal(&user)
	fmt.Fprintf(w, string(output))
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Se crea un usuario")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Se actualiza un usuario")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Se elimina usuario")
}
