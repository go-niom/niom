package engine

import (
	"errors"
	"fmt"
	"os"
	"runtime"
	"strings"

	appLogger "github.com/go-niom/niom/pkg/logger"
	"github.com/go-niom/niom/pkg/utils"
	"github.com/go-niom/niom/res/misc"
	"github.com/go-niom/niom/res/pkg/common"
	"github.com/go-niom/niom/res/pkg/config"
	"github.com/go-niom/niom/res/pkg/logger"
	"github.com/go-niom/niom/res/pkg/middleware"
	"github.com/go-niom/niom/res/pkg/response"
	pkgUtils "github.com/go-niom/niom/res/pkg/utils"
	"github.com/go-niom/niom/res/server"
	"github.com/go-niom/niom/res/src"
)

func CreateInitialFiles(moduleName string) {

	appName := utils.GetAppName(moduleName)
	// create(moduleName, "app")
	// return
	// init go.mod file
	modFile(moduleName)

	// create niom-cli.json config file
	CreateNiomCli(moduleName)

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

	//utils
	appUtils(moduleName)

	//appCommon
	appCommonFile(moduleName)

	//appMiddleware
	appMiddleware(moduleName)

	// create README.md with basic commands to run and application
	// create server folder
	// create /server/server.go
	createServer(moduleName)

	//create app
	create(moduleName, "app")

}

func CreateNiomCli(moduleName string) {
	if moduleName != "" {
		appName := utils.GetAppName(moduleName)
		config := `{
	"module_name":"` + moduleName + `",
	"app_name": "{{ .NameLowerCase}}",
	"sourceRoot": "src"
	"configFile": ".env"
}
  `
		utils.RenderWriteToFile(config, appName, appName+"/niom-cli.json")
	} else {
		if _, err := os.Stat("niom-cli.json"); errors.Is(err, os.ErrNotExist) {
			config := `{
	"module_name":"",
	"app_name": "",
	"sourceRoot": "src",
	"configFile": ".env"
}
 `
			utils.RenderWriteToFile(config, "", "niom-cli.json")
		} else {
			appLogger.Warn("Already Exist")
		}

	}

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
	utils.RenderWriteToFile(misc.MiscDockerIgnore, appName, appName+"/.dockerignore")
	utils.RenderWriteToFile(misc.MiscDockerFile, appName, appName+"/.Dockerfile")
	utils.RenderWriteToFile(misc.MiscEnv, appName, appName+"/.env")
	utils.RenderWriteToFile(misc.MiscEnv, appName, appName+"/env.example")
	utils.RenderWriteToFile(misc.MiscReadme, appName, appName+"/README.md")
}

func createModuleFile(moduleName string) {
	appName := utils.GetAppName(moduleName)
	utils.RenderMain(misc.MiscMainGo, appName+"/main.go", "main", moduleName)
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
	utils.RenderWriteToFile(response.HelperResponse, appName, directory+"/response.go")
}

func appUtils(moduleName string) {
	appName := utils.GetAppName(moduleName)
	directory := appName + "/pkg/utils"
	utils.RenderWriteToFileModule(pkgUtils.PkgUtils, directory+"/utils.go", "utils", moduleName)
}

func createServer(moduleName string) {
	appName := utils.GetAppName(moduleName)
	directory := appName + "/server"
	utils.RenderWriteToFileModule(server.Server, directory+"/server.go", "server", moduleName)
	utils.RenderWriteToFileModule(server.Routers, directory+"/router.go", "server", moduleName)
	utils.RenderWriteToFileModule(server.Connectors, directory+"/connecters.go", "server", moduleName)
	utils.RenderWriteToFileModule(server.Middleware, directory+"/middleware.go", "server", moduleName)
}

func create(moduleName, resName string) {
	appName := utils.GetAppName(moduleName)
	dir := appName + "/src/app/"
	CreateResource(moduleName, dir, resName)
}

func CreateResource(moduleName, dir, resName string) {
	directoryFile := dir + resName
	utils.RenderWriteToFileModule(src.ControllerTmpl, directoryFile+".controller.go", resName, moduleName)
	utils.RenderWriteToFileModule(src.ServiceTmpl, directoryFile+".service.go", resName, moduleName)
	utils.RenderWriteToFileModule(src.RouterTmpl, directoryFile+".router.go", resName, moduleName)

	utils.RenderWriteToFileModule(src.DTO, dir+"dto/"+resName+".dto.go", resName, moduleName)
	utils.RenderWriteToFileModule(src.Model, dir+"model/"+resName+".model.go", resName, moduleName)
}
