package product

type ProductRequest struct {
	Id       int    `json:"id"`
	Name     string `json:"name" validate:"required,min=3"`
	Category string `json:"category" validate:"required"`
	Price    int    `json:"price" validate:"required,numeric"`
	Stock    int    `json:"stock" validate:"required,numeric"`
}

type ProductResponse struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Price    int    `json:"price"`
	Stock    int    `json:"stock"`
}
