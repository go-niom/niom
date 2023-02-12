package misc

const MiscMainGo = `
package main

import (
	// load API Docs files (Swagger)

	"{{ .ModuleName}}/server"
	_ "{{ .ModuleName}}/docs"
	"{{ .ModuleName}}/pkg/config"
)

// @title Travel App
// @version 1.0
// @description Travel App Backend REST API
// @in header
// @name Authorization
// @host localhost:7000
// @BasePath /api
func main() {

	// setup various configuration for app
	config.LoadAllConfigs(".env")
	server.Serve()
}


`
