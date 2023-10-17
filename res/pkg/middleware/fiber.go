package middleware

const MiddlewareFiber = `package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/gofiber/fiber/v2/middleware/logger"
	"{{ .ModuleName}}/pkg/common"
	"{{ .ModuleName}}/pkg/logger"
	"{{ .ModuleName}}/pkg/response"
)

func FiberMiddleware(a *fiber.App) {
	a.Use(
		// Add CORS to each route.
		cors.New(),
		// Add simple logger.
		fiberLogger.New(),
	)
}

func AuthMiddleware(ctx *fiber.Ctx) error {
	log := "method: AuthMiddleware"
	logger.Info(log)
	err := common.ValidateJwtToken(ctx)
	if err != nil {
		msg := "Unauthorized"
		logger.Error(fmt.Sprintf("%s %s", log, msg), err)
		return response.Unauthorized(ctx, msg)
	}
	return ctx.Next()
}

`
