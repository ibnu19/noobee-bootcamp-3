package product

import (
	"github.com/go-playground/validator/v10"
)

type Service struct {
	repository Repository
	validate   *validator.Validate
}

func NewService(repository Repository, validate *validator.Validate) Service {
	return Service{
		repository: repository,
		validate:   validate,
	}
}

func (service *Service) Save(request ProductRequest) (err error) {
	err = service.validate.Struct(request)
	if err != nil {
		return
	}

	product := Product{}
	ConvertToStruct(request, &product)
	return service.repository.Save(product)
}

func (service *Service) Update(request ProductRequest) (err error) {
	err = service.validate.Struct(request)
	if err != nil {
		return
	}

	product, err := service.repository.FindById(request.Id)
	if err != nil {
		return
	}

	ConvertToStruct(request, &product)
	return service.repository.Update(product)
}

func (service *Service) Delete(id int) (err error) {
	product, err := service.repository.FindById(id)
	if err != nil {
		return
	}
	return service.repository.Delete(product)
}

func (service *Service) FindAll() (products []Product, err error) {
	return service.repository.FindAll()
}

func (service *Service) FindById(id int) (product Product, err error) {
	return service.repository.FindById(id)
}
