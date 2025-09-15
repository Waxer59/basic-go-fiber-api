package userModels

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/waxer59/basic-go-fiber-api/internal/upload/uploadModels"
	"github.com/waxer59/basic-go-fiber-api/internal/utils/bcryptUtils"
	"gorm.io/gorm"
)

var validate = validator.New()

type User struct {
	ID       uuid.UUID             `gorm:"type:uuid;primaryKey"`
	Name     string                `json:"name" gorm:"not null" validate:"required"`
	Email    string                `json:"email" gorm:"unique;not null" validate:"required,email"`
	Password string                `json:"password" gorm:"not null" validate:"required"` // omited on response
	Uploads  []uploadModels.Upload `json:"-"`
}

type UpdateUser struct {
	Name     string  `json:"name"`
	Email    *string `json:"email,omitempty" validate:"omitempty,email"`
	Password string  `json:"-"` // omited on response
}

func (u *User) BeforeSave(tx *gorm.DB) error {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return nil
}

func (u User) ValidateFields() error {
	err := validate.Struct(u)

	if err != nil {
		return err
	}

	return nil
}

func (u *User) HashPassword() error {
	userPassword, err := bcryptUtils.Encript(u.Password)

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
		userPassword, err := bcryptUtils.Encript(u.Password)
		if err != nil {
			return err
		}
		u.Password = string(userPassword)
	}

	return nil
}
