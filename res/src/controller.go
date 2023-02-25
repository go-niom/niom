package src

var ControllerTmpl = `package {{ .NameLowerCase}}

import (
	"{{ .ModuleName}}/pkg/logger"
	"{{ .ModuleName}}/pkg/response"
	dto "{{ .ModuleName}}/src/{{ .NameLowerCase}}/dto"
//	model "{{ .ModuleName}}/src/{{ .NameLowerCase}}/model"
	"github.com/gofiber/fiber/v2"
)

var {{ .Name}}Controller {{ .NameLowerCase}}ControllerInterface = &{{ .NameLowerCase}}Controller{}

type {{ .NameLowerCase}}Controller struct {
}

type {{ .NameLowerCase}}ControllerInterface interface {
	Create{{ .Name}}(*fiber.Ctx) error
	GetAll{{ .Name}}(*fiber.Ctx) error
	Get{{ .Name}}ById(*fiber.Ctx) error
	Update{{ .Name}}(*fiber.Ctx) error
	Delete{{ .Name}}(*fiber.Ctx) error
}

// @title Create{{ .Name}}
// @Description Create{{ .Name}} func creates a {{ .NameLowerCase}}.
// @Summary Create a {{ .NameLowerCase}}
// @Tags {{ .Name}}
// @Accept json
// @Produce json
// @Param request body dto.Create{{ .Name}}Dto true "Create {{ .Name}} Request Body"
// @Success 201 {object} model.{{ .Name}} "Created"
// @Failure 400 {object} response.ErrMessage "Error"
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

		return response.Error(ctx, "Error Descriptions", err)
	}
	return response.Create(ctx, res)
}

// @title GetAll{{ .Name}}
// @Description GetAll{{ .Name}} func get all {{ .NameLowerCase}}.
// @Summary Get all {{ .NameLowerCase}}
// @Tags {{ .Name}}
// @Accept json
// @Produce json
// @Success 200 {array} model.{{ .Name}} "Ok"
// @Failure 400 {object} response.ErrMessage "Error"
// @Security ApiKey{{ .Name}}
// @Router /v1/{{ .NameLowerCase}} [Get]
func (d *{{ .NameLowerCase}}Controller) GetAll{{ .Name}}(ctx *fiber.Ctx) error {
	log := "Controller : {{ .Name}}, Method : RetrieveAll{{ .Name}}"
	logger.Info(log)
	query := new(dto.Query{{ .Name}}Dto)
	if err := ctx.QueryParser(query); err != nil {
		return err
	}
	res, err := {{ .Name}}Service.findAll(query)
	if err != nil {
		return response.Error(ctx, "Error Descriptions", err)
	}
	return response.GetList(ctx, res)
}

// @title Get{{ .Name}}ByID
// @Description GetAll{{ .Name}} func get {{ .NameLowerCase}} by ID.
// @Summary Get a {{ .NameLowerCase}}
// @Tags {{ .Name}}
// @Accept json
// @Produce json
// @Param id path string true "{{ .Name}} ID"
// @Success 200 {object} model.{{ .Name}} "Ok"
// @Failure 400 {object} response.ErrMessage "Error"
// @Security ApiKey{{ .Name}}
// @Router /v1/{{ .NameLowerCase}}/{id} [Get]
func (d *{{ .NameLowerCase}}Controller) Get{{ .Name}}ById(ctx *fiber.Ctx) error {
	log := "Controller : {{ .Name}}, Method : Retrieve{{ .Name}}"
	logger.Info(log)

	id := ctx.Params("id")
	res, err := {{ .Name}}Service.findOne(id)
	if err != nil {

		return response.Error(ctx, "Error Descriptions", err)
	}
	return response.Get(ctx, res)
}

// @title Update{{ .Name}}
// @Description Update{{ .Name}} func update a {{ .NameLowerCase}}.
// @Summary Update {{ .Name}} by Id
// @Tags {{ .Name}}
// @Accept json
// @Produce json
// @Param id path string true "{{ .Name}} ID"
// @Param request body dto.Update{{ .Name}}Dto true "Update {{ .Name}} Request Body"
// @Success 200 {object} model.{{ .Name}} "Ok"
// @Success 204 {object} response.ResMessage "Updated"
// @Failure 400 {object} response.ErrMessage "Error"
// @Security ApiKey{{ .Name}}
// @Router /v1/{{ .NameLowerCase}} [Patch]
func (d *{{ .NameLowerCase}}Controller) Update{{ .Name}}(ctx *fiber.Ctx) error {
	log := "Controller : {{ .Name}}, Method : Update{{ .Name}}"
	logger.Info(log)
	id := ctx.Params("id")
	body := new(dto.Update{{ .Name}}Dto)
	if err := ctx.BodyParser(body); err != nil {
		logger.Error(log, err)
		return response.Error(ctx, "Error Descriptions", err)
	}
	res, err := {{ .Name}}Service.update(id, body)
	if err != nil {

		return response.Error(ctx, "Error Descriptions", err)
	}
	return response.UpdatedData(ctx, res)
}

// Delete{{ .Name}} func delete a {{ .NameLowerCase}}.
// @Description delete {{ .NameLowerCase}}
// @Summary delete a {{ .NameLowerCase}}
// @Tags {{ .Name}}
// @Accept json
// @Produce json
// @Param id path string true "{{ .Name}} ID"
// @Success 204 {object} response.ResMessage "Ok"
// @Failure 400 {object} response.ErrMessage "Error"
// @Security ApiKey{{ .Name}}
// @Router /v1/{{ .NameLowerCase}}/{id} [delete]
func (d *{{ .NameLowerCase}}Controller) Delete{{ .Name}}(ctx *fiber.Ctx) error {
	log := "Controller : {{ .Name}}, Method : Delete{{ .Name}}"
	logger.Info(log)

	id := ctx.Params("id")
	res, err := {{ .Name}}Service.remove(id)
	if err != nil {
		return response.Error(ctx, "Error Descriptions", err)
	}
	return response.Delete(ctx, res)
}

`

var ControllerEmptyTmpl = `package {{ .NameLowerCase}}

import (
	"errors"

	"{{ .ModuleName}}/pkg/logger"
	"{{ .ModuleName}}/pkg/response"
	"github.com/gofiber/fiber/v2"
)

var {{ .Name}}Controller {{ .NameLowerCase}}ControllerInterface = &{{ .NameLowerCase}}Controller{}

type {{ .NameLowerCase}}Controller struct {
}

type {{ .NameLowerCase}}ControllerInterface interface {
	Create{{ .Name}}(*fiber.Ctx) error
	GetAll{{ .Name}}(*fiber.Ctx) error
	Get{{ .Name}}ById(*fiber.Ctx) error
	Update{{ .Name}}(*fiber.Ctx) error
	Delete{{ .Name}}(*fiber.Ctx) error
}

// @title Create{{ .Name}}
// @Description Create{{ .Name}} func creates a {{ .NameLowerCase}}.
// @Summary Create a {{ .NameLowerCase}}
// @Tags {{ .Name}}
// @Accept json
// @Produce json
// @Param request body map[string]interface{}  true "Create Name Request Body"
// @Success 201 {object} response.ResMessage "Created"
// @Failure 400 {object} response.ErrMessage "Error"
// @Security ApiKey{{ .Name}}
// @Router /v1/{{ .NameLowerCase}} [Post]
func (d *{{ .NameLowerCase}}Controller) Create{{ .Name}}(ctx *fiber.Ctx) error {
	log := "Controller : Name, Method : CreateName"
	logger.Info(log)

	var body *interface{} //Replace *interface{} with your dto create
	if err := ctx.BodyParser(body); err != nil {
		return err
	}

	//Add Your logic for response(res) and error
	res, err := "Create func. of controller", errors.New("Error String")
	if err != nil {
		return response.Error(ctx, "Error Descriptions", err)
	}

	return response.Create(ctx, res)
}

// @title GetAll{{ .Name}}
// @Description GetAll{{ .Name}} func get all {{ .NameLowerCase}}.
// @Summary Get all {{ .NameLowerCase}}
// @Tags {{ .Name}}
// @Accept json
// @Produce json
// @Success 200 {array} string "Ok"
// @Failure 400 {object} response.ErrMessage "Error"
// @Security ApiKey{{ .Name}}
// @Router /v1/{{ .NameLowerCase}} [Get]
func (d *{{ .NameLowerCase}}Controller) GetAll{{ .Name}}(ctx *fiber.Ctx) error {
	log := "Controller : Name, Method : RetrieveAllName"
	logger.Info(log)

	var query *interface{} //Replace *interface{} with your dto query
	if err := ctx.QueryParser(query); err != nil {
		return err
	}

	//Add Your logic for response(res) and error
	var res, err *interface{} = query, nil
	if err != nil {
		return response.Error(ctx, "Error Descriptions", err)
	}
	return response.GetList(ctx, res)
}

// @title Get{{ .Name}}ByID
// @Description GetAll{{ .Name}} func get {{ .NameLowerCase}} by ID.
// @Summary Get a {{ .NameLowerCase}}
// @Tags {{ .Name}}
// @Accept json
// @Produce json
// @Param id path string true "{{ .Name}} ID"
// @Success 200 {object} string "Ok"
// @Failure 400 {object} response.ErrMessage "Error"
// @Security ApiKey{{ .Name}}
// @Router /v1/{{ .NameLowerCase}}/{id} [Get]
func (d *{{ .NameLowerCase}}Controller) Get{{ .Name}}ById(ctx *fiber.Ctx) error {
	log := "Controller : Name, Method : RetrieveName"
	logger.Info(log)

	id := ctx.Params("id")

	//Add Your logic for response(res) and error
	res, err := id, errors.New("Error String")
	if err != nil {

		return response.Error(ctx, "Error Descriptions", err)
	}
	return response.Get(ctx, res)
}

// @title Update{{ .Name}}
// @Description Update{{ .Name}} func update a {{ .NameLowerCase}}.
// @Summary Update {{ .Name}} by Id
// @Tags {{ .Name}}
// @Accept json
// @Produce json
// @Param id path string true "{{ .Name}} ID"
// @Param request body map[string]interface{} true "Update Name Request Body"
// @Success 200 {object} string "Ok"
// @Success 204 {object} response.ResMessage "Updated"
// @Failure 400 {object} response.ErrMessage "Error"
// @Security ApiKey{{ .Name}}
// @Router /v1/{{ .NameLowerCase}} [Patch]
func (d *{{ .NameLowerCase}}Controller) Update{{ .Name}}(ctx *fiber.Ctx) error {
	log := "Controller : Name, Method : UpdateName"
	logger.Info(log)
	id := ctx.Params("id")
	var body *interface{} //Replace *interface{} with your dto update
	if err := ctx.BodyParser(body); err != nil {
		logger.Error(log, err)
		return response.Error(ctx, "Error Descriptions", err)
	}

	//Add Your logic for response(res) and error
	res, err := id, errors.New("Error String")
	if err != nil {
		return response.Error(ctx, "Error Descriptions", err)
	}
	return response.UpdatedData(ctx, res)
}

// Delete{{ .Name}} func delete a {{ .NameLowerCase}}.
// @Description delete {{ .NameLowerCase}}
// @Summary delete a {{ .NameLowerCase}}
// @Tags {{ .Name}}
// @Accept json
// @Produce json
// @Param id path string true "{{ .Name}} ID"
// @Success 204 {object} response.ResMessage "Ok"
// @Failure 400 {object} response.ErrMessage "Error"
// @Security ApiKey{{ .Name}}
// @Router /v1/{{ .NameLowerCase}}/{id} [delete]
func (d *{{ .NameLowerCase}}Controller) Delete{{ .Name}}(ctx *fiber.Ctx) error {
	log := "Controller : Name, Method : DeleteName"
	logger.Info(log)

	id := ctx.Params("id")

	//Add Your logic for response(res) and error
	res, err := id, errors.New("Error String")
	if err != nil {
		return response.Error(ctx, "Error Descriptions", err)
	}
	return response.Delete(ctx, res)
}

`
