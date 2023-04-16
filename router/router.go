package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
	_ "github.com/waxer59/basic-go-fiber-api/docs"
	"github.com/waxer59/basic-go-fiber-api/internal/user/userController"
)

func Setup(app *fiber.App) {
	app.Get("/swagger/*", swagger.HandlerDefault)

	api := app.Group("/api", logger.New())

	userController.Setup(api)
}
