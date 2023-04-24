package uploadController

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/skip"
	"github.com/waxer59/basic-go-fiber-api/router/routeHandlers"
	"github.com/waxer59/basic-go-fiber-api/router/routeMiddlewares"
)

func Setup(router fiber.Router) {
	upload := router.Group("/upload", skip.New(routeHandlers.ProtectRouteHandler, routeMiddlewares.ProtectRouteMiddleware))

	upload.Post("/", func(c *fiber.Ctx) error {
		return c.SendString("upload")
	})

}
