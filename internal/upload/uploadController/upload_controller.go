package uploadController

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/skip"
	"github.com/waxer59/basic-go-fiber-api/internal/upload/uploadService"
	"github.com/waxer59/basic-go-fiber-api/router/routeHandlers"
	"github.com/waxer59/basic-go-fiber-api/router/routeMiddlewares"
)

const FORM_PARAM_NAME = "file"

func Setup(router fiber.Router) {
	upload := router.Group("/upload", skip.New(routeHandlers.ProtectRouteHandler, routeMiddlewares.ProtectRouteMiddleware))

	upload.Post("/", func(c *fiber.Ctx) error {
		c.Accepts("multipart/form-data")

		file, err := c.FormFile(FORM_PARAM_NAME)

		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		uploadService.CreateUpload(file)

		return c.SendString("upload")
	})

}
