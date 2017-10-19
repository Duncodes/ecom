package main

import (
	"flag"
	"log"

	"github.com/Duncodes/ecom/database"
	"github.com/Duncodes/ecom/econfig"
	"github.com/Duncodes/ecom/server"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	flag.Parse()
	err := econfig.LoadConfig("config.json")

	if err != nil {
		log.Println(err)
	}

	log.Println(econfig.Config.Port)
	database.InitDB()
	server.StartServer()
}
