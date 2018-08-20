package utils

import (
	"encoding/json"
	"log"
	"os"
)

type (
	configuration struct {
		Server, DbHost, DbUser, DbPwd, DbName string
		DbPort, LogLevel                      int
	}
)

// AppConfig holds the configuration values from config.json file
var AppConfig configuration

// Reads config.json and decode into AppConfig
func loadConfig() {
	file, err := os.Open("config.json")
	defer file.Close()
	if err != nil {
		log.Fatalf("[loadConfig]: %s\n", err)
	}

	decoder := json.NewDecoder(file)
	AppConfig = configuration{}
	err = decoder.Decode(&AppConfig)
	if err != nil {
		log.Fatalf("[loadAppConfig]: %s\n", err)
	}
}
