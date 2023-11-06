package models

import "time"

type BorrowItem struct {
	BrwiId     int    `gorm:"primaryKey"`
	BrwCode    string `gorm:"not null"`
	BookCode   string `gorm:"not null"`
	BrwiStatus int    `gorm:"not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
