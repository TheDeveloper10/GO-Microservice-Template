package controller

import (
	"msvc-template/internal/dto"
	"msvc-template/internal/service"
	"msvc-template/internal/util"

	"github.com/gofiber/fiber/v2"
)

type Example struct {
	svc service.IExample
}

func (ctrl *Example) Handler(c *fiber.Ctx) error {
	id := dto.ID{}
	if err := id.Fill(c); err != nil {
		return err
	}

	body := dto.Body{}
	if err := c.BodyParser(&body); err != nil {
		return err
	} else if err := body.Validate(); err != nil {
		return err
	}

	status := ctrl.svc.Func()

	if status == util.StatusSuccess {
		return c.SendStatus(fiber.StatusOK)
	} else {
		return c.SendStatus(fiber.StatusBadRequest)
	}
}
