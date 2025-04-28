package models

import "time"

type Member struct {
	MemId     int    `json:"mem_id" gorm:"primaryKey"`
	MemName   string `json:"mem_name" gorm:"not null"`
	MemEmail  string `json:"mem_email" gorm:"not null"`
	MemPhone  int    `json:"mem_phone" gorm:"not null"`
	MemBd     string `json:"mem_bd" gorm:"not null"`
	MemGend   string `json:"mem_gend" gorm:"not null"`
	MemAddr   string `json:"mem_addr" gorm:"not null"`
	MemStatus int    `json:"mem_status" gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
