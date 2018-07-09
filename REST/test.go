package main

import (
	"fmt"

	"./models"
)

func main() {
	models.CreateConnection()
	models.Ping()
	result := models.ExistsTable("user")
	fmt.Println(result)
	models.CreateTable("user", models.UserSchema)
	models.CloseConnection()
}
