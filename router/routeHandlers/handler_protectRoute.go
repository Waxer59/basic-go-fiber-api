package routeHandlers

import "github.com/gofiber/fiber/v2"

func ProtectRouteHandler(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusUnauthorized)
}
