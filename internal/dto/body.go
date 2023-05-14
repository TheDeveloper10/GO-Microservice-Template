package dto

import "github.com/gofiber/fiber/v2"

type Body struct {
	Field string `json:"field"`
}

func (body *Body) Validate() error {
	if len(body.Field) <= 4 {
		return fiber.NewError(fiber.StatusBadRequest, "Body must have a length greater than 4")
	}

	return nil
}
