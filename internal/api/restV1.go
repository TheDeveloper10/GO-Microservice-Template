package api

import (
	"msvc-template/internal/api/middleware"
	"msvc-template/internal/controller"
	"msvc-template/internal/repository"
	"msvc-template/internal/service"

	"github.com/gofiber/fiber/v2"
)

func SetUpRESTV1(app *fiber.App) {
	var (
		exampleRepo = repository.NewExampleRepository()
	)

	var (
		exampleSvc = service.NewExampleService(exampleRepo)
	)

	var (
		exampleCtrl = controller.NewExampleController(exampleSvc)
	)

	api := app.Group("/v1")

	api.Use(middleware.CORS)

	api.Get("/example", exampleCtrl.Handler)
}
