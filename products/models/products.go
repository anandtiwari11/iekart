package models

import (
	"time"
)

type Product struct {
    ID          uint       `gorm:"primaryKey"`
    Name        string     `json:"name" binding:"required"`
    Description string     `json:"description"`
    Price       float64    `json:"price" binding:"required"`
    SellerID    uint       `json:"seller_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
    BookedByID  *uint      `json:"booked_by_id"`
    BuyerID     *uint      `json:"buyer_id"`
    CreatedAt   time.Time  `json:"created_at" gorm:"autoCreateTime"`
    UpdatedAt   time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
}

type ProductReq struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"required"`
}