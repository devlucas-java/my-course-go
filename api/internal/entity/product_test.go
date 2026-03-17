package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	product := NewProduct("vr", 300.0)

	assert.NotEmpty(t, product.ID)
	assert.NotNil(t, product.Name)
	assert.NotNil(t, product.Price)
	assert.NotNil(t, product.CreatedAt)
	assert.Equal(t, product.Name, "vr")
	assert.Equal(t, product.Price, 300.0)
}

func TestProduct_Validate(t *testing.T) {
	product := NewProduct("vr", 300.0)

	assert.NoError(t, product.Validate())
}

func TestProduct_ValidateError(t *testing.T) {
	product := NewProduct("", 300.0)

	assert.Error(t, product.Validate())
}
