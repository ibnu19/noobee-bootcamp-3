package product

import "gorm.io/gorm"

type repositoryGorm struct {
	db *gorm.DB
}

func NewGormRepository(db *gorm.DB) Repository {
	return &repositoryGorm{
		db: db,
	}
}

func (repository *repositoryGorm) Save(product Product) (err error) {
	return repository.db.Create(&product).Error
}

func (repository *repositoryGorm) Update(product Product) (err error) {
	return repository.db.Updates(&product).Error
}

func (repository *repositoryGorm) Delete(product Product) (err error) {
	return repository.db.Delete(&product).Error
}

func (repository *repositoryGorm) FindAll() (products []Product, err error) {
	err = repository.db.Find(&products).Error
	return
}

func (repository *repositoryGorm) FindById(id int) (product Product, err error) {
	err = repository.db.Where("id = ?", id).First(&product).Error
	return
}
