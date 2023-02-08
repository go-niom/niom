package config

const AppConfig = `package config

import (
	"strconv"
	"time"
)

// App holds the App configuration
type App struct {
	AppName     string
	AppEnv      string
	Host        string
	Port        int
	Debug       bool
	LogLevel    string
	ReadTimeout time.Duration
}

var app = &App{}

// AppCfg returns the default App configuration
func AppCfg() *App {
	return app
}

// LoadApp loads App configuration
func LoadAppCfg() {
	app.AppEnv = getEnv("APP_ENV", "development")
	app.AppName = getEnv("APP_NAME", "APP_NAME")
	app.Host = getEnv("APP_HOST", "localhost")
	app.LogLevel = getEnv("LOG_LEVEL", "0")
	app.Port, _ = strconv.Atoi(getEnv("APP_PORT", "3100"))
	app.Debug, _ = strconv.ParseBool(getEnv("APP_DEBUG", "false"))
	timeOut, _ := strconv.Atoi(getEnv("APP_READ_TIMEOUT", "1000"))
	app.ReadTimeout = time.Duration(timeOut) * time.Second
}
`
