package controller

import (
	"backend/internal/gateway"
	"backend/internal/presenter"
	"backend/internal/rule"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func EnrollUser(service rule.RegistryService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var params gateway.EnrollUserRegistry
		err := ctx.BodyParser(&params)
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(presenter.BadResponse)
		}

		res, err := service.EnrollUser(params.UserId, params.CourseId)
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(presenter.BadResponse)
		}

		ctx.Status(http.StatusOK)
		return ctx.JSON(presenter.SuccessRequest(res))
	}
}

func ConfirmEnrollship(service rule.RegistryService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var params gateway.EnrollUserRegistry
		err := ctx.BodyParser(&params)
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(presenter.BadResponse)
		}

		res, err := service.ConfirmEnrollship(params.UserId, params.CourseId)
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(presenter.BadResponse)
		}

		ctx.Status(http.StatusOK)
		return ctx.JSON(presenter.SuccessRequest(res))
	}
}

func UploadFile(service rule.RegistryService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		file, err := ctx.FormFile("upload")
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(presenter.BadResponse)
		}

		if err = ctx.SaveFile(file, "../../uploads/"+file.Filename); err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(presenter.BadResponse)
		}

		ctx.Status(http.StatusOK)
		return ctx.JSON(presenter.SuccessRequest("message: file written successfuly"))
	}
}
