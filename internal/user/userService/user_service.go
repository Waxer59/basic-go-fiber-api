package userService

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/waxer59/basic-go-fiber-api/database"
	"github.com/waxer59/basic-go-fiber-api/internal/user/userModels"
)

func GetUser(c *fiber.Ctx) error {
	db := database.DB

	var user []userModels.User

	db.Find(&user)

	if len(user) <= 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No user found", "data": nil})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "User found", "data": user})
}

func CreateUser(c *fiber.Ctx) error {
	db := database.DB
	user := new(userModels.User)

	err := c.BodyParser(user)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't parse user", "data": nil})
	}

	user.SetUUID()
	user.ValidateFields(c)
	user.HashPassword(c)

	err = db.Create(&user).Error

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create user", "data": nil})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "User created", "data": user})
}

func UpdateUser(c *fiber.Ctx) error {
	db := database.DB

	var user userModels.User

	id := c.Params("id")

	db.Find(&user, "id = ?", id)

	if user.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No user found", "data": nil})
	}

	var updateUser userModels.UpdateUser

	err := c.BodyParser(&updateUser)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't parse user", "data": nil})
	}

	updateUser.ValidateFields(c)

	user.Name = updateUser.Name
	user.Email = updateUser.Email
	updateUser.HashPassword(c)

	db.Model(&user).Updates(&updateUser)

	return c.JSON(fiber.Map{"status": "success", "message": "User updated", "data": user})
}

func DeleteUser(c *fiber.Ctx) error {
	db := database.DB
	id := c.Params("id")

	var user userModels.User

	db.Find(&user, "id = ?", id)

	if user.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No user found", "data": nil})
	}

	err := db.Delete(&user, "id = ?", id).Error

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't delete user", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "User deleted", "data": nil})
}

func GetAllUsers(c *fiber.Ctx) error {
	db := database.DB

	var users []userModels.User

	db.Find(&users)

	if len(users) <= 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No users found", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Users found", "data": users})
}
