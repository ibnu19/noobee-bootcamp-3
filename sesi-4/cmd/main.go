package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

var Payload = []User{}

func main() {
	router := fiber.New(fiber.Config{
		Prefork: true,
		AppName: "Noobee",
	})

	router.Use(requestid.New(requestid.Config{
		Header:     "X-TRACE-ID",
		ContextKey: "X-TRACE-ID",
	}))

	router.Use(TraceId())

	setupHandler(router)
	router.Listen(":8080")
}

func setupHandler(router fiber.Router) {
	userRouter := router.Group("/users")
	{
		userRouter.Get("", GetAllUsers)
		userRouter.Post("", CreateUsers)
		userRouter.Put("/:id", UpdateUsers)
		userRouter.Delete("/:id", DeleteUsers)
	}
}
