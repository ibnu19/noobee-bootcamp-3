package main

import (
	"math/rand"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GetAllUsers(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(ApiResponses(nil, fiber.StatusOK, "get all success", Payload))
}

func CreateUsers(c *fiber.Ctx) error {
	rand.NewSource(int64(time.Now().Second()))
	user := new(User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ApiResponses(err, fiber.StatusBadRequest, fiber.ErrBadRequest.Message, nil))
	}

	user.Id = rand.Intn(100)
	Payload = append(Payload, *user)
	return c.Status(fiber.StatusCreated).JSON(ApiResponses(nil, fiber.StatusCreated, "created success", nil))
}

func UpdateUsers(c *fiber.Ctx) error {
	user := new(User)

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(ApiResponses(err, fiber.StatusNotFound, fiber.ErrNotFound.Message, nil))
	}

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ApiResponses(err, fiber.StatusBadRequest, fiber.ErrBadRequest.Message, user))
	}

	for i, p := range Payload {
		if id == p.Id {
			Payload[i].Name = user.Name
			Payload[i].Email = user.Email
			Payload[i].Address = user.Address
		}
	}

	return c.Status(fiber.StatusCreated).JSON(ApiResponses(nil, fiber.StatusCreated, "update success", nil))
}

func DeleteUsers(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(ApiResponses(err, fiber.StatusNotFound, fiber.ErrNotFound.Message, nil))
	}

	for i, p := range Payload {
		if id == p.Id {
			Payload = append(Payload[:i], Payload[i+1:]...)
		}
	}

	return c.Status(fiber.StatusOK).JSON(ApiResponses(nil, fiber.StatusOK, "delete success", nil))
}
