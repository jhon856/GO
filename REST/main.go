package main

import (
	"log"
	"net/http"

	"./handlers"
	"./models"
	"github.com/gorilla/mux"
)

func main() {
	mux := mux.NewRouter()

	models.SetDefaulUser()

	mux.HandleFunc("/api/v1/users/", handlers.GetUsers).Methods("GET")
	mux.HandleFunc("/api/v1/users/{id:[0-9]+}", handlers.GetUser).Methods("GET")
	mux.HandleFunc("/api/v1/users/", handlers.CreateUser).Methods("POST")
	mux.HandleFunc("/api/v1/users/{id:[0-9]+}", handlers.UpdateUser).Methods("PUT")
	mux.HandleFunc("/api/v1/users/{id:[0-9]+}", handlers.DeleteUser).Methods("DELETE")

	log.Println("el servidor esta a la eschua")
	log.Fatal(http.ListenAndServe(":8000", mux))
}
