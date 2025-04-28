package controllers

import (
	"api_perpustakaan/configs"
	"api_perpustakaan/models"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func generateNumber() (string, error) {
	today := time.Now().Format("060102") // Format YYMMDD
	currentPrefix := fmt.Sprintf("BR.%s.", today)

	db := configs.ConnectDB()
	var lastTransaction models.Borrow
	if err := db.Where("brw_code LIKE ?", currentPrefix+"%").Order("brw_code DESC").First(&lastTransaction).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return currentPrefix + "0001", nil // Nomor transaksi pertama hari ini
		}
		return "", err
	}

	lastNumberStr := strings.TrimPrefix(lastTransaction.BrwCode, currentPrefix)
	lastNumber, err := strconv.Atoi(lastNumberStr)
	if err != nil {
		return "", fmt.Errorf("invalid last transaction number format: %w", err)
	}

	nextNumber := lastNumber + 1
	nextNumberStr := fmt.Sprintf("%04d", nextNumber) // Format dengan leading zeros
	return currentPrefix + nextNumberStr, nil

}

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
	newBrwCode, _ := generateNumber()
	borrow.BrwCode = newBrwCode

	db.Create(&borrow)
	brwItemsLen := len(borrow.BorrowItems)

	if brwItemsLen == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Item must be selected",
		})
	}

	// for _, bi := range borrow.BorrowItems {
	// 	fmt.Println("Buku : %s , Kode %d \n", bi.BookCode, bi.BrwiStatus)
	// 	db.Model(&borrow).Association("BorrowItems").Append(&bi)
	// }

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

func ReturnBook(c *fiber.Ctx) error {
	code := c.Params("code")

	var borrow models.Borrow
	db := configs.ConnectDB()
	if err := db.Where("brw_code = ? ", code).First(&borrow).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Borrow not found"})
	}

	// var updateBorrow models.Borrow
	// if err := c.BodyParser(&updateBorrow); err != nil {
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Fill the empty field!"})
	// }

	borrow.BrwStatus = 2

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
