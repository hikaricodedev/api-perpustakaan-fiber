package controllers

import (
	"api_perpustakaan/configs"
	"api_perpustakaan/models"

	"github.com/gofiber/fiber/v2"
)

func SearchBorrow(c *fiber.Ctx) error {
	code := c.Query("code")
	brw_date := c.Query("brw_date")

	var borrow []models.Borrow

	db := configs.ConnectDB() // Mengambil instance GORM dari local storage di Fiber

	query := db
	if code != "" {
		query = query.Where("brw_code LIKE ?", "%"+code+"%")
	}
	if brw_date != "" {
		query = query.Where("brw_date LIKE ?", "%"+brw_date+"%")
	}

	query.Find(&borrow)

	return c.JSON(borrow)
}

func GetSingleBorrow(c *fiber.Ctx) error {
	// userID := c.Params("id") sample get params
	// return c.SendString("User ID: " + userID)
	book_code := c.Params("code")

	var books models.Borrow

	db := configs.ConnectDB() // Mengambil instance GORM dari local storage di Fiber

	query := db

	query.Find(&books, book_code)

	return c.JSON(books)
}

func CreateBorrow(c *fiber.Ctx) error {
	var borrow models.Borrow
	db := configs.ConnectDB() // Mengambil instance GORM dari local storage di Fiber
	if err := c.BodyParser(&borrow); err != nil {
		return err
	}
	db.Create(&borrow)

	return c.JSON(borrow)
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
