package controller

import (
	"msvc-template/internal/service"
	"msvc-template/internal/util"

	"github.com/gofiber/fiber/v2"
)

type Example struct {
	svc service.IExample
}

func (ctrl *Example) Handler(c *fiber.Ctx) error {
	status := ctrl.svc.Func()

	if status == util.StatusSuccess {
		return c.SendStatus(fiber.StatusOK)
	} else {
		return c.SendStatus(fiber.StatusBadRequest)
	}
}
