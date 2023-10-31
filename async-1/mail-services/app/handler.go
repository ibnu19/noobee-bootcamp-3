package app

import (
	"async-1/mail-services/utils"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	service
}

func NewHandler(service service) handler {
	return handler{
		service: service,
	}
}

func (handler *handler) SendEmail(ctx *fiber.Ctx) error {
	email := EmailRequest{}
	err := ctx.BodyParser(&email)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(utils.ErrResponse(err, fiber.ErrBadRequest.Message))
	}

	err = handler.service.SendEmail(email)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(utils.ErrResponse(err, fiber.ErrBadRequest.Message))
	}

	return ctx.Status(fiber.StatusOK).
		JSON(utils.ApiResponse("the email has been sent successfully"))
}
