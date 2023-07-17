package userController

import (
	"github.com/gofiber/fiber/v2"
	"github.com/waxer59/basic-go-fiber-api/internal/user/userModels"
	"github.com/waxer59/basic-go-fiber-api/internal/user/userService"
)

func Setup(router fiber.Router) {
	user := router.Group("/user")

	user.Get("/", userGetAll)

	user.Get("/:id", userGet)

	user.Delete("/:id", userDelete)

	user.Patch("/:id", userPatch)
}

// Update a user by ID
//	@Description	Update a user by ID
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	userModels.User
//	@Router			/users/:id [put]
func userPatch(c *fiber.Ctx) error {
	id := c.Params("id")

	var updateUser userModels.UpdateUser

	err := c.BodyParser(&updateUser)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't parse user", "data": nil})
	}

	user, err := userService.UpdateUser(id, updateUser)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": err.Error(), "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "User updated", "data": user})
}

// Delete a user by ID
//	@Description	Delete a user by ID
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	userModels.User
//	@Router			/users/:id [delete]
func userDelete(c *fiber.Ctx) error {
	id := c.Params("id")

	user, err := userService.DeleteUser(id)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": err.Error(), "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "User deleted", "data": user})
}

// Get a user
//	@Description	Get a user
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	userModels.User
//	@Router			/users:id [get]
func userGet(c *fiber.Ctx) error {
	id := c.Params("id")

	user, err := userService.GetUserById(id)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": err.Error(), "data": nil})
	}

	return c.Status(200).JSON(fiber.Map{"message": "User found", "status": "success", "data": user})
}

// Get all users
//	@Description	Get all users
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	userModels.User
//	@Router			/users [get]
func userGetAll(c *fiber.Ctx) error {
	users := userService.GetAllUsers()

	if len(users) <= 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No users found", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Users found", "data": users})
}
