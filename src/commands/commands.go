package commands

import (
	"fmt"

	"github.com/go-niom/niom/pkg/utils"
	"github.com/go-niom/niom/src/handler"
)

func help(cmd string) bool {
	if utils.ListContains(cmd, []string{"h", "-h", "help", "--help"}) {
		handler.Help()
		return true
	}
	return false
}

func version(cmd string) bool {
	if utils.ListContains(cmd, []string{"v", "-v", "version", "--version"}) {
		handler.Version()
		return true
	}
	return false
}

func new(cmd string, args []string) bool {
	if utils.ListContains(cmd, []string{"n", "-n", "new", "--new"}) {
		handler.NewApp(args)
		return true
	}
	return false
}

func build(cmd string, args []string) bool {
	if utils.ListContains(cmd, []string{"b", "-b", "build", "--build"}) {
		handler.Build()
		return true
	}
	return false
}

func info(cmd string) bool {
	if utils.ListContains(cmd, []string{"i", "info"}) {
		handler.Info()
		return true
	}
	return false
}

func update(cmd string, args []string) bool {
	if utils.ListContains(cmd, []string{"u", "update"}) {
		handler.UpdateApp()
		return true
	}
	return false
}

func generate(cmd string, args []string) bool {
	if utils.ListContains(cmd, []string{"g", "generate"}) {
		handler.Generate()
		return true
	}
	return false
}

func swagger(cmd string, args []string) bool {
	if utils.ListContains(cmd, []string{"sg", "swagger"}) {
		handler.SwagExecute(".")
		return true
	}
	return false
}

func start(cmd string, args []string) {
	switch cmd {
	case "start:dev":
		handler.Dev()
	case "start":
		handler.Start()
	default:
		fmt.Printf("Command not available %s\n", cmd)
	}
}

func Commands(args []string) {
	cmd := args[1]
	if help(cmd) {
		return
	}
	if version(cmd) {
		return
	}
	if info(cmd) {
		return
	}
	if new(cmd, args) {
		return
	}
	if build(cmd, args) {
		return
	}
	if generate(cmd, args) {
		return
	}
	if update(cmd, args) {
		return
	}
	if swagger(cmd, args) {
		return
	}
	start(cmd, args)
}
