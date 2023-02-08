package server

const Routers = `package server

import (
	"github.com/gofiber/fiber/v2"

	"{{ .ModuleName}}/pkg/common"
	"{{ .ModuleName}}/src/app"
)

func registerRouters(a fiber.Router) {
	common.GeneralRoute(a)
	app.AppRouter(a)
}
`
