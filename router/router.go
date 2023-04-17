package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/swagger"
	"github.com/waxer59/basic-go-fiber-api/internal/auth/authController"
	"github.com/waxer59/basic-go-fiber-api/internal/user/userController"
)

func Setup(app *fiber.App) {
	app.Get("/swagger/*", swagger.HandlerDefault)

	app.Get("/metrics", monitor.New(monitor.Config{Title: "Fiber Metrics"}))

	api := app.Group("/api", logger.New())

	userController.Setup(api)
	authController.Setup(api)
}
