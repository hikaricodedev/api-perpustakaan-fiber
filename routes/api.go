package routes

import (
	"api_perpustakaan/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/books", controllers.SearchBooks)
	api.Post("/books/store", controllers.CreateBook)
}
