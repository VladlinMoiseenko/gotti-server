package main

import (
	"log"
	"os"

	"github.com/VladlinMoiseenko/gotti/db"
	"github.com/VladlinMoiseenko/gotti/routes"

	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"

	_ "github.com/VladlinMoiseenko/gotti/docs/gottiserver"
	swagger "github.com/arsmn/fiber-swagger/v2"
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

// @title Fiber Swagger Example API
// @version 2.0
// @description This is a sample server.

// @host localhost:9000
// @BasePath /
// @schemes http
func main() {

	godotenv.Load(".env.local")

	godotenv.Load()

	app := fiber.New()

	prometheus := fiberprometheus.New("gotti-server")
	prometheus.RegisterAt(app, "/metrics")
	app.Use(prometheus.Middleware)

	app.Use(cors.New())
	app.Use(logger.New())
	app.Use(limiter.New())

	db.ConnectDB()

	app.Get("/swagger/*", swagger.HandlerDefault)

	setupRoutes(app)

	port := os.Getenv("PORT")
	err := app.Listen(":" + port)

	if err != nil {
		log.Fatal("Error app failed to start")
		panic(err)
	}
}
