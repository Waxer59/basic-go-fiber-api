package userModels

import (
	"github.com/go-playground/validator"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var validate = validator.New()

type User struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name     string    `json:"name" gorm:"not null" validate:"required"`
	Email    string    `json:"email" gorm:"unique;not null" validate:"required,email"`
	Password string    `json:"password" gorm:"not null" validate:"required"`
}

type UpdateUser struct {
	Name     string `json:"name"`
	Email    string `json:"email" validate:"email"`
	Password string `json:"password"`
}

func (u User) ValidateFields() error {
	err := validate.Struct(u)

	if err != nil {
		return err
	}

	return nil
}

func (u *User) SetUUID() {
	u.ID = uuid.New()
}

func (u *User) HashPassword() error {
	userPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	u.Password = string(userPassword)

	return nil
}

func (u UpdateUser) ValidateFields() error {
	err := validate.Struct(u)
	if err != nil {
		return err
	}
	return nil
}

func (u *UpdateUser) HashPassword() error {
	if u.Password != "" {
		userPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		u.Password = string(userPassword)
	}

	return nil
}
