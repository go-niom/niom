package utils

var PkgUtils = `package utils

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"{{ .ModuleName}}/pkg/logger"
)

func GetRoutes(app *fiber.App) {
	for _, r := range app.GetRoutes() {
		if r.Method == "HEAD" {
			continue
		}
		logger.Info(fmt.Sprintf("%s	%s", r.Method, r.Path))
	}
}`
