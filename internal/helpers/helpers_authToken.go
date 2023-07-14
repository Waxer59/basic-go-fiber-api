package helpers

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

func GetJwtToken(c *fiber.Ctx) (string, error) {
	auth := c.Get("Authorization")
	if auth == "" {
		return "", fiber.ErrUnauthorized
	}
	return strings.Split(auth, "Bearer ")[1], nil
}
