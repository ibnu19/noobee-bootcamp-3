package product

// var (
// 	ErrEmptyName     = errors.New("kolom nama tidak boleh kosong")
// 	ErrEmptyCategory = errors.New("kolom category tidak boleh kosong")
// 	ErrEmptyPrice    = errors.New("kolom price tidak boleh kosong")
// 	ErrEmptyStock    = errors.New("kolom stock tidak boleh kosong")
// )

type Product struct {
	Id       int    `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	Category string `json:"category" db:"category"`
	Price    int    `json:"price" db:"price"`
	Stock    int    `json:"stock" db:"stock"`
}
