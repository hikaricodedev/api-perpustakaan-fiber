package controllers

import (
	"api_perpustakaan/configs"
	"api_perpustakaan/models"

	"github.com/gofiber/fiber/v2"
)

func SearchBooks(c *fiber.Ctx) error {
	title := c.Query("title")
	author := c.Query("author")
	pub_date := c.Query("pub_date")

	var books []models.Book

	db := configs.ConnectDB() // Mengambil instance GORM dari local storage di Fiber

	query := db
	if title != "" {
		query = query.Where("book_title LIKE ?", "%"+title+"%")
	}
	if author != "" {
		query = query.Where("book_author LIKE ?", "%"+author+"%")
	}
	if pub_date != "" {
		query = query.Where("book_pub_date = ?", pub_date)
	}

	query.Find(&books)

	return c.JSON(books)
}
