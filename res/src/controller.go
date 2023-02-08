package src

var ControllerTmpl = `package {{ .NameLowerCase}}

import (
	"github.com/gofiber/fiber/v2"
	"{{ .ModuleName}}/pkg/logger"
	helpers "{{ .ModuleName}}/pkg/helpers"
	dto "{{ .ModuleName}}/src/{{ .NameLowerCase}}/dto"
)

var {{ .Name}}Controller {{ .NameLowerCase}}ControllerInterface = &{{ .NameLowerCase}}Controller{}

type {{ .NameLowerCase}}Controller struct {
}

type {{ .NameLowerCase}}ControllerInterface interface {
	Create{{ .Name}}(*fiber.Ctx) error
	RetrieveAll{{ .Name}}(*fiber.Ctx) error
	Retrieve{{ .Name}}(*fiber.Ctx) error
	Update{{ .Name}}(*fiber.Ctx) error
	Delete{{ .Name}}(*fiber.Ctx) error
}

// Create{{ .Name}} func create a {{ .NameLowerCase}}.
// @Description create {{ .NameLowerCase}}
// @Summary create a {{ .NameLowerCase}}
// @Tags {{ .Name}}
// @Accept json
// @Produce json
// @Param request body Create{{ .Name}}Dto true "Update {{ .Name}} Request Body"
// @Success 200 {object} helpers.ResMessage(data={{ .Name}}) "Ok"
// Failure 400,404,401,500 {object} helpers.ResMessage "Error"
// @Security ApiKey{{ .Name}}
// @Router /v1/{{ .NameLowerCase}} [Post]
func (d *{{ .NameLowerCase}}Controller) Create{{ .Name}}(ctx *fiber.Ctx) error {
	log := "Controller : {{ .Name}}, Method : Create{{ .Name}}"
	logger.Info(log)

	body := new(dto.Create{{ .Name}}Dto)
	if err := ctx.BodyParser(body); err != nil {
		return err
	}
	res, err := {{ .Name}}Service.create(body)
	if err != nil {
		return helpers.ErrorResponse(ctx, "Error Descriptions", err)
	}
	return helpers.MsgResponse(ctx, "{{ .Name}} Created", res)
}

// GetAll{{ .Name}} func get all {{ .NameLowerCase}}.
// @Description get all {{ .NameLowerCase}}
// @Summary Get all {{ .NameLowerCase}}
// @Tags {{ .Name}}
// @Accept json
// @Produce json
// @Success 200 {object} helpers.ResMessage(data=[]{{ .Name}}) "Ok"
// {object} helpers.SuccessResponse
// Failure 400,404,401,500 {object} helpers.ResMessage "Error"
// @Security ApiKey{{ .Name}}
// @Router /v1/{{ .NameLowerCase}} [Get]
func (d *{{ .NameLowerCase}}Controller) RetrieveAll{{ .Name}}(ctx *fiber.Ctx) error {
	log := "Controller : {{ .Name}}, Method : RetrieveAll{{ .Name}}"
	logger.Info(log)

	query := new(dto.Query{{ .Name}}Dto)
	if err := ctx.QueryParser(query); err != nil {
		return err
	}
	res, err := {{ .Name}}Service.findAll(query)
	if err != nil {
		return helpers.ErrorResponse(ctx, "Error Descriptions", err)
	}
	return helpers.MsgResponse(ctx, "All {{ .Name}} Retrieved", res)
}

// Get{{ .Name}} func get a {{ .NameLowerCase}}.
// @Description get {{ .NameLowerCase}}
// @Summary get a {{ .NameLowerCase}}
// @Tags {{ .Name}}
// @Accept json
// @Produce json
// @Param id path string true "{{ .Name}} ID"
// @Success 200 {object} helpers.ResMessage(data={{ .Name}}) "Ok"
// {object} helpers.SuccessResponse
// Failure 400,404,401,500 {object} helpers.ResMessage "Error"
// @Security ApiKey{{ .Name}}
// @Router /v1/{{ .NameLowerCase}}/{id} [Get]
func (d *{{ .NameLowerCase}}Controller) Retrieve{{ .Name}}(ctx *fiber.Ctx) error {
	log := "Controller : {{ .Name}}, Method : Retrieve{{ .Name}}"
	logger.Info(log)

	id := ctx.Params("id")
	res, err := {{ .Name}}Service.findOne(id)
	if err != nil {

		return helpers.ErrorResponse(ctx, "Error Descriptions", err)
	}
	return helpers.MsgResponse(ctx, "{{ .Name}} Retrieved", res)
}

// GetAll{{ .Name}} func update a {{ .NameLowerCase}}.
// @Description update {{ .NameLowerCase}}
// @Summary Update {{ .Name}} by Id
// @Tags {{ .Name}}
// @Accept json
// @Produce json
// @Param id path string true "{{ .Name}} ID"
// @Param request body Update{{ .Name}}Dto true "Create {{ .Name}} Request Body"
// @Success 200 {object} helpers.ResMessage(data={{ .Name}}) "Ok"
// {object} helpers.SuccessResponse
// Failure 400,404,401,500 {object} helpers.ResMessage "Error"
// @Security ApiKey{{ .Name}}
// @Router /v1/{{ .NameLowerCase}} [Patch]
func (d *{{ .NameLowerCase}}Controller) Update{{ .Name}}(ctx *fiber.Ctx) error {
	log := "Controller : {{ .Name}}, Method : Update{{ .Name}}"
	logger.Info(log)
	id := ctx.Params("id")
	body := new(dto.Update{{ .Name}}Dto)
	if err := ctx.BodyParser(body); err != nil {
		logger.Error(log, err)
		return helpers.ErrorResponse(ctx, "Error Descriptions", err)
	}
	res, err := {{ .Name}}Service.update(id, body)
	if err != nil {

		return helpers.ErrorResponse(ctx, "Error Descriptions", err)
	}
	return helpers.MsgResponse(ctx, "{{ .Name}} Updated", res)
}

// Delete{{ .Name}} func delete a {{ .NameLowerCase}}.
// @Description delete {{ .NameLowerCase}}
// @Summary delete a {{ .NameLowerCase}}
// @Tags {{ .Name}}
// @Accept json
// @Produce json
// @Param id path string true "{{ .Name}} ID"
// Success 200 "Ok"
// {object} helpers.SuccessResponse
// Failure 400,404,401,500 {object} helpers.ResMessage "Error"
// @Security ApiKey{{ .Name}}
// @Router /v1/{{ .NameLowerCase}}/{id} [delete]
func (d *{{ .NameLowerCase}}Controller) Delete{{ .Name}}(ctx *fiber.Ctx) error {
	log := "Controller : {{ .Name}}, Method : Delete{{ .Name}}"
	logger.Info(log)

	id := ctx.Params("id")
	res, err := {{ .Name}}Service.remove(id)
	if err != nil {
		return helpers.ErrorResponse(ctx, "Error Descriptions", err)
	}
	return helpers.MsgResponse(ctx, "{{ .Name}} Deleted", res)
}

`
