package product

import (
	"github.com/jmoiron/sqlx"
)

type sqlxRepository struct {
	db *sqlx.DB
}

func NewSqlXRepository(db *sqlx.DB) Repository {
	return &sqlxRepository{
		db: db,
	}
}

func (repository *sqlxRepository) Save(product Product) (err error) {
	query := `
		INSERT INTO products(
			name, category, price, stock
		) VALUES (
			:name, :category, :price, :stock
		)
	`
	_, err = repository.db.NamedExec(query, product)
	return
}

func (repository *sqlxRepository) Update(product Product) (err error) {
	query := `
		UPDATE products
		SET name = :name,
			category = :category,
			price = :price,
			stock = :stock
		WHERE id = :id
	`
	_, err = repository.db.NamedExec(query, product)
	return
}

func (repository *sqlxRepository) Delete(product Product) (err error) {
	query := `DELETE FROM products WHERE id=:id`
	_, err = repository.db.NamedExec(query, product)
	return
}

func (repository *sqlxRepository) FindAll() (products []Product, err error) {
	query := `SELECT id, name, category, price, stock FROM products ORDER BY id`
	err = repository.db.Select(&products, query)
	return
}

func (repository *sqlxRepository) FindById(id int) (product Product, err error) {
	query := `SELECT id, name, category, price, stock FROM products WHERE id=$1`
	err = repository.db.Get(&product, query, id)
	return
}
