package main

import (
	"flag"

	"github.com/loggercode/ecom/database"
	"github.com/loggercode/ecom/server"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	flag.Parse()

	database.InitDB()
	server.StartServer()
}
