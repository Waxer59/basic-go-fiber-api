package routeMiddlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/waxer59/basic-go-fiber-api/internal/helpers"
	"github.com/waxer59/basic-go-fiber-api/internal/utils/jwtUtils"
)

func ProtectRouteMiddleware(c *fiber.Ctx) bool {

	token, err := helpers.GetJwtToken(c)

	if err != nil {
		return false
	}

	_, err = jwtUtils.ParseJwt(token)

	if err != nil {
		return false
	}

	return true
}
