package engine

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/niom/res/misc"
	"github.com/niom/res/pkg/common"
	"github.com/niom/res/pkg/config"
	"github.com/niom/res/pkg/helpers"
	"github.com/niom/res/pkg/logger"
	"github.com/niom/res/pkg/middleware"
	"github.com/niom/res/server"
	"github.com/niom/res/src"
	"github.com/niom/utils"
)

func CreateInitialFiles(moduleName string) {

	appName := utils.GetAppName(moduleName)
	// createApp(moduleName, "app")
	// return
	// init go.mod file
	modFile(moduleName)

	// create niom-cli.json config file
	createNiomCli(moduleName)

	// create .env file with host and db examples
	// create .dockerignore file with host and db examples
	// create Dockerfile file with host and db examples
	createFiles(appName)

	// create main.go
	createModuleFile(moduleName)

	//create config files
	createConfigFiles(appName)

	//create helpers
	appHelper(appName)

	//logger
	loggerFile(moduleName)

	//appCommon
	appCommonFile(moduleName)

	//appMiddleware
	appMiddleware(moduleName)

	// create README.md with basic commands to run and application
	// create server folder
	// create /server/server.go
	createServer(moduleName)

	//create app
	createApp(moduleName, "app")

}

func createNiomCli(moduleName string) {
	appName := utils.GetAppName(moduleName)
	config := `{
	"module_name":"` + moduleName + `",
	"app_name": "{{ .NameLowerCase}}",
	"sourceRoot": "src"
}
  `
	utils.RenderWriteToFile(config, appName, appName+"/niom-cli.json")
}

func modFile(appName string) {
	moduleName := appName
	version := strings.Split(runtime.Version(), "go")
	init := fmt.Sprintf("module %s \n\ngo %s", moduleName, version[1])

	split := strings.Split(moduleName, "/")
	appName = split[len(split)-1]

	utils.RenderWriteToFile(init, appName, appName+"/go.mod")
}

func createFiles(appName string) {
	utils.RenderWriteToFile(utils.DockerIgnore, appName, appName+"/.dockerignore")
	utils.RenderWriteToFile(utils.DockerFile, appName, appName+"/.Dockerfile")
	utils.RenderWriteToFile(utils.Env, appName, appName+"/.env")
	utils.RenderWriteToFile(utils.Env, appName, appName+"/env.example")
	utils.RenderWriteToFile(misc.MiscReadme, appName, appName+"/README.md")
}

func createModuleFile(moduleName string) {
	appName := utils.GetAppName(moduleName)
	utils.RenderWriteToFileModule(utils.MainGo, appName+"/main.go", "main", moduleName)
}

func createConfigFiles(appName string) {
	directory := appName + "/pkg/config"
	utils.RenderWriteToFile(config.AppConfig, appName, directory+"/app.go")
	utils.RenderWriteToFile(config.PkgConfig, appName, directory+"/config.go")
	utils.RenderWriteToFile(config.DBConfig, appName, directory+"/db.go")
	utils.RenderWriteToFile(config.HelperConfig, appName, directory+"/helper.go")
	utils.RenderWriteToFile(config.JWTConfig, appName, directory+"/jwt.go")
}

func loggerFile(moduleName string) {
	appName := utils.GetAppName(moduleName)
	directory := appName + "/pkg/logger"
	utils.RenderWriteToFileModule(logger.Logger, directory+"/logger.go", "logger", moduleName)
}

func appCommonFile(moduleName string) {
	appName := utils.GetAppName(moduleName)
	directory := appName + "/pkg/common"
	utils.RenderWriteToFileModule(common.CommonRouter, directory+"/common.go", "common", moduleName)
}

func appMiddleware(moduleName string) {
	appName := utils.GetAppName(moduleName)
	directory := appName + "/pkg/middleware"
	utils.RenderWriteToFileModule(middleware.MiddlewareFiber, directory+"/fiber.go", "middleware", moduleName)
}

func appHelper(appName string) {
	directory := appName + "/pkg/helpers"
	utils.RenderWriteToFile(helpers.HelperResponse, appName, directory+"/response.go")
}

func createServer(moduleName string) {
	appName := utils.GetAppName(moduleName)
	directory := appName + "/server"
	utils.RenderWriteToFileModule(server.Server, directory+"/server.go", "server", moduleName)
	utils.RenderWriteToFileModule(server.Routers, directory+"/router.go", "server", moduleName)
	utils.RenderWriteToFileModule(server.Connectors, directory+"/connecters.go", "server", moduleName)
	utils.RenderWriteToFileModule(server.Middleware, directory+"/middleware.go", "server", moduleName)
}

func createApp(moduleName, resName string) {
	appName := utils.GetAppName(moduleName)
	directoryRes := appName + "/src/app/" + resName
	directoryRes2 := appName + "/src/app/"
	utils.RenderWriteToFileModule(src.ControllerTmpl, directoryRes+".controller.go", resName, moduleName)
	utils.RenderWriteToFileModule(src.ServiceTmpl, directoryRes+".service.go", resName, moduleName)
	utils.RenderWriteToFileModule(src.RouterTmpl, directoryRes+".router.go", resName, moduleName)

	utils.RenderWriteToFileModule(src.DTO, directoryRes2+"dto/"+resName+".dto.go", resName, moduleName)
	utils.RenderWriteToFileModule(src.Model, directoryRes2+"model/"+resName+".model.go", resName, moduleName)
}
