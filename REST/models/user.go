package models

import "errors"

//User para  con JSON
type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

const UserSchema string = `CREATE TABLE user (
    id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(30) NOT NULL,
    password VARCHAR(64) NOT NULL,
    emial VARCHAR(40),
    created_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)`

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
