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
	_, err = DB.Exec(`create table if not exists category(id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
					uuid varchar(100) NOT NULL unique , name varchar(100) unique, description text ,
					picture varchar(100));`)
	if err != nil {
		log.Fatal("Error creating table : ", err)
	}
	_, err = DB.Exec(`create table if not exists products(id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
					uuid varchar(100) NOT NULL unique , name varchar(255) unique, photoid varchar(255),
					description varchar(255), price decimal default 0.0 , productstock decimal default 0.0,
					update_date datetime default CURRENT_TIMESTAMP, quantitypreunit int default 1,
					categoryid int NOT NULL,FOREIGN KEY (categoryid) REFERENCES category (id));`)
	if err != nil {
		log.Fatal("Error creating table : ", err)
	}
	_, err = DB.Exec(`create table if not exists users(id int NOT NULL AUTO_INCREMENT PRIMARY KEY ,
					uuid varchar(100) NOT NULL unique, username varchar(255) unique NOT NULL,
					location varchar(255) NOT NULL, phonenumber varchar(100) unique NOT NULL,
					email varchar(100),password_hash varchar(255) NOT NULL, joined_on datetime default CURRENT_TIMESTAMP);`)
	if err != nil {
		log.Fatal("Error creating table : ", err)
	}
	_, err = DB.Exec(`create table if not exists payment(id int NOT NULL AUTO_INCREMENT PRIMARY KEY ,
					uuid varchar(100) NOT NULL unique , paymenttype varchar(100) unique,allowed bool default true);`)
	if err != nil {
		log.Panic("Error creating table : ", err)
	}
	_, err = DB.Exec(`create table if not exists orders(id int NOT NULL AUTO_INCREMENT PRIMARY KEY ,
					uuid varchar(100) NOT NULL unique , customerid int NOT NULL, paymentid int NOT NULL,
					shippingadress varchar(100), paid bool default false, fulfilled bool default false,
					timestamp datetime default CURRENT_TIMESTAMP, FOREIGN KEY(customerid) REFERENCES users(id) ,
					FOREIGN KEY(paymentid) REFERENCES payment(id));`)
	if err != nil {
		log.Panic("Error creating table : ", err)
	}
	_, err = DB.Exec(`create table if not exists orderdetails(id int NOT NULL AUTO_INCREMENT PRIMARY KEY ,
					uuid varchar(100) NOT NULL unique , orderid int NOT NULL, productid int NOT NULL ,
					quantity int NOT NULL default 1, price decimal NOT NULL , total decimal NOT NULL,
					FOREIGN KEY(orderid) REFERENCES orders(id) , FOREIGN KEY(productid) REFERENCES products(id));`)

	if err != nil {
		log.Fatal("Error creating table : ", err)
	}
}
