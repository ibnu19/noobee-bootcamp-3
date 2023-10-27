package product

type Repository interface {
	Save(product Product) (err error)
	Update(product Product) (err error)
	Delete(product Product) (err error)
	FindAll() (products []Product, err error)
	FindById(id int) (product Product, err error)
}
