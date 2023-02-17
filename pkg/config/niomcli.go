package config

import (
	"encoding/json"
	"os"

	"github.com/go-niom/niom/pkg/logger"
)

// niomCli holds niom-app details
type niomCli struct {
	ModuleName string `json:"module_name"`
	AppName    string `json:"app_name"`
	SourceRoot string `json:"sourceRoot"`
	ConfigFile string `json:"configFile"`
}

// GetNiomCliConfig return niom-app details
func GetNiomCliConfig() *niomCli {

	// Reads app details from niom-cli.json
	data, err := os.ReadFile("./niom-cli.json")
	if err != nil {
		logger.Error("File Error: ", err.Error())
		logger.Error("Please run `niom init` ", "")
		return nil
	}

	niomCli := niomCli{}
	// Unmarshals niom-cli.json content to the niomCli variable
	err = json.Unmarshal(data, &niomCli)

	if err != nil {
		logger.Error("Error while reading niom-cli.json", err.Error())
	}

	return &niomCli
}
