package commands

import (
	"fmt"

	"github.com/go-niom/niom/src/handler"
)

func Commands(args []string) {
	cmd := args[1]

	switch cmd {
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
	default:
		fmt.Printf("Command not available %s\n", cmd)
	}

}
