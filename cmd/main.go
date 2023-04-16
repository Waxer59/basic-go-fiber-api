package main

import (
	"crud_api/config"
	"crud_api/database"
	"fmt"
	"log"

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

	serverPort := fmt.Sprintf(":%s", config.GetEnv("PORT"))

	log.Fatal(app.Listen(serverPort))
}

func middlewares(app *fiber.App) {
	app.Use(cors.New(cors.Config{}))
}
