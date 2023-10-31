package app

import (
	"async-1/app-services/utils"
	"log"

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

func (handler *handler) SendEmail(c *fiber.Ctx) error {
	email := EmailRequest{}
	err := c.BodyParser(&email)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(utils.ApiResponse(fiber.ErrBadRequest.Message, err))
	}

	response := handler.service.SendEmail(c, email)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
