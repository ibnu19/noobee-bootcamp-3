package product

import (
	"sesi-6/pkg/database"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func RegisterProductService(router fiber.Router, validate *validator.Validate, db database.ConnDB) {
	// repository := NewGormRepository(db.Gorm)
	repository := NewSqlXRepository(db.SqlX)
	service := NewService(repository, validate)
	handler := NewHandler(service)

	productRouter := router.Group("/products")
	{
		productRouter.Post("", handler.CreateProduct)
		productRouter.Put("/:id", handler.UpdateProduct)
		productRouter.Delete("/:id", handler.DeleteProduct)
		productRouter.Get("", handler.GetAllProduct)
		productRouter.Get("/:id", handler.GetOneProduct)
	}
}
