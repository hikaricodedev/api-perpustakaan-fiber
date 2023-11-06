package models

import "time"

type Return struct {
	RetCode   string `gorm:"primaryKey"`
	BrwCode   string `gorm:"not null"`
	RetStatus int    `gorm:"not null"`
	RetDate   int    `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
