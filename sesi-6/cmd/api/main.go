package main

import (
	"log"
	"sesi-6/config"
	"sesi-6/pkg/database"

	"github.com/gofiber/fiber/v2"
)

func main() {

	err := config.LoadConfig("./config/config.yaml")
	if err != nil {
		panic(err)
	}

	db, err := database.ConnectPostgres(config.Cfg.DB)
	if err != nil {
		panic(err)
	}

	if db != nil {
		log.Println("conncet database")
	}

	router := fiber.New(fiber.Config{
		Prefork: true,
		AppName: "Product Service",
	})

	router.Listen(config.Cfg.App.Port)
}
