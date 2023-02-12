package commands

import (
	"fmt"

	"github.com/go-niom/niom/src/handler"
)

func help(cmd string) {
	switch cmd {
	case "h":
		fallthrough
	case "-h":
		fallthrough
	case "help":
		fallthrough
	case "--help":
		handler.Help()
	}
}

func version(cmd string) {
	switch cmd {
	case "v":
		fallthrough
	case "-v":
		fallthrough
	case "version":
		fallthrough
	case "--version":
		handler.Version()
	}
}

func new(cmd string, args []string) {
	switch cmd {
	case "n":
		fallthrough
	case "-n":
		fallthrough
	case "new":
		fallthrough
	case "--new":
		handler.NewApp(args)
	}
}

func build(cmd string, args []string) {
	switch cmd {
	case "b":
		fallthrough
	case "-b":
		fallthrough
	case "build":
		fallthrough
	case "--build":
		handler.Build()
	}
}

func info(cmd string) {
	switch cmd {
	case "i":
		fallthrough
	case "info":
		handler.Info()
	}
}

func update(cmd string, args []string) {
	switch cmd {
	case "u":
		fallthrough
	case "update":
		handler.UpdateApp()
	}
}

func generate(cmd string, args []string) {
	switch cmd {
	case "g":
		fallthrough
	case "generate":
		handler.Generate()
	}
}

func swagger(cmd string, args []string) {
	switch cmd {
	case "sg":
		fallthrough
	case "swagger":
		handler.SwagExecute(".")
	}
}

func start(cmd string, args []string) {
	switch cmd {
	case "start:dev":
		handler.Dev()
	case "start":
		handler.Start()
	default:
		fmt.Printf("Command not available %s\n", cmd)
		handler.Help()
	}
}

func Commands(args []string) {
	cmd := args[1]
	help(cmd)
	version(cmd)
	info(cmd)
	new(cmd, args)
	build(cmd, args)
	generate(cmd, args)
	update(cmd, args)
	start(cmd, args)
	swagger(cmd, args)
}
