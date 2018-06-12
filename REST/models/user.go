package models

import "errors"

//User para  con JSON
type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}
type Users []User

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

//GetUsers get
func GetUsers() Users {
	list := Users{}
	for _, user := range users {
		list = append(list, user)
	}
	return list
}

func SaveUser(user User) User { // si fuera con base de datos. el parametro seria un apuntador &
	user.Id = len(users) + 1
	users[user.Id] = user
	return user
}

func UpdateUser(user User, username string, passwor string) User {
	user.Username = username
	user.Password = passwor
	users[user.Id] = user
	return user
}
func DeleteUser(id int) {
	delete(users, id)
}
