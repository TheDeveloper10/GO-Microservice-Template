package api

import (
	"msvc-template/internal/api/middleware"

	"github.com/gofiber/fiber/v2"
)

func NewFiberApp() *fiber.App {
	return fiber.New(fiber.Config{
		ErrorHandler: middleware.GeneralErrorHandler,
	})
}
