package userController

import (
	"github.com/gofiber/fiber/v2"
	"github.com/waxer59/basic-go-fiber-api/internal/user/userService"
)

func Setup(router fiber.Router) {
	user := router.Group("/user")

	user.Post("/", userService.CreateUser)

	user.Get("/", userService.GetAllUsers)

	user.Get("/:id", userService.GetUser)

	user.Delete("/:id", userService.DeleteUser)

	user.Patch("/:id", userService.UpdateUser)
}
