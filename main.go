package main

import (
	"fmt"
	"os"

	"github.com/go-niom/niom/pkg/constants"
	"github.com/go-niom/niom/pkg/terminal"
	"github.com/go-niom/niom/pkg/utils"
	"github.com/go-niom/niom/pkg/watcher"
	"github.com/go-niom/niom/src/engine"
	"github.com/gookit/color"
)

func help() {
	// info|i                                          Display Niom project details.
	// update|u [options]                              Update Niom dependencies.
	println(`OPTIONS:
  -v, --version                                   Output the current version.
  -h, --help                                      Output usage information.
	`)

	print(`COMMANDS:
  new|n [options] [name]                          Generate Niom application.
  build [options] [app]                           Build Niom application.
  start:dev [options] [app]                       Run app rebuild/watch mode.
  update|u [options]                              Update Niom CLI.
  swagger|sg [options]                            Generate Swagger docs
  generate|g [options] <schematic> [name] [path]  Generate a Niom element.`)
	print(`
	┌───────────────────────────────────────────────────────────────────┐
	│                       Available schematics:                       │
	├─────────────┬─────────────┬───────────────────────────────────────┤
	│ NAME        │ ALIAS       │  DESCRIPTION                          │
	├─────────────┼─────────────┼───────────────────────────────────────┤
	│ resource    │ res         │  Generate a new CRUD resource         │
	│ controller  │ co          │  Generate a controller declaration    │
	│ service     │ s           │  Generate a service declaration       │
	│ router      │ ro          │  Generate a router                    │
	│ interface   │ in          │  Generate an interface                │
	│ middleware  │ mi          │  Generate a middleware declaration    │
	│ model       │ mo          │  Generate a model                     │
	│ dto         │ dto         │  Generate a dto                       │
	└─────────────┴─────────────┴───────────────────────────────────────┘
	Example: niom g res user` + "\n")
}

func swagInit(appName string) {
	terminal.CmdExecute(appName, "go", []string{"install", "github.com/swaggo/swag/cmd/swag@latest"})
	swagExecute(appName)
}

func swagExecute(appName string) {
	terminal.CmdExecute(appName, "swag", []string{"init"})
}

func execute(moduleName string) {

	appName := utils.GetAppName(moduleName)
	color.Greenln("\nInstalling dependencies....")
	swagInit(appName)
	terminal.CmdExecute(appName, "go", []string{"mod", "tidy"})

	fmt.Println(`
🚀  Successfully created project ` + appName + `
👉  Get started with the following commands:`)
	color.Redln("\n\t$ cd " + appName)
	fmt.Println("\t$ niom start:dev")
	color.Greenln("\t$ niom -h\n")
	color.Cyanln("🙏 Thanks for installing Niom 🙏\n")

	// spinUp(appName)
}

func newApp(cmd []string) {
	if len(cmd) < 3 {
		fmt.Print(`
Example usage:
        'niom new app_name' to initialize a v0 or v1 module
        'niom new example.com/m' to initialize a v0 or v1 module
        'niom new example.com/m/v2' to initialize a v2 module
`)
		return
	}
	moduleName := cmd[2]
	engine.CreateInitialFiles(moduleName)
	execute(moduleName)
}

func build() {
	terminal.CmdExecute(".", "go", []string{"build", "."})
}
func info() {
	welcome()
}

func g() {
	fmt.Println("Generate app")
}

func spinUp(appName string) {
	watcher.Watch()
	terminal.CmdExecute(appName, "go", []string{"run", "."})

}

func dev() {
	watcher.Watch()
	terminal.CmdExecute(".", "go", []string{"run", "."})
	<-terminal.TerminalChannel
}

func start() {
	terminal.CmdExecute(".", "go", []string{"run", "."})
}

func version() {
	welcome()
}

func updateApp() {
	terminal.CmdExecute(".", "go", []string{"install", "github.com/go-niom/niom@latest"})
}

func welcome() {

	fmt.Println(`
--------------------- Welcome to th world of ----------------------
-------------------------------------------------------------------
----------________    ___   ________   _____ ______   -------------
--------- |\   ___  \ |\  \ |\   __  \ |\   _ \  _   \ ------------
--------- \ \  \\ \  \\ \  \\ \  \|\  \\ \  \\\__\ \  \ -----------
---------- \ \  \\ \  \\ \  \\ \  \\\  \\ \  \\|__| \  \ ----------
----------- \ \  \\ \  \\ \  \\ \  \\\  \\ \  \    \ \  \ ---------
------------ \ \__\\ \__\\ \__\\ \_______\\ \__\    \ \__\ --------
------------- \|__| \|__| \|__| \|_______| \|__|     \|__| --------
-------------------------------------------------------------------
---------------------------- Version: ` + constants.AppVersion + ` ------------------------ `)

	// 	println(
	// 		`
	// ---------------------------- Welcome to the world of ------------------------------
	// `)
	// 	myFigure := figure.NewFigure("...STUPA ...", "larry3d", true)
	// 	myFigure.Print()
	// 	println(`
	// -----------------------------------------------------------------------------------
	// --------------------------------- Version: 0.1 ------------------------------------`)
	fmt.Println("\nTry -h, --help  for usage information.")

}

func commands(args []string) {

	cmd := args[1]
	switch cmd {

	case "h":
		fallthrough
	case "-h":
		fallthrough
	case "help":
		fallthrough
	case "--help":
		help()

	case "v":
		fallthrough
	case "-v":
		fallthrough
	case "version":
		fallthrough
	case "--version":
		version()
	case "n":
		fallthrough
	case "new":
		newApp(args)
	case "b":
		fallthrough
	case "build":
		build()
	case "i":
		fallthrough
	case "info":
		info()
	case "u":
		fallthrough
	case "update":
		updateApp()
	case "g":
		fallthrough
	case "generate":
		g()
	case "start:dev":
		dev()
	case "start":
		start()
	case "sg":
		fallthrough
	case "swagger":
		swagExecute(".")
	default:
		fmt.Printf("Command not available %s\n", cmd)
		help()
	}
}

func main() {
	args := os.Args
	if len(args) == 1 {
		welcome()
		return
	}
	commands(args)
}
