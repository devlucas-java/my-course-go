package entity

import (
	"errors"
	"time"

	"github.com/devlucas-java/curso/api/pkg/entity"
)

type Product struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

var (
	ErrNotFound          = errors.New("Not found")
	ErrIDInvalid         = errors.New("Id is invalid")
	ErrNameRequired      = errors.New("Name is required")
	ErrPriceRequired     = errors.New("Price is required")
	ErrCreatedAtRequired = errors.New("Created at is required")
)

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:        entity.NewUUID(),
		Name:      name,
		Price:     price,
		CreatedAt: time.Now(),
	}
}

func (p *Product) Validate() error {
	if p.Name == "" {
		return ErrNameRequired
	}
	if p.Price == 0 {
		return ErrPriceRequired
	}
	if p.CreatedAt.IsZero() {
		return ErrCreatedAtRequired
	}
	return nil
}
