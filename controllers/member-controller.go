package controllers

import (
	"api_perpustakaan/configs"
	"api_perpustakaan/models"

	"github.com/gofiber/fiber/v2"
)

func SearchMember(c *fiber.Ctx) error {
	member_name := c.Query("name")
	member_email := c.Query("email")
	member_phone := c.Query("phone")
	member_gender := c.Query("gender")
	member_status := c.Query("status")

	var members []models.Member

	db := configs.ConnectDB() // Mengambil instance GORM dari local storage di Fiber

	query := db
	if member_name != "" {
		query = query.Where("mem_name LIKE ?", "%"+member_name+"%")
	}
	if member_email != "" {
		query = query.Where("mem_email LIKE ?", "%"+member_email+"%")
	}
	if member_phone != "" {
		query = query.Where("mem_phone = ?", "%"+member_phone+"%")
	}
	if member_status != "" {
		query = query.Where("mem_status = ?", member_status)
	}
	if member_gender != "" {
		query = query.Where("mem_status = ?", member_gender)
	}

	query.Find(&members)

	return c.JSON(members)
}

func GetSingleMember(c *fiber.Ctx) error {
	// userID := c.Params("id") sample get params
	// return c.SendString("User ID: " + userID)
	member_code := c.Params("id")

	var members models.Member

	db := configs.ConnectDB() // Mengambil instance GORM dari local storage di Fiber

	query := db

	query.Find(&members, member_code)

	return c.JSON(members)
}

func CreateMember(c *fiber.Ctx) error {
	var member models.Member
	db := configs.ConnectDB() // Mengambil instance GORM dari local storage di Fiber
	if err := c.BodyParser(&member); err != nil {
		return err
	}
	db.Create(&member)

	return c.JSON(member)
}

func UpdateMember(c *fiber.Ctx) error {
	id := c.Params("id")

	var member models.Member
	db := configs.ConnectDB()
	if err := db.Where("mem_id = ? ", id).First(&member).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Member not found"})
	}

	var updateMember models.Member
	if err := c.BodyParser(&updateMember); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Fill the empty field!"})
	}

	member.MemName = updateMember.MemName
	member.MemEmail = updateMember.MemEmail
	member.MemPhone = updateMember.MemPhone
	member.MemBd = updateMember.MemBd
	member.MemGend = updateMember.MemGend
	member.MemAddr = updateMember.MemAddr
	member.MemStatus = updateMember.MemStatus

	if err := db.Save(&member).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal update data"})
	}

	return c.JSON(member)
}

func DeleteMember(c *fiber.Ctx) error {
	id := c.Params("id")

	var member models.Member
	db := configs.ConnectDB()
	if err := db.Where("mem_id = ? ", id).First(&member).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Member not found"})
	}

	if err := db.Delete(&member).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to delete user")
	}
	return c.SendString("Member deleted!")
}
