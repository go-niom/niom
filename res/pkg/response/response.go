package response

const HelperResponse = `package response

import "github.com/gofiber/fiber/v2"

type ErrMessage struct {
	Message string      ` + "`" + `json:"message,omitempty""` + "`" + `
	
	Error   interface{} ` + "`" + `json:"error,omitempty""` + "`" + `
}

type ResMessage struct {
	Message string       ` + "`" + `json:"message,omitempty""` + "`" + `
	Data    interface{}  ` + "`" + `data:"message,omitempty""` + "`" + `
}


// Rest API Success Responses

// Success Response
// Method Get
// Get single item
// HTTP Response Code: 200
// API  /:id
func Get(c *fiber.Ctx, data interface{}) error {
	c.Set("Content-Type", " application/json")
	c.Status(fiber.StatusOK)
	return c.JSON(data)
}

// Success Response
// Method Get
// Get Item List
// HTTP Response Code: 200
// API /
func GetList(c *fiber.Ctx, data interface{}) error {
	c.Set("Pagination-Count", "100")
	c.Set("Pagination-Page", "5")
	c.Set("Pagination-Limit", "20")
	c.Set("Content-Type", "application/json")
	c.Status(fiber.StatusOK)
	return c.JSON(data)
}

// Success Response
// Method POST
// Create a new item
// HTTP Response Code: 201
// API /
func Create(c *fiber.Ctx, message string) error {
	c.Set("Content-Type", " application/json")
	c.Status(fiber.StatusCreated)
	data := &ResMessage{
		Message: message,
	}
	return c.JSON(data)
}

// Success Response
// Method PUT/PATCH
// Returned updated object
// HTTP Response Code: 200
// API /
func UpdatedData(c *fiber.Ctx, data interface{}) error {
	c.Set("Content-Type", " application/json")
	c.Status(fiber.StatusOK)
	return c.JSON(data)
}

// Success Response
// Method PUT/PATCH
// Returned Updated Message
// HTTP Response Code: 204
// API /
func Update(c *fiber.Ctx, message string) error {
	c.Set("Content-Type", " application/json")
	c.Status(fiber.StatusNoContent)
	data := &ResMessage{
		Message: message,
	}
	return c.JSON(data)
}

// Success Response
// Method DELETE
// Returned Appropriate Message
// HTTP Response Code: 204
// API /
func Delete(c *fiber.Ctx, message string) error {
	c.Set("Content-Type", " application/json")
	c.Status(fiber.StatusNoContent)
	data := &ResMessage{
		Message: message,
	}
	return c.JSON(data)
}

//Rest API Error Responses

// Error Response
// Method GET/DELETE
// GET/DELETE single item
// HTTP Response Code: 404
// API  /:id
func NotFound(c *fiber.Ctx) error {
	c.Set("Content-Type", " application/json")
	c.Status(fiber.StatusNotFound)
	data := &ResMessage{
		Message: "The item does not exist",
	}
	return c.JSON(data)
}

// Error Response
// Method Get
// Get Item List
// HTTP Response Code: 200
// API /
func NotFoundList(c *fiber.Ctx) error {
	c.Set("Content-Type", " application/json")
	c.Status(fiber.StatusNotFound)
	data := &ResMessage{
		Message: "The item does not exist",
	}
	return c.JSON(data)
}

// Error Response
// HTTP Response Code: 400
// API /
func Error(c *fiber.Ctx, message string, err interface{}) error {
	c.Set("Content-Type", " application/json")
	c.Status(fiber.StatusBadRequest)
	data := &ErrMessage{
		Message: message,
		Error:   err,
	}
	return c.JSON(data)
}

// Unauthorized Response
// HTTP Response Code: 401
// API /
func Unauthorized(c *fiber.Ctx, message string) error {
	c.Set("Content-Type", " application/json")
	c.Status(fiber.StatusUnauthorized)
	data := &ResMessage{
		Message: message,
	}
	return c.JSON(data)
}

// Forbidden Response
// HTTP Response Code: 401
// API /
func Forbidden(c *fiber.Ctx, message string) error {
	c.Set("Content-Type", " application/json")
	c.Status(fiber.StatusForbidden)
	data := &ResMessage{
		Message: message,
	}
	return c.JSON(data)
}

// Conflict Response
// HTTP Response Code: 401
// API /
func Conflict(c *fiber.Ctx, message string) error {
	c.Set("Content-Type", " application/json")
	c.Status(fiber.StatusConflict)
	data := &ResMessage{
		Message: message,
	}
	return c.JSON(data)
}

// Too Many Requests Response
// HTTP Response Code: 401
// API /
func TooManyRequests(c *fiber.Ctx) error {
	c.Set("Content-Type", " application/json")
	c.Status(fiber.StatusTooManyRequests)
	data := &ResMessage{
		Message: "The request cannot be served due to the rate limit having been exhausted for the resource",
	}
	return c.JSON(data)
}

// Internal Server Error Response
// HTTP Response Code: 500
// API /
func InternalServerError(c *fiber.Ctx) error {
	c.Set("Content-Type", " application/json")
	c.Status(fiber.StatusInternalServerError)
	data := &ResMessage{
		Message: "Something is broken",
	}
	return c.JSON(data)
}

// Service Unavailable Error Response
// HTTP Response Code: 503
// API /
func ServiceUnavailableError(c *fiber.Ctx) error {
	c.Set("Content-Type", " application/json")
	c.Status(fiber.StatusServiceUnavailable)
	data := &ResMessage{
		Message: "Something is broken",
	}
	return c.JSON(data)
}

`
