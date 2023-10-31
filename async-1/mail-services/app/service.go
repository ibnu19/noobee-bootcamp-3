package app

import (
	"github.com/go-playground/validator/v10"
)

type service struct {
	repository repository
	validate   *validator.Validate
}

func NewService(repository repository, validate *validator.Validate) service {
	return service{
		repository: repository,
		validate:   validate,
	}
}

func (service *service) SendEmail(request EmailRequest) (err error) {
	err = service.validate.Struct(request)
	if err != nil {
		return
	}

	email := Email{}
	ConvertToStruct(request, &email)
	return service.repository.SendEmail(email)
}
