package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("mysql", "root:duncan@tcp(127.0.0.1:3306)/ecom")
	if err != nil {
		log.Fatal("Error open database ", err)
	}
	// Create tables
	_, err = DB.Exec(`create table if not exists items(id int NOT NULL AUTO_INCREMENT PRIMARY KEY,uuid varchar(255) NOT NULL unique , name varchar(255) unique, photoid varchar(255), description varchar(255), price decimal default 0.0);`)
	if err != nil {
		log.Fatal("Error creation table table : ", err)
	}
}
