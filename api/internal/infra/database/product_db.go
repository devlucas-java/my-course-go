package database

import (
	"github.com/devlucas-java/curso/api/internal/entity"
	"gorm.io/gorm"
)

type Product struct {
	DB *gorm.DB
}

func NewProduct(db *gorm.DB) *Product {
	return &Product{DB: db}
}

func (p *Product) CreateProduct(product *entity.Product) error {
	return p.DB.Create(product).Error
}

func (p *Product) FindById(id string) (*entity.Product, error) {
	var product entity.Product
	if err := p.DB.Where("id = ?", id).First(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (p *Product) UpdateProduct(product *entity.Product) error {
	var productFound entity.Product
	err := p.DB.Where("id = ?", product.ID).First(&productFound).Error
	if err == nil {
		return p.DB.Save(product).Error
	}
	return err
}

func (p *Product) DeleteProduct(id string) error {
	return p.DB.Where("id = ?", id).Delete(&entity.Product{}).Error

}

func (p *Product) FindAllProducts(page, limit int, sort string) ([]entity.Product, error) {

	var products []entity.Product

	if sort != "asc" && sort != "desc" {
		sort = "asc"
	}

	if page <= 0 {
		page = 1
	}

	if limit <= 0 || limit > 50 {
		limit = 10
	}

	offset := (page - 1) * limit

	err := p.DB.
		Order("created_at " + sort).
		Limit(limit).
		Offset(offset).
		Find(&products).Error

	return products, err
}
