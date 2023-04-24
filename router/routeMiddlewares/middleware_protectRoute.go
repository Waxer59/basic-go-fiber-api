package routeMiddlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/waxer59/basic-go-fiber-api/internal/utils/jwtUtils"
)

const AUTHORIZATION_HEADER = "Bearer-Token"

func ProtectRouteMiddleware(c *fiber.Ctx) bool {
	headerToken := c.Get(AUTHORIZATION_HEADER)

	_, err := jwtUtils.ParseJwt(headerToken)

	if err != nil {
		return false
	}

	return true
}
