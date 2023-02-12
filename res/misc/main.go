package misc

const MiscMainGo = `
package main

import (
	// load API Docs files (Swagger)

	"{{ .ModuleName}}/server"
	_ "{{ .ModuleName}}/docs"
	"{{ .ModuleName}}/pkg/config"
)

// @title {{ .Name}}
// @version 1.0
// @description {{ .Name}} Backend REST API
// @in header
// @name Authorization
// @host 127.0.0.1:7000
// @BasePath /api
func main() {

	// setup various configuration for app
	config.LoadAllConfigs(".env")
	server.Serve()
}


`
