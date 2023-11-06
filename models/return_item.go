package models

import "time"

type ReturnItem struct {
	RetiId     int    `gorm:"primaryKey"`
	RetCode    string `gorm:"not null"`
	BookCode   string `gorm:"not null"`
	RetiStatus int    `gorm:"not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
