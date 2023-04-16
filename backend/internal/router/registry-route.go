package router

import (
	"backend/internal/controller"
	"backend/internal/middleware"
	"backend/internal/rule"

	"github.com/gofiber/fiber/v2"
)

func registryRouter(app *fiber.App, service rule.RegistryService) {
	app.Post("/registry/enroll", middleware.Protected(), controller.EnrollUser(service))
	app.Put("/registry/confirmEnrollship", middleware.Protected(), controller.ConfirmEnrollship(service))
	app.Post("/upload", middleware.Protected(), controller.UploadFile(service))
}
