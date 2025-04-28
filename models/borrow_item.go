package models

import "time"

type BorrowItem struct {
	BrwiId     int    `gorm:"primaryKey" json:"brwi_id"`
	BrwCode    string `gorm:"not null" json:"brw_code"`
	BookCode   string `gorm:"not null" json:"book_code"`
	BrwiStatus int    `gorm:"not null" json:"brwi_status"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
