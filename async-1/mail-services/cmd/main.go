package main

import (
	"async-1/mail-services/app"
	"async-1/mail-services/config"
	"async-1/mail-services/utils"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	router := fiber.New(fiber.Config{
		AppName: "Mail Services",
	})
	router.Use(cors.New())
	router.Use(logger.New())

	validate := utils.NewValidator()

	config, err := config.LoadConfig()
	if err != nil {
		log.Println(err.Error())
	}

	app.RegisterMailService(router, validate, config.Email)
	log.Println("running email server...")

	if err := router.Listen(config.App.Port); err != nil {
		log.Println(err.Error())
	}
}
