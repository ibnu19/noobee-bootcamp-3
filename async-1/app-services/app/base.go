package app

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterEmailService(router fiber.Router, URL string) {
	repository := NewRepository()
	service := NewService(repository, URL)
	handler := NewHandler(service)

	emailRouter := router.Group("/send")
	{
		emailRouter.Post("", handler.SendEmail)
	}
}
