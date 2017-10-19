package config

import (
	"encoding/json"
	"fmt"
	"os"
)

func init() {
	// set config to default
}

// Config the appllication
var Config struct {
	Port             string `json:"port"`
	ServerName       string `json:"server_name"`
	DatabaseHost     string `json:"db_host"`
	DatabaseUserName string `json:"db_user_name"`
	DatabasePassword string `json:"db_password"`
	DatabaseName     string `json:"db_name"`
}

//LoadConfig loads config given a file from json
func LoadConfig(path string) (err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Config)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	return
}
