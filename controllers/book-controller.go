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

func GetSingleBook(c *fiber.Ctx) error {
	// userID := c.Params("id") sample get params
	// return c.SendString("User ID: " + userID)
	book_code := c.Params("code")

	var books models.Book

	db := configs.ConnectDB() // Mengambil instance GORM dari local storage di Fiber

	query := db
	query = query.Where("book_code", book_code)
	query.First(&books)

	return c.JSON(books)
}

func CreateBook(c *fiber.Ctx) error {
	var book models.Book
	db := configs.ConnectDB() // Mengambil instance GORM dari local storage di Fiber
	if err := c.BodyParser(&book); err != nil {
		return err
	}
	db.Create(&book)

	return c.JSON(book)
}

func UpdateBook(c *fiber.Ctx) error {
	code := c.Params("code")

	var book models.Book
	db := configs.ConnectDB()
	if err := db.Where("book_code = ? ", code).First(&book).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Book not found"})
	}

	var updateBook models.Book
	if err := c.BodyParser(&updateBook); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Fill the empty field!"})
	}

	book.BookTitle = updateBook.BookTitle
	book.BookAuthor = updateBook.BookAuthor
	book.BookPub = updateBook.BookPub
	book.CatId = updateBook.CatId
	book.BookPubDate = updateBook.BookPubDate
	book.BookLang = updateBook.BookLang

	if err := db.Save(&book).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal update data"})
	}

	return c.JSON(book)
}

func DeleteBook(c *fiber.Ctx) error {
	code := c.Params("code")

	var book models.Book
	db := configs.ConnectDB()
	if err := db.Where("book_code = ? ", code).First(&book).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Book not found"})
	}

	if err := db.Delete(&book).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to delete user")
	}
	return c.SendString("Book deleted!")
}
