package server

const Connectors = `package server

import (
	"runtime/debug"

	"{{ .ModuleName}}/pkg/logger"
)

func initConnectors() {
	log := logger.Get()
	buildInfo, _ := debug.ReadBuildInfo()
	log.Info().Str("go_version", buildInfo.GoVersion).Msg("Logger initialized")
}
`
