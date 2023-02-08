package helpers

const HelperResponse = `package helpers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type Pagination struct {
	Data interface{} ` + "`" + `json:"data,omitempty"` + "`" + `
	Next bool        ` + "`" + `json:"next,omitempty"` + "`" + `
	Page int64       ` + "`" + `json:"page,omitempty"` + "`" + `
}

type ResMessage struct {
	Code    int         ` + "`" + `json:"code,omitempty"` + "`" + `
	Message string      ` + "`" + `json:"message,omitempty"` + "`" + `
	Result  interface{} ` + "`" + `json:"result,omitempty"` + "`" + `
	Error   interface{} ` + "`" + `json:"error,omitempty"` + "`" + `
}

type SuccessResponse struct {
	ResMessage
	Code    int  ` + "`" + `json:"code,omitempty" example:"200"` + "`" + `
	Success bool ` + "`" + `json:"success,omitempty"` + "`" + `
}

type ErrorResponses struct {
	ResMessage
	Success bool ` + "`" + `json:"success,omitempty"` + "`" + `
}

type WarningResponses struct {
	ResMessage
	Success bool ` + "`" + `json:"success,omitempty"` + "`" + `
}

func MsgResponse(c *fiber.Ctx, msg string, data interface{}) error {
	resPonse := &ResMessage{
		Code:    fiber.StatusOK,
		Message: msg,
		Result:  data,
	}
	return c.Status(fiber.StatusOK).JSON(resPonse)
}

func CrudResponse(c *fiber.Ctx, msg string, data interface{}) error {
	resPonse := &ResMessage{
		Code:    fiber.StatusOK,
		Message: fmt.Sprintf(" %s data succesfully", msg),
		Result:  data,
	}
	return c.Status(fiber.StatusOK).JSON(resPonse)
}

func ErrorResponse(c *fiber.Ctx, msg string, data interface{}) error {
	resPonse := &ResMessage{
		Code:    fiber.StatusBadRequest,
		Message: msg,
		Result:  data,
	}
	return c.Status(fiber.StatusBadRequest).JSON(resPonse)
}

func BadResponse(c *fiber.Ctx, msg string, data interface{}) error {
	resPonse := &ResMessage{
		Code:    fiber.StatusBadRequest,
		Message: msg,
		Result:  data,
	}
	return c.Status(fiber.StatusBadRequest).JSON(resPonse)
}

func ServerResponse(c *fiber.Ctx, msg string, data interface{}) error {
	resPonse := &ResMessage{
		Code:    fiber.StatusInternalServerError,
		Message: msg,
		Result:  data,
	}
	return c.Status(fiber.StatusInternalServerError).JSON(resPonse)
}

func NotFoundResponse(c *fiber.Ctx, data interface{}) error {
	resPonse := &ResMessage{
		Code:    fiber.StatusNotFound,
		Message: "Not Found",
		Result:  data,
	}
	return c.Status(fiber.StatusNotFound).JSON(resPonse)
}

type ValidateErrorResponse struct {
	FailedField string ` + "`" + `json:"code,omitempty"` + "`" + `
	Tag         string ` + "`" + `json:"message,omitempty"` + "`" + `
	Value       string ` + "`" + `json:"result,omitempty"` + "`" + `
}
`
