package userModels

import (
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
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

func (u User) ValidateFields(c *fiber.Ctx) error {
	err := validate.Struct(u)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid fields", "data": nil})
	}

	return err
}

func (u *User) SetUUID() {
	u.ID = uuid.New()
}

func (u *User) HashPassword(c *fiber.Ctx) error {
	userPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Internal server error", "data": nil})
	}

	u.Password = string(userPassword)

	return err
}

func (u UpdateUser) ValidateFields(c *fiber.Ctx) error {
	err := validate.Struct(u)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid fields", "data": nil})
	}
	return err
}

func (u *UpdateUser) HashPassword(c *fiber.Ctx) error {
	if u.Password != "" {
		userPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Internal server error", "data": nil})
		}
		u.Password = string(userPassword)
	}

	return nil
}
