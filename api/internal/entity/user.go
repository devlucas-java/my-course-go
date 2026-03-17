package entity

import (
	"github.com/devlucas-java/curso/api/pkg/entity"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       entity.ID `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"-"`
}

func NewUser(name, email, password string) (*User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return &User{
		ID:       entity.NewUUID(),
		Name:     name,
		Email:    email,
		Password: string(hash),
	}, nil

}

func (u *User) ValidatePassword(str string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(str))
	return err
}
