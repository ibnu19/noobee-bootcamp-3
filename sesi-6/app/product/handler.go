package product

import (
	"errors"
	"sesi-6/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) Handler {
	return Handler{
		service: service,
	}
}

// create new product
func (handler *Handler) CreateProduct(ctx *fiber.Ctx) error {
	request := ProductRequest{}
	err := ctx.BodyParser(&request)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(utils.ErrorResponse(fiber.ErrBadRequest.Message, err))
	}

	err = handler.service.Save(request)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(utils.ErrorResponse(fiber.ErrBadRequest.Message, err))
	}

	return ctx.Status(fiber.StatusCreated).
		JSON(utils.ApiResponse("CREATE SUCCESS", ""))
}

// update product
func (handler *Handler) UpdateProduct(ctx *fiber.Ctx) error {
	request := ProductRequest{}
	err := ctx.BodyParser(&request)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(utils.ErrorResponse(fiber.ErrBadRequest.Message, err))
	}

	productId, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(utils.ErrorResponse(fiber.ErrBadRequest.Message, errors.New("invalid id")))
	}

	request.Id = productId

	err = handler.service.Update(request)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(utils.ErrorResponse(fiber.ErrBadRequest.Message, err))
	}

	return ctx.Status(fiber.StatusCreated).
		JSON(utils.ApiResponse("UPDATE SUCCESS", ""))
}

// delete product by id
func (handler *Handler) DeleteProduct(ctx *fiber.Ctx) error {
	productId, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(utils.ErrorResponse(fiber.ErrBadRequest.Message, errors.New("invalid id")))
	}

	err = handler.service.Delete(productId)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).
			JSON(utils.ErrorResponse(fiber.ErrNotFound.Message, errors.New("data tidak ditemukan")))
	}

	return ctx.Status(fiber.StatusCreated).
		JSON(utils.ApiResponse("DELETE DATA SUCCESS", ""))
}

// get all products
func (handler *Handler) GetAllProduct(ctx *fiber.Ctx) error {
	products, err := handler.service.FindAll()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).
			JSON(utils.ErrorResponse(fiber.ErrInternalServerError.Message, errors.New("ada masalah pada server")))
	}

	return ctx.Status(fiber.StatusOK).
		JSON(utils.ApiResponse("GET ALL SUCCESS", products))
}

// get single product
func (handler *Handler) GetOneProduct(ctx *fiber.Ctx) error {
	productId, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(utils.ErrorResponse(fiber.ErrBadRequest.Message, errors.New("invalid id")))
	}

	product, err := handler.service.FindById(productId)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).
			JSON(utils.ErrorResponse(fiber.ErrNotFound.Message, errors.New("data tidak ditemukan")))
	}

	return ctx.Status(fiber.StatusOK).
		JSON(utils.ApiResponse("GET DATA SUCCESS", product))
}
