package src

var RouterTmpl = `package {{ .NameLowerCase}}

import (
	"github.com/gofiber/fiber/v2"
)

func {{ .Name}}Router(router fiber.Router) {

	route := router.Group("/{{ .NameLowerCase}}")
	route.Get("/:id", {{ .Name}}Controller.GetAll{{ .Name}})
	route.Get("/", {{ .Name}}Controller.Get{{ .Name}}ById)
	route.Post("/", {{ .Name}}Controller.Create{{ .Name}})
	route.Patch("/", {{ .Name}}Controller.Update{{ .Name}})
	route.Delete("/:id", {{ .Name}}Controller.Delete{{ .Name}})
}

`
