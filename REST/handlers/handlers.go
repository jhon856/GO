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
	models.SendData(w, models.GetUsers())
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	//user := models.User{Id: 1, Username: "jhonnathan", Password: "clave123"} para ejemplos
	//w.Header().Set("Content-type", "application/json") esto se fue a response########
	// &&&&&& este es para UPDATEEEEEEEEEEEEEEEEE
	//vars := mux.Vars(r)
	//userId, _ := strconv.Atoi(vars["id"])

	if user, err := getUserByRequest(r); err != nil {
		models.SendNotFound(w)
	} else {

		models.SendData(w, user)
	}
	// estas serializando con esto marshal
	/* esto se fue al response#############
	output, _ := json.Marshal(&response)
	fmt.Fprintf(w, string(output))
	*/
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&user); err != nil {
		models.SendUnprocessableEntity(w)
	} else {

		models.SendData(w, models.SaveUser(user))
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Se actualiza un usuario")
	user, err := getUserByRequest(r)
	if err != nil {
		models.SendNotFound(w)
		return
	}
	userResponse := models.User{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userResponse); err != nil {
		models.SendUnprocessableEntity(w)
		return
	}
	user = models.UpdateUser(user, userResponse.Username, userResponse.Password)
	models.SendData(w, user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Se elimina usuario")
	if user, err := getUserByRequest(r); err != nil {
		models.SendNotFound(w)
	} else {
		models.DeleteUser(user.Id)
		models.SendNotContent(w)
	}
}

//getUserByRequest sirve para todos pero es privada
func getUserByRequest(r *http.Request) (models.User, error) {
	vars := mux.Vars(r)
	userId, _ := strconv.Atoi(vars["id"])
	if user, err := models.GetUser(userId); err != nil {
		return user, err
	} else {
		return user, nil
	}
}

/*   &&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&      CURL      &&&&&&&&&&&&&&&&&
$ curl http://localhost:8000/api/v1/users/ -X POST -d '{"username" :"loco","password" : "passwwww123"}' -H "Content-Type: application/json"
peticion post created

*/
