package controllers

import (
	"api_perpustakaan/configs"
	"api_perpustakaan/models"

	"github.com/gofiber/fiber/v2"
)

func SearchCategory(c *fiber.Ctx) error {
	cat_name := c.Query("cat_name")

	var category []models.Category

	db := configs.ConnectDB() // Mengambil instance GORM dari local storage di Fiber

	query := db
	if cat_name != "" {
		query = query.Where("cat_name LIKE ?", "%"+cat_name+"%")
	}

	query.Find(&category)

	return c.JSON(category)
}

func GetSingleCategory(c *fiber.Ctx) error {
	// userID := c.Params("id") sample get params
	// return c.SendString("User ID: " + userID)
	book_code := c.Params("id")

	var books models.Category

	db := configs.ConnectDB() // Mengambil instance GORM dari local storage di Fiber

	query := db

	query.Find(&books, book_code)

	return c.JSON(books)
}

func CreateCategory(c *fiber.Ctx) error {
	var book models.Category
	db := configs.ConnectDB() // Mengambil instance GORM dari local storage di Fiber
	if err := c.BodyParser(&book); err != nil {
		return err
	}
	db.Create(&book)

	return c.JSON(book)
}

func UpdateCategory(c *fiber.Ctx) error {
	id := c.Params("id")

	var category models.Category
	db := configs.ConnectDB()
	if err := db.Where("cat_id = ? ", id).First(&category).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Category not found"})
	}

	var updateCategory models.Category
	if err := c.BodyParser(&updateCategory); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Fill the empty field!"})
	}

	category.CatId = updateCategory.CatId
	category.CatName = updateCategory.CatName
	category.CatStatus = updateCategory.CatStatus
	category.CatOrder = updateCategory.CatOrder

	if err := db.Save(&category).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal update data"})
	}

	return c.JSON(category)
}

func DeleteCategory(c *fiber.Ctx) error {
	id := c.Params("id")

	var category models.Category
	db := configs.ConnectDB()
	if err := db.Where("id = ? ", id).First(&id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Category not found"})
	}

	if err := db.Delete(&category).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to delete category")
	}
	return c.SendString("Category deleted!")
}
