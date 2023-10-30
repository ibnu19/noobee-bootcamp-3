package main

import (
	"async-1/app-services/app"
	"async-1/app-services/utils"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

const BASE_URL = "http://localhost:8080/send"

func main() {
	router := fiber.New(fiber.Config{
		AppName: "App Services",
	})
	router.Use(cors.New())
	router.Use(logger.New())
	repo := app.NewRepository()

	router.Post("/send", func(c *fiber.Ctx) error {
		email := app.Email{}
		err := c.BodyParser(&email)
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusBadRequest).JSON(utils.ApiResponse(fiber.ErrBadRequest.Message, err))
		}

		email.Type = "text/html"
		response := repo.SendEmail(c, BASE_URL, email)
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusBadRequest).JSON(response)
		}

		return c.Status(fiber.StatusOK).JSON(response)
	})

	if err := router.Listen(":3000"); err != nil {
		log.Println(err)
	}
}
