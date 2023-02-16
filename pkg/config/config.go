package config

import (
	"log"

	"github.com/joho/godotenv"
)

// LoadAllConfigs set various configs
func LoadAllConfigs(envFile string) {

	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatalf("can't load .env file. error: %v", err)
	}

	LoadDBCfg()
}
