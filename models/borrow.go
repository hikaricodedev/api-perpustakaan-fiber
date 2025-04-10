package models

import (
	"time"
)

type Borrow struct {
	BrwCode     string `gorm:"primaryKey" json:"brw_code"`
	MemId       int    `gorm:"not null" json:"mem_id"`
	BrwDate     string `gorm:"not null" json:"brw_date"`
	BrwTime     int    `gorm:"not null" json:"brw_time"`
	BrwStatus   int    `gorm:"not null" json:"brw_status"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	BorrowItems []BorrowItem `gorm:"foreignKey:BrwCode" json:"brw_items" binding:"-"`
}
