package main

import (
	"flag"

	"github.com/Duncodes/ecom/database"
	"github.com/Duncodes/ecom/server"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	flag.Parse()

	database.InitDB()
	server.StartServer()
}
