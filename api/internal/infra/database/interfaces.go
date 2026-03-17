package database

import "github.com/devlucas-java/curso/api/internal/entity"

type UserInterface interface {
	CreateUser(user User) error
	FindUserByEmail(email string) (*entity.User, error)
	FindUserById(id string) (*entity.User, error)
	UpdateUser(user *entity.User) error
	FindAll(page, limit int) ([]entity.User, error)
}

type ProductInterface interface {
	CreateProduct(product *entity.Product) error
	FindById(id string) (*entity.Product, error)
	UpdateProduct(product *entity.Product) error
	FindAllProducts(page, limit int, sort string) ([]entity.Product, error)
	DeleteProduct(id string) error
}
