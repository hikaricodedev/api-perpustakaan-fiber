package models

import "time"

type Book struct {
	BookCode    string `gorm:"primaryKey"`
	BookTitle   string `gorm:"not null"`
	BookAuthor  string `gorm:"not null"`
	BookPub     string `gorm:"not null"`
	CatId       int    `gorm:"not null"`
	BookPubDate string `gorm:"not null"`
	BookLang    string `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
