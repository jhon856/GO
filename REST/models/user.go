package models

import "errors"

//User para  con JSON
type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var users = make(map[int]User)

//SetDefaulUser set
func SetDefaulUser() {
	user := User{Id: 1, Username: "jhonnathan henriquez", Password: "clave123"}
	users[user.Id] = user
}

//GetUser get
func GetUser(userId int) (User, error) {
	if user, ok := users[userId]; ok {
		return user, nil
	}
	return User{}, errors.New("el usuario no se encuentra dentro del mapa")
}
