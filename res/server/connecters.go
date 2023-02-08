package server

const Connectors = `package server

import "{{ .ModuleName}}/pkg/logger"

func initConnectors() {
	log := logger.Get()
	log.Info().Msg("Logger initialized")
	// connect to DB
	// if err := database.ConnectDB(); err != nil {
	// 	logr.Panicf("failed database setup. error: %v", err)
	// }
}
`
