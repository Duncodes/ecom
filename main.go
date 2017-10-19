package main

import (
	"fmt"
	"os"

	"github.com/Duncodes/ecom/cmd"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// flag.Parse()

	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
