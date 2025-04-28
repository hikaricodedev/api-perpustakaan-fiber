package routes

import (
	"api_perpustakaan/controllers"
	"api_perpustakaan/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/books", controllers.SearchBooks)
	api.Get("/books/:code", controllers.GetSingleBook)
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
	api.Get("/borrow", controllers.SearchBorrow)
	api.Get("/borrow/:code", controllers.GetSingleBorrow)
	api.Post("/borrow/store", controllers.CreateBorrow)
	api.Put("/borrow/:code/update", controllers.UpdateBorrow)
	api.Put("/borrow/:code/return", controllers.ReturnBook)
	api.Delete("/borrow/:code", middleware.Protected, controllers.DeleteBorrow)
}
