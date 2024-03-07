package controllers

import (
	"api_perpustakaan/configs"
	"api_perpustakaan/models"

	"github.com/gofiber/fiber/v2"
)

func SearchBorrow(c *fiber.Ctx) error {
	title := c.Query("title")
	author := c.Query("author")
	pub_date := c.Query("pub_date")

	var books []models.Borrow

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

func GetSingleBorrow(c *fiber.Ctx) error {
	// userID := c.Params("id") sample get params
	// return c.SendString("User ID: " + userID)
	book_code := c.Params("id")

	var books models.Borrow

	db := configs.ConnectDB() // Mengambil instance GORM dari local storage di Fiber

	query := db

	query.Find(&books, book_code)

	return c.JSON(books)
}

func CreateBorrow(c *fiber.Ctx) error {
	var book models.Borrow
	db := configs.ConnectDB() // Mengambil instance GORM dari local storage di Fiber
	if err := c.BodyParser(&book); err != nil {
		return err
	}
	db.Create(&book)

	return c.JSON(book)
}

func UpdateBorrow(c *fiber.Ctx) error {
	code := c.Params("code")

	var borrow models.Borrow
	db := configs.ConnectDB()
	if err := db.Where("brw_code = ? ", code).First(&borrow).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Borrow not found"})
	}

	var updateBorrow models.Borrow
	if err := c.BodyParser(&updateBorrow); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Fill the empty field!"})
	}

	borrow.MemId = updateBorrow.MemId
	borrow.BrwDate = updateBorrow.BrwDate
	borrow.BrwTime = updateBorrow.BrwTime
	borrow.BrwStatus = updateBorrow.BrwStatus

	if err := db.Save(&borrow).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal update data"})
	}

	return c.JSON(borrow)
}

func DeleteBorrow(c *fiber.Ctx) error {
	code := c.Params("code")

	var borrow models.Borrow
	db := configs.ConnectDB()
	if err := db.Where("brw_code = ? ", code).First(&borrow).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Borrow not found"})
	}

	if err := db.Delete(&borrow).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to delete user")
	}
	return c.SendString("Borrow deleted!")
}
