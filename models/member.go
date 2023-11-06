package models

import "time"

type Member struct {
	MemId     int    `gorm:"primaryKey"`
	MemName   string `gorm:"not null"`
	MemEmail  string `gorm:"not null"`
	MemPhone  int    `gorm:"not null"`
	MemBd     string `gorm:"not null"`
	MemGend   string `gorm:"not null"`
	MemAddr   string `gorm:"not null"`
	MemStatus int    `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
