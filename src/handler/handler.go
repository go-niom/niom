package handler

import (
	"fmt"

	"github.com/go-niom/niom/pkg/constants"
	"github.com/go-niom/niom/pkg/logger"
	"github.com/go-niom/niom/pkg/terminal"
	"github.com/go-niom/niom/pkg/utils"
	"github.com/go-niom/niom/pkg/watcher"
	"github.com/go-niom/niom/src/engine"
	"github.com/gookit/color"
)

// NewApp initialize niom project with give project name
// The project directory will be populated with scaffolds to manage and run the application
// `niom new 'project_name'` may used to invoke this function
func NewApp(cmd []string) {
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
	Execute(moduleName)
}

// Welcome shows the niom ASCII banner
// Niom Application details
// `niom info` may used to invoke this function
func Welcome() {
	fmt.Println(`
--------------------- Welcome to th world of ----------------------
-------------------------------------------------------------------
---------- ________    ___   ________   _____ ______  -------------
--------- |\   ___  \ |\  \ |\   __  \ |\   _ \  _   \ ------------
--------- \ \  \\ \  \\ \  \\ \  \|\  \\ \  \\\__\ \  \ -----------
---------- \ \  \\ \  \\ \  \\ \  \\\  \\ \  \\|__| \  \ ----------
----------- \ \  \\ \  \\ \  \\ \  \\\  \\ \  \    \ \  \ ---------
------------ \ \__\\ \__\\ \__\\ \_______\\ \__\    \ \__\ --------
------------- \|__| \|__| \|__| \|_______| \|__|     \|__| --------
-------------------------------------------------------------------
---------------------------- Version: ` + constants.AppVersion + ` ------------------------ `)
	fmt.Println("\nTry -h, --help  for usage information.")
}

// `niom -v` may used to invoke this function
func Version() {
	Welcome()
}

// Help shows the list available commands
// `niom help` may used to invoke this function
func Help() {
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

// SwagInit initialize niom project with give project name
func SwagInit(appName string) {
	terminal.CmdExecute(appName, "go", []string{"install", "github.com/swaggo/swag/cmd/swag@latest"}, false)
	SwagExecute(appName)
}

// SwagExecute regenerates swagger documentation
// `niom sg` may used to invoke this function
func SwagExecute(appName string) {
	terminal.CmdExecute(appName, "swag", []string{"init"}, false)
}

// Execute install the dependencies
func Execute(moduleName string) {
	appName := utils.GetAppName(moduleName)
	color.Greenln("\nInstalling dependencies....")
	SwagInit(appName)
	terminal.CmdExecute(appName, "go", []string{"mod", "tidy"}, false)

	fmt.Println(`
🚀  Successfully created project ` + appName + `
👉  Get started with the following commands:`)
	color.Redln("\n\t$ cd " + appName)
	fmt.Println("\t$ niom start:dev")
	color.Greenln("\t$ niom -h\n")
	color.Cyanln("🙏 Thanks for installing Niom 🙏\n")
}

// Info calls Welcome func. to show niom details
func Info() {
	Welcome()
}

// Build executes `go build .`
// `niom build` may used to invoke this function
func Build() {
	terminal.CmdExecute(".", "go", []string{"build", "."}, false)
}

// Install executes `go install .`
// `niom install` may used to invoke this function
func Install(args []string) {
	terminal.CmdExecute(".", "go", []string{"install", "."}, true)
}

// InstallDev executes `go install .` and watch the file changes
// Whenever there is/are any files it reruns the `go install .`
// `niom install:dev` may used to invoke this function
func InstallDev(args []string) {
	watcher.Watch()
	terminal.CmdExecute(".", "go", []string{"install", "."}, true)
	<-terminal.TerminalChannel
}

// Start executes `go run .` by default
// User may specify the path and the command to be run
// `niom start` may used to invoke this function
func Start(args []string) {
	path := "."
	res := utils.ArgsStruct{
		Prefix: "-c=",
		Args:   args,
	}
	if p := utils.ReadArgs("-p=", args); p != "" {
		path = p
	}

	res.AppAndArgs()
	appArgs := res.Result
	cmdArgs := []string{"run", "."}
	app := "go"
	if appArgs.App != "" {
		cmdArgs = appArgs.Args
		app = appArgs.App
	}
	logger.Info(fmt.Sprintf("Running command: %s %s\n", app, cmdArgs))
	terminal.CmdExecute(path, app, cmdArgs, true)
}

// Dev invoke Start func. file watcher mode
// Whenever there is/are any files it recall the Start func
// `niom start:dev` may used to invoke this function
func Dev(args []string) {
	watcher.Watch()
	Start(args)
	// terminal.CmdExecute(".", "go", []string{"run", "."})
	<-terminal.TerminalChannel
}

// Update niom app
func UpdateApp() {
	terminal.CmdExecute(".", "go", []string{"install", "github.com/go-niom/niom@latest"}, false)
}

// TODO future use
func SpinUp(appName string) {
	watcher.Watch()
	terminal.CmdExecute(appName, "go", []string{"run", "."}, false)
}
