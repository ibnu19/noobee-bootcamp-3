package main

import (
	"async-1/app-services/app"
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

	app.RegisterEmailService(router, BASE_URL)

	if err := router.Listen(":3000"); err != nil {
		log.Println(err)
	}
}
