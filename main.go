package main

import (
	"os"
	"fmt"
	"log"

	"github.com/waxer59/basic-go-fiber-api/config"
	"github.com/waxer59/basic-go-fiber-api/database"
	"github.com/waxer59/basic-go-fiber-api/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/waxer59/basic-go-fiber-api/docs"
)

// @title			Basic Go Fiber API
// @version		1.0
// @description	This is a basic go fiber api
// @BasePath		/api
func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file. Err: %s", err)
	}

	app := fiber.New(fiber.Config{
		AppName:       "Fiber App",
		CaseSensitive: true,
	})

	database.Connect()

	middlewares(app)

	router.Setup(app)

	log.Fatal(app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT"))))
}

func middlewares(app *fiber.App) {
	app.Use(cors.New(cors.Config{}))
}
