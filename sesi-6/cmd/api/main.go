package main

import (
	"log"
	"sesi-6/app/product"
	"sesi-6/config"
	"sesi-6/pkg/database"
	"sesi-6/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

func main() {
	router := fiber.New(fiber.Config{
		AppName: "Product Services",
	})

	validate := utils.NewValidator()

	err := config.LoadConfig("./config/config.yaml")
	if err != nil {
		log.Println("error to try loadconfig with err :", err.Error())
	}

	dbGORM, err := database.ConnectGORM(config.Cfg.DB)
	if err != nil {
		log.Println("cannot connect database")
	}

	dbSQLX, err := database.ConnectSQLX(config.Cfg.DB)
	if err != nil {
		log.Println("cannot connect database")
	}

	db := database.ConnDB{
		Gorm: dbGORM,
		SqlX: dbSQLX,
	}

	product.RegisterProductService(router, validate, db)

	router.Listen(config.Cfg.App.Port)
}
