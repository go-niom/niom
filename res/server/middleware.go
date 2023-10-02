package server

const Middleware = `package server

import (
	"github.com/gofiber/fiber/v2"
	"{{ .ModuleName}}/pkg/database"
	"{{ .ModuleName}}/pkg/logger"
	"{{ .ModuleName}}/pkg/middleware"
)

//register global Middleware
func registerMiddleware(fApp *fiber.App) {

	//initialize postgres
	pgsql := database.GetDB()
	if pgsql == nil {
		logger.Error("failed to initialize postgres", nil)
		return
	}

	//initialize redish
	// redis := database.ConnetRedis()
	// if redis == nil {
	// 	logger.Error("failed to initialize redis", nil)
	// 	return
	// }

	fApp.Use(func(c *fiber.Ctx) error {
		c.Locals("rdb", redis)
		c.Locals("pgsql", pgsql)
		return c.Next()
	})

	middleware.FiberMiddleware(fApp)
}
`
