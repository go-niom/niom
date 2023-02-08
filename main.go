package main

import (
	"errors"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-niom/niom/pkg/terminal"
	"github.com/go-niom/niom/pkg/watcher"
	"github.com/go-niom/niom/src/engine"
	"github.com/go-niom/niom/utils"
	"github.com/gookit/color"
)

type TemplateArgs struct {
	Name          string
	NameLowerCase string
}

func renderWriteToFile(tmpl string, func_name string, file_name string) {

	td := TemplateArgs{strings.Title(func_name), func_name}
	t, err := template.New("name").Parse(tmpl)
	if err != nil {
		fmt.Println("errrror", err)
	}
	f, err := os.Create(file_name)
	if err != nil {
		fmt.Println("create file: ", err)
		return
	}
	err = t.Execute(f, td)
	if err != nil {
		panic(err)
	}
}

func ensureDir(dirName string) error {
	err := os.MkdirAll(dirName, os.ModePerm)
	if err == nil {
		return nil
	}
	if os.IsExist(err) {
		// check that the existing path is a directory
		info, err := os.Stat(dirName)
		if err != nil {
			return err
		}
		if !info.IsDir() {
			return errors.New("path exists but is not a directory")
		}
		return nil
	}
	return err
}

func getFileDirectory(filetype string, file_path string, filename string, isBase bool) string {
	mdir := filepath.Join("./src", file_path, filename)

	dir := filepath.Join(mdir, filetype)

	if isBase {
		dir = mdir
	}
	fmt.Println(dir)
	if err := ensureDir(dir); err != nil {
		fmt.Println("Directory creation failed with error: " + err.Error())
		os.Exit(1)
	}
	return fmt.Sprintf("%s/%s.%s.go", dir, filename, filetype)
}

// func controller(file_name string, file_path string) {
// 	renderWriteToFile(utils.ControllerTmpl, file_name, getFileDirectory("controllers", file_path, file_name, true))
// }

// func interfaces(file_name string, file_path string) {
// 	renderWriteToFile(utils.InterfaceTmpl, file_name, getFileDirectory("interfaces", file_path, file_name, false))
// }

// func model(file_name string, file_path string) {
// 	renderWriteToFile(utils.ModelTmpl, file_name, getFileDirectory("models", file_path, file_name, false))
// }

// func router(file_name string, file_path string) {
// 	renderWriteToFile(utils.RouterTmpl, file_name, getFileDirectory("router", file_path, file_name, true))
// }

// func service(file_name string, file_path string) {
// 	renderWriteToFile(utils.ServiceTmpl, file_name, getFileDirectory("service", file_path, file_name, true))
// }

// func dto(file_name string, file_path string) {
// 	renderWriteToFile(utils.DtoTmpl, file_name, getFileDirectory("dto", file_path, file_name, false))
// }

func res(file_name string, file_path string) {
	// controller(file_name, file_path)
	// interfaces(file_name, file_path)
	// model(file_name, file_path)
	// router(file_name, file_path)
	// service(file_name, file_path)
	// dto(file_name, file_path)
}

func generate(cmd string, name string) {
	// df := strings.Split(name, "/")
	// file_name := df[len(df)-1]
	// file_path := strings.Join(df[0:len(df)-1], "/")

	// switch cmd {
	// case "res":
	// 	{
	// 		res(file_name, file_path)
	// 		break
	// 	}
	// case "co":
	// 	{
	// 		controller(file_name, file_path)
	// 		break
	// 	}
	// case "se":
	// 	{
	// 		service(file_name, file_path)
	// 		break
	// 	}
	// case "ro":
	// 	{
	// 		router(file_name, file_path)
	// 		break
	// 	}
	// case "mo":
	// 	{
	// 		model(file_name, file_path)
	// 		break
	// 	}
	// case "in":
	// 	{
	// 		interfaces(file_name, file_path)
	// 		break
	// 	}
	// case "dto":
	// 	{
	// 		dto(file_name, file_path)
	// 		break
	// 	}
	// }

}

func Contains(collection []string, element string) bool {
	for _, item := range collection {
		if item == element {
			return true
		}
	}

	return false
}

func help() {

	println(`OPTIONS:
  -v, --version                                   Output the current version.
  -h, --help                                      Output usage information.
	`)

	print(`COMMANDS:
  new|n [options] [name]                          Generate Niom application.
  build [options] [app]                           Build Niom application.
  dev [options] [app]                           Run app rebuild/watch mode.
  info|i                                          Display Niom project details.
  update|u [options]                              Update Niom dependencies.
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

func execute(moduleName string) {
	appName := utils.GetAppName(moduleName)

	println(`
🚀  Successfully created project ` + appName + `
👉  Get started with the following commands:`)
	color.Redln("\n\t$ cd " + appName)
	color.Println("\t$ go mod tidy")
	color.Greenln("\t$ go run main.go\n")
	color.Cyanln("🙏 Thanks for installing Niom 🙏")

	terminal.CmdExecute(appName, "go", []string{"mod", "tidy"})
	dev()
}
func newApp(cmd []string) {
	if len(cmd) < 3 {
		println(`
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
	println("build app")
}
func info() {
	welcome()
}
func update() {
	println("update app")
}

func g() {
	println("Generate app")
}

func dev() {
	watcher.Watch()
	terminal.CmdExecute("myapp", "go", []string{"run", "main.go"})
}

func version() {
	println("Generate app")
}

func welcome() {

	println(`
-------------------- Welcome to world of  -------------------------
-------------------------------------------------------------------
-------- ____    ______  __  __  ____    ______  ------------------
------- /\  _ \ /\__  _\/\ \/\ \/\  _ \ /\  _  \ ----- Version: ---
------- \ \,\L\_\/_/\ \/\ \ \ \ \ \ \L\ \ \ \L\ \ ------ 0.1 ------
-------- \/_\__ \  \ \ \ \ \ \ \ \ \ ,__/\ \  __ \ ----------------
---------- /\ \L\ \ \ \ \ \ \ \_\ \ \ \/  \ \ \/\ \ ---------------
---------- \  \____\ \ \_\ \ \_____\ \_\   \ \_\ \_\ --------------
----------- \/_____/  \/_/  \/_____/\/_/    \/_/\/_/ --------------
-------------------------------------------------------------------
------------------------- Version: 0.1 ---------------------------- `)
	// 	println(
	// 		`
	// ---------------------------- Welcome to the world of ------------------------------
	// `)
	// 	myFigure := figure.NewFigure("...STUPA ...", "larry3d", true)
	// 	myFigure.Print()
	// 	println(`
	// -----------------------------------------------------------------------------------
	// --------------------------------- Version: 0.1 ------------------------------------`)
	println("\nTry -h, --help  for usage information.\n")

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
		update()
	case "g":
		fallthrough
	case "generate":
		g()
	case "dev":
		g()
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
