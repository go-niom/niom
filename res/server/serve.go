package server

const Server = `package server

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"{{ .ModuleName}}/pkg/config"
	"{{ .ModuleName}}/pkg/logger"
)


// Serve ..
func Serve() {
	appCfg := config.AppCfg()

	// Define Fiber config & app.
	fiberCfg := config.FiberConfig()
	app := fiber.New(fiberCfg)

	//initial the connecters
	initConnectors()

	//Attach Middleware
	registerMiddleware(app)
	registerRouters(app)

	//Print app routes
	//utils.GetRoutes(app)

	// signal channel to capture system calls
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	// start shutdown goroutine
	go func() {
		// capture sigterm and other system call here
		<-sigCh
		fmt.Println("Shutting down server...")
		_ = app.Shutdown()
	}()

	//Print server info
	serverAddPrint := fmt.Sprintf("App Serving on http://%s:%d", config.AppCfg().Host, config.AppCfg().Port)
	logger.Info(serverAddPrint)
	
	// start http server
	serverAddr := fmt.Sprintf("%s:%d", appCfg.Host, appCfg.Port)
	if err := app.Listen(serverAddr); err != nil {
		fmt.Println("Oops... server is not running! error: %v", err)
	}

}
`
