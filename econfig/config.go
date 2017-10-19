package econfig

import (
	"encoding/json"
	"os"
)

// Config the appllication
var Config struct {
	Port             string `json:"port"`
	ServerName       string `json:"server_name"`
	DatabaseHost     string `json:"db_host"`
	DatabasePassword string `json:"db_password"`
}

//LoadConfig loads config given a file from json
func LoadConfig(path string) (err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	var data []byte
	file.Read(data)
	err = json.Unmarshal(data, Config)
	return
}
