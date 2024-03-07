package main

import (
	"fmt"
	"os"

	"api_perpustakaan/configs"
	"api_perpustakaan/models"
	"api_perpustakaan/routes"

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

	sqlStr := `SELECT b.brw_code AS brw_code, m.mem_name AS mem_name , b.brw_date AS brw_date , b.brw_status AS brw_status, b.created_at AS created_at , b.updated_at AS updated_at FROM borrows b LEFT JOIN members m ON b.mem_id = m.mem_id`
	args := '-'

	tx := db.Exec(sqlStr, args)

	if tx.Error != nil {
		fmt.Println("error migrate")
	}
}

func main() {
	goenv.Load(".env")
	if os.Getenv("DO_MIGRATION") == "1" {
		migrate()
	}
	app := fiber.New()

	routes.SetupRoutes(app)
	app.Listen(":" + os.Getenv("RUNNING_PORT"))
}
