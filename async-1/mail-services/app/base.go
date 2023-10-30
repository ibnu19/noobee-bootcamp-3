package app

import (
	"async-1/mail-services/config"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func RegisterMailService(router fiber.Router, validate *validator.Validate, config config.EmailConfig) {

	repository := NewRepository(config)
	service := NewService(repository, validate)
	handler := NewHandler(service)

	router.Post("/send", handler.SendEmail)

}
