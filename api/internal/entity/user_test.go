package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("lucas", "lucas@lucas", "123456")

	assert.NoError(t, err)
	assert.NotNil(t, user.ID)
	assert.NotNil(t, user.Name)
	assert.NotNil(t, user.Email)
	assert.NotNil(t, user.Password)
	assert.Equal(t, user.Name, "lucas")
	assert.Equal(t, user.Email, "lucas@lucas")

}

func TestValidatePassword(t *testing.T) {
	user, err := NewUser("lucas", "lucas@lucas", "123456")

	assert.NoError(t, err)
	err = user.ValidatePassword("123456")

	assert.NoError(t, err)
	assert.NotNil(t, user.Password)
	assert.NoError(t, user.ValidatePassword("123456"))
	assert.Error(t, user.ValidatePassword("1234567"))
}
