package models

import (
	"time"
)

type Product struct {
    ID          uint   `json:"id" gorm:"primaryKey"`
    Name        string `json:"name"`
    Description string `json:"description"`
    CreatedAt   time.Time
    UpdatedAt   time.Time
}
