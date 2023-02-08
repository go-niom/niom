package server

const Middleware = `package server

import (
	"github.com/gofiber/fiber/v2"
	"{{ .ModuleName}}/pkg/middleware"
)

//register global Middleware
func registerMiddleware(fApp *fiber.App) {
	middleware.FiberMiddleware(fApp)
}
`
