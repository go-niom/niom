package common

const CommonRouter = `package common

import (
	"github.com/gofiber/swagger"
	"github.com/gofiber/fiber/v2"
)

func GeneralRoute(a fiber.Router) {
	a.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"msg":    "Hello World!",
			"docs":   "/swagger/index.html",
			"status": "/h34l7h",
		})
	})

	a.Get("/h34l7h", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"msg":       "Health Check",
		})
	})
}

func SwaggerRoute(a fiber.Router) {
	//Create route group.
	a.Get("/swagger/*", swagger.HandlerDefault)
}

func NotFoundRoute(a *fiber.App) {
	a.Use(
		func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"msg": "sorry, endpoint is not found",
			})
		},
	)
}
`
