package commands

import (
	"encoding/json"
	"fmt"

	"github.com/go-niom/niom/pkg/constants"
	"github.com/go-niom/niom/pkg/utils"

	"github.com/go-niom/niom/src/handler"
)

type promptMsg struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	Title string `json:"title"`
}

func InitNiomCli() {

	fmt.Println(`This utility will walk you through creating a niom-cli.json file.
It only covers the most common items, and tries to guess sensible defaults.`)

	var niomCliPrompt []promptMsg
	prompts := map[string]string{}
	err := json.Unmarshal([]byte(constants.CliPrompt), &niomCliPrompt)
	if err != nil {
		fmt.Printf("Error while parsing prompt: %s", err.Error())
	}

	for _, nc := range niomCliPrompt {
		value := utils.UserPrompt(fmt.Sprintf("%s?(%s)", nc.Title, nc.Value))
		if value == "" {
			value = nc.Value
		}
		prompts[nc.Key] = value
	}
	bt, err := json.MarshalIndent(prompts, "", "	")
	if err != nil {
		panic(err)
	}
	println(string(bt))

}

// Commands check and intercept user entered commands
// As per the user this function redirects to the function to carry the task
func Commands(args []string) {
	cmd := args[1]

	switch cmd {
	case "init":
		InitNiomCli()
	case "h", "-h", "help", "--help":
		handler.Help()
	case "v", "-v", "version", "--version":
		handler.Version()
	case "n", "-n", "new", "--new":
		handler.NewApp(args)
	case "b", "-b", "build", "--build":
		handler.Build()
	case "i", "info":
		handler.Info()
	case "u", "update":
		handler.UpdateApp()
	case "g", "generate":
		generate(cmd, args)
	case "sg", "swagger":
		handler.SwagExecute(".")
	case "start:dev":
		handler.Dev(args[2:])
	case "start":
		handler.Start(args[2:])
	case "install":
		handler.Install(args[2:])
	case "install:dev":
		handler.InstallDev(args[2:])
	case "kill", "kl":
		handler.KillPort(args[2:])
	case "migration", "mg":
		migrations(args[2:])

	default:
		fmt.Printf("Command not available %s\n", cmd)
	}

}
