package dto

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ID struct {
	ID uint64
}

func (id *ID) Fill(c *fiber.Ctx) error {
	str := c.Params("id")
	value, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "ID must be a positive integer")
	} else if value <= 0 {
		return fiber.NewError(fiber.StatusBadRequest, "ID must be greater than 0")
	}

	id.ID = value
	return nil
}
