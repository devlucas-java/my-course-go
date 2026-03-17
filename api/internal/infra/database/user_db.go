package database

import (
	"github.com/devlucas-java/curso/api/internal/entity"
	"gorm.io/gorm"
)

type User struct {
	DB *gorm.DB
}

func NewUser(db *gorm.DB) *User {
	return &User{DB: db}
}

func (u *User) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	if err := u.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *User) FindById(id string) (*entity.User, error) {
	var user entity.User
	return &user, u.DB.Where("id = ?", id).First(&user).Error
}

func (u *User) CreateUser(user *entity.User) error {
	return u.DB.Create(user).Error
}

func (u *User) UpdateUser(user *entity.User) (*entity.User, error) {
	var userFound entity.User

	err := u.DB.Where("id = ?", user.ID).First(userFound).Error
	if err != nil {
		return nil, err
	}
	err = u.DB.Save(user).Error
	return &userFound, err
}
