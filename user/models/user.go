package models

import (
	"time"

	"github.com/anandtiwari11/IEKart-go/products/models"
)

type User struct {
    ID        uint             `gorm:"primaryKey"`
    Name      string           `json:"name"`
    Username  string           `gorm:"unique;not null" json:"username" binding:"required"`
    Email     string           `gorm:"unique;not null" json:"email" binding:"required,email"`
    Password  string           `gorm:"not null" json:"password" binding:"required,min=6"`
    Products  []models.Product        `gorm:"many2many:user_products;"`
    Customers []User           `gorm:"many2many:user_customers;joinForeignKey:UserID;joinReferences:CustomerID;"`
    CreatedAt time.Time        `json:"created_at" gorm:"autoCreateTime"`
    UpdatedAt time.Time        `json:"updated_at" gorm:"autoUpdateTime"`
    IsActive  bool             `json:"is_active" gorm:"default:true"`
}


type LoginUser struct {
	Username string `json:"username" gorm:"unique;not null"`
	Password string `json:"-" gorm:"not null"`
}
