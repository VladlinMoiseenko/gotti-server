package main

import (
	"log"
	"os"

	"github.com/VladlinMoiseenko/gotti/db"
	"github.com/VladlinMoiseenko/gotti/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success":     true,
			"message":     "You are at the root endpoint",
			"github_repo": "https://github.com/VladlinMoiseenko/gotti-server",
		})
	})

	api := app.Group("/api")

	routes.GottiRoute(api.Group("/gotti"))
}

func main() {

	erro := godotenv.Load()
	if erro != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()

	app.Use(cors.New())
	app.Use(logger.New())

	db.ConnectDB()

	setupRoutes(app)

	port := os.Getenv("PORT")
	err := app.Listen(":" + port)

	if err != nil {
		log.Fatal("Error app failed to start")
		panic(err)
	}
}
