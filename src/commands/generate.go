package commands

import (
	"fmt"

	"github.com/go-niom/niom/pkg/logger"
	"github.com/go-niom/niom/pkg/utils"
	"github.com/go-niom/niom/res/src"
	"github.com/go-niom/niom/src/engine"
)

func generate(cmd string, args []string) {
	niomConfig := utils.GetNiomCliConfig()
	if niomConfig == nil {
		return
	}
	moduleName := niomConfig.ModuleName

	if len(args) < 3 {
		logger.Error("Schematic type is missing", "")
		logger.Info(fmt.Sprintf("Valid Syntax is %s <schematic> [name]", cmd))
		return
	}

	resType := args[2]
	if len(args) < 4 {
		logger.Error("Schematic Name is missing", "")
		logger.Info(fmt.Sprintf("Valid Syntax is %s %s [name]", resType, cmd))
		return
	}

	filePath := "src/"
	resName := args[3]
	if len(args) == 5 {
		logger.Info(fmt.Sprintf("Valid Syntax is %s %s [name]", resType, cmd))
		filePath = fmt.Sprintf("%s%s/", filePath, args[4])
	} else {
		filePath = fmt.Sprintf("%s%s/", filePath, resName)
	}

	logger.Info(fmt.Sprintf("Generating schematic of type: `%s`, name: %s at `%s`", resType, resName, filePath))
	directoryFile := filePath + resName

	switch resType {
	case "resource", "res":
		engine.CreateResource(moduleName, filePath, resName)
	case "controller", "co":
		utils.RenderWriteToFileModule(src.ControllerTmpl, directoryFile+".controller.go", resName, moduleName)
	case "service", "s":
		utils.RenderWriteToFileModule(src.ServiceTmpl, directoryFile+".service.go", resName, moduleName)
	case "router", "ro":
		utils.RenderWriteToFileModule(src.RouterTmpl, directoryFile+".router.go", resName, moduleName)
	case "model", "mo":
		utils.RenderWriteToFileModule(src.Model, filePath+"model/"+resName+".model.go", resName, moduleName)
	case "dto", "d":
		utils.RenderWriteToFileModule(src.DTO, filePath+"dto/"+resName+".dto.go", resName, moduleName)
	case "interface", "in":
	case "middleware", "mi":

	default:
		fmt.Printf("Invalid Command %s\n", resType)
	}
}
