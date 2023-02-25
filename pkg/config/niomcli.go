package config

import (
	"encoding/json"
	"os"

	"github.com/go-niom/niom/pkg/logger"
)

// NiomCli holds niom-app details
type NiomCli struct {
	ModuleName string `json:"module_name"`
	AppName    string `json:"app_name"`
	SourceRoot string `json:"source_root"`
	ConfigFile string `json:"config_file"`
}

// GetNiomCliConfig return niom-app details
func GetNiomCliConfig() *NiomCli {

	// Reads app details from niom-cli.json
	data, err := os.ReadFile("./niom-cli.json")
	if err != nil {
		logger.Error("File Error: ", err.Error())
		logger.Error("Please run `niom init` ", "")
		return nil
	}

	niomCli := NiomCli{}
	// Unmarshals niom-cli.json content to the NiomCli variable
	err = json.Unmarshal(data, &niomCli)

	if err != nil {
		logger.Error("Error while reading niom-cli.json", err.Error())
	}

	return &niomCli
}
