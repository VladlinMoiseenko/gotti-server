package routes

import (
	"github.com/VladlinMoiseenko/gotti/controllers"

	"github.com/gofiber/fiber/v2"
)

func GottiRoute(route fiber.Router) {
	route.Get("/", controllers.GetAllanimation)
	route.Get("/:id", controllers.GetAnimation)
	route.Post("/", controllers.AddAnimation)
	//route.Put("/:id", controllers.UpdateAnimation)
	//route.Delete("/:id", controllers.DeleteAnimation)
}
