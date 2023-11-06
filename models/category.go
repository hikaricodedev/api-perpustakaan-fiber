package models

import "time"

type Category struct {
	CatId     int `gorm:"primaryKey"`
	CatName   int `gorm:"not null"`
	CatStatus int `gorm:"not null"`
	CatOrder  int `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
