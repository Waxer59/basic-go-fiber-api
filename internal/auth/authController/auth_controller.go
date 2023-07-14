package authController

import (
	"github.com/gofiber/fiber/v2"
	"github.com/waxer59/basic-go-fiber-api/internal/auth/authModels"
	"github.com/waxer59/basic-go-fiber-api/internal/auth/authService"
	"github.com/waxer59/basic-go-fiber-api/internal/user/userModels"
	"github.com/waxer59/basic-go-fiber-api/internal/user/userService"
)

func Setup(router fiber.Router) {
	auth := router.Group("/auth")

	auth.Post("/register", registerUser)

	auth.Post("/login", userLogin)
}

// Create a new user
// @Description Create a new user
// @Tags Auth
// @Accept json
// @Produce json
// @Success 200 {object} userModels.User
// @Router /auth/register [post]
func registerUser(c *fiber.Ctx) error {
	user := new(userModels.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't parse user", "data": nil})
	}

	if userEmail, err := userService.GetUserByEmail(user.Email); err == nil || userEmail != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "User already exists", "data": nil})
	}

	if err := user.HashPassword(); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't hash password", "data": nil})
	}

	_, err := userService.CreateUser(user)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create user", "data": nil})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "User created", "data": *user})
}

// Login a user
// @Description Login a user
// @Tags Auth
// @Accept json
// @Produce json
// @Success 200 {object} userModels.User
// @Router /auth/login [post]
func userLogin(c *fiber.Ctx) error {
	userLogin := new(authModels.UserLogin)

	if err := c.BodyParser(userLogin); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't parse user", "data": nil})
	}

	token, err := authService.UserLogin(userLogin.Email, userLogin.Password)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't login user", "data": nil})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "User logged in", "data": token})
}
