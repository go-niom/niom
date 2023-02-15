package handler

import (
	"fmt"

	"github.com/go-niom/niom/pkg/constants"
	"github.com/go-niom/niom/pkg/logger"
	"github.com/go-niom/niom/pkg/migrate"
	"github.com/go-niom/niom/pkg/terminal"
	"github.com/go-niom/niom/pkg/utils"
	"github.com/go-niom/niom/pkg/watcher"
	"github.com/go-niom/niom/src/engine"
	"github.com/gookit/color"
)

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

func Version() {
	Welcome()
}

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

func SwagInit(appName string) {
	terminal.CmdExecute(appName, "go", []string{"install", "github.com/swaggo/swag/cmd/swag@latest"}, false)
	SwagExecute(appName)
}

func SwagExecute(appName string) {
	terminal.CmdExecute(appName, "swag", []string{"init"}, false)
}

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

func Build() {
	terminal.CmdExecute(".", "go", []string{"build", "."}, false)
}

func Info() {
	Welcome()
}

func Generate() {
	fmt.Println("Generate app")
}

func SpinUp(appName string) {
	watcher.Watch()
	terminal.CmdExecute(appName, "go", []string{"run", "."}, false)
}

func Dev(args []string) {
	watcher.Watch()
	Start(args)
	// terminal.CmdExecute(".", "go", []string{"run", "."})
	<-terminal.TerminalChannel
}

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

func Migrate(args []string) {
	migrate.Up()
	return
	// path := "."
	// res := utils.ArgsStruct{
	// 	Prefix: "-c=",
	// 	Args:   args,
	// }
	// if p := utils.ReadArgs("-p=", args); p != "" {
	// 	path = p
	// }

	// res.AppAndArgs()
	// appArgs := res.Result
	// cmdArgs := []string{"run", "."}
	// app := "migrate"
	// if appArgs.App != "" {
	// 	cmdArgs = appArgs.Args
	// 	app = appArgs.App
	// }
	// logger.Info(fmt.Sprintf("Running command: %s %s\n", app, cmdArgs))
}

func UpdateApp() {
	terminal.CmdExecute(".", "go", []string{"install", "github.com/go-niom/niom@latest"}, false)
}
