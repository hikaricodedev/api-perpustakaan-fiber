package main

import (
	"os"

	"api_perpustakaan/configs"
	"api_perpustakaan/models"

	"github.com/gofiber/fiber/v2"
	goenv "github.com/subosito/gotenv"
)

func migrate() {
	db := configs.ConnectDB()

	db.AutoMigrate(&models.Book{})
	db.AutoMigrate(&models.Borrow{})
	db.AutoMigrate(&models.BorrowItem{})
	db.AutoMigrate(&models.Member{})
	db.AutoMigrate(&models.Return{})
	db.AutoMigrate(&models.ReturnItem{})
	db.AutoMigrate(&models.Category{})
}

func main() {
	goenv.Load(".env")
	if os.Getenv("DO_MIGRATION") == "1" {
		migrate()
	}
	app := fiber.New()
	//refference
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	//refference
	app.Listen(":3000")
}
