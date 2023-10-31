package app

import (
	"async-1/app-services/utils"
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

type repository struct{}

func NewRepository() repository {
	return repository{}
}

func (repository *repository) SendEmail(ctx *fiber.Ctx, URL string, email Email) (response utils.WebResponse) {
	agent := fiber.Post(URL)
	emailByte, err := json.Marshal(email)
	if err != nil {
		return utils.WebResponse{
			Success: false,
			Message: fiber.ErrBadRequest.Message,
			Error:   err.Error(),
		}
	}

	agent.Body(emailByte)
	agent.ContentType("application/json")
	statusCode, body, errs := agent.Bytes()
	if len(errs) > 0 {
		return utils.WebResponse{
			Success: false,
			Message: fiber.ErrBadRequest.Message,
			Error:   err.Error(),
		}
	}

	err = ctx.Status(statusCode).Send(body)
	if err != nil {
		return utils.WebResponse{
			Success: false,
			Message: fiber.ErrBadRequest.Message,
			Error:   err.Error(),
		}
	}

	var data utils.WebResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		return utils.WebResponse{
			Success: false,
			Message: fiber.ErrBadRequest.Message,
			Error:   err.Error(),
		}
	}

	return utils.WebResponse{
		Success: data.Success,
		Message: data.Message,
		Error:   data.Error,
	}
}
