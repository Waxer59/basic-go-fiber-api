package main

import (
	"fmt"
	"log"

	"github.com/waxer59/basic-go-fiber-api/config"
	"github.com/waxer59/basic-go-fiber-api/database"
	"github.com/waxer59/basic-go-fiber-api/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName:       "Fiber App",
		CaseSensitive: true,
	})

	database.Connect()

	middlewares(app)

	router.Setup(app)

	log.Fatal(app.Listen(fmt.Sprintf(":%s", config.GetEnv("PORT"))))
}

func middlewares(app *fiber.App) {
	app.Use(cors.New(cors.Config{}))
}
