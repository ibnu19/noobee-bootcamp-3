package product

func ConvertToStruct(request ProductRequest, product *Product) {
	product.Id = request.Id
	product.Name = request.Name
	product.Category = request.Category
	product.Price = request.Price
	product.Stock = request.Stock
}
