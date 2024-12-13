package model

import "time"

type College struct {
	Id        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `json:"college_name"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}