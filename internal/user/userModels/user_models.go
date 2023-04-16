package userModels

import (
	"github.com/go-playground/validator"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Name     string    `json:"name" gorm:"not null" validate:"required"`
	Email    string    `json:"email" gorm:"unique;not null" validate:"required,email"`
	Password string    `json:"password" gorm:"not null" validate:"required"`
}

type UpdateUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var validate = validator.New()

func (u User) Validate() error {
	err := validate.Struct(u)

	return err
}
