package models

import "time"

type BorrowView struct {
	BrwCode   string `gorm:"primaryKey"`
	MemName   int    `gorm:"not null"`
	BrwDate   string `gorm:"not null"`
	BrwTime   int    `gorm:"not null"`
	BrwStatus int    `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
