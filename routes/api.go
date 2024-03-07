package routes

import (
	"api_perpustakaan/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/books", controllers.SearchBooks)
	api.Get("/books/:id", controllers.GetSingleBook)
	api.Post("/books/store", controllers.CreateBook)
	api.Put("/books/:code/update", controllers.UpdateBook)
	api.Delete("/books/:code", controllers.DeleteBook)
	api.Get("/category", controllers.SearchCategory)
	api.Get("/category/:id", controllers.GetSingleCategory)
	api.Post("/category/store", controllers.CreateCategory)
	api.Put("/category/:id/update", controllers.UpdateCategory)
	api.Delete("/category/:code", controllers.DeleteCategory)
	api.Get("/member", controllers.SearchMember)
	api.Get("/member/:id", controllers.GetSingleMember)
	api.Post("/member/store", controllers.CreateMember)
	api.Put("/member/:id/update", controllers.UpdateMember)
	api.Delete("/member/:id", controllers.DeleteMember)
}
