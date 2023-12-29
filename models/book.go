package models

import "time"

type Book struct {
	BookCode    string    `gorm:"primaryKey" json:"book_code"`
	BookTitle   string    `gorm:"not null" json:"book_title"`
	BookAuthor  string    `gorm:"not null" json:"book_author"`
	BookPub     string    `gorm:"not null" json:"book_publisher"`
	CatId       int       `gorm:"not null" json:"category"`
	BookPubDate string    `gorm:"not null" json:"book_publish_date"`
	BookLang    string    `gorm:"not null" json:"book_language"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type BookResponse struct {
	BookCode    string    `json:"book_code"`
	BookTitle   string    `json:"book_title"`
	BookAuthor  string    `json:"book_author"`
	BookPub     string    `json:"book_publisher"`
	Category    string    `json:"categoy"`
	BookPubDate string    `json:"book_publish_date"`
	BookLang    string    `json:"book_language"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
