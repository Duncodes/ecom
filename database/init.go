package database

import (
	"database/sql"
	"log"
)

// DB ...
var DB *sql.DB

// InitDB ...
func InitDB() {
	var err error
	DB, err = sql.Open("mysql", "root:duncan@tcp(127.0.0.1:3306)/ecom")
	if err != nil {
		log.Fatal("Error open database ", err)
	}
	// Create tables
	_, err = DB.Exec(`create table if not exists products(id int NOT NULL AUTO_INCREMENT PRIMARY KEY,uuid varchar(255) NOT NULL unique , name varchar(255) unique, photoid varchar(255), description varchar(255), price decimal default 0.0);`)
	_, err = DB.Exec(`create table if not exists users(id int NOT NULL AUTO_INCREMENT PRIMARY KEY , uuid varchar(255) NOT NULL unique, username varchar(255) unique NOT NULL, location varchar(255) NOT NULL, phonenumber varchar(100) unique NOT NULL, email varchar(100),password_hash varchar(255) NOT NULL, joined_on datetime default CURRENT_TIMESTAMP);`)

	if err != nil {
		log.Fatal("Error creating table : ", err)
	}
}
