package config

import (
	"log"

	"github.com/joho/godotenv"
)

// This function dedicated to the read env file from the project directory
// Expect env file path from where env will loaded using godotev package
// LoadAllConfigs set various configs
func LoadAllConfigs(envFile string) {

	//Loads env variable from the env file
	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatalf("can't load %s file. error: %v", envFile, err)
	}

	//Initializing Project db configs
	LoadDBCfg()
}
