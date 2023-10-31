package app

import (
	"async-1/app-services/utils"

	"github.com/gofiber/fiber/v2"
)

type service struct {
	repository
	URL string
}

func NewService(repository repository, url string) service {
	return service{
		repository: repository,
		URL:        url,
	}
}

func (service *service) SendEmail(c *fiber.Ctx, request EmailRequest) (response utils.WebResponse) {
	email := Email{}
	request.Type = "text/html"
	ConvertToStruct(request, &email)
	return service.repository.SendEmail(c, service.URL, email)
}
