package models

import (
	"database/sql"
	"fmt"
	"log"
	//database!
	_ "github.com/go-sql-driver/mysql"
)

const usarname string = "root"

const password string = ""

const host string = "localhost"

const port int = 3306

const database string = "project_go_web"

var db *sql.DB

//CreateConnection conexion a una DataBase
func CreateConnection() {
	if connection, err := sql.Open("mysql", generateURL()); err != nil {
		panic(err)
	} else {
		fmt.Println("conex exitosa")
		db = connection
	}

}
func CreateTable(tableName, schema string) {
	if !ExistsTable(tableName) {
		_, err := db.Exec(schema)
		if err != nil {
			log.Println(err)
		}
	}

}
func ExistsTable(tableName string) bool {
	// show tables like 'algo';
	sql := fmt.Sprintf("show tables like '%s'", tableName)
	rows, err := db.Query(sql)
	if err != nil {
		log.Println(err)
	}
	return rows.Next()
}

func Ping() {
	if err := db.Ping(); err != nil {
		panic(err)
	}
}

func CloseConnection() {
	db.Close()
}

func generateURL() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", usarname, password, host, port, database)
}

//<username>:<password>@tcp(<host>)/<database>
