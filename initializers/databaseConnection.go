package initializers

import (
	"fmt"
	"log"

	userModel "github.com/anandtiwari11/IEKart-go/user/models"
	productModel "github.com/anandtiwari11/IEKart-go/products/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func ConnectDB() {
	DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to Connect to the Database", err)
	}
	err = DB.AutoMigrate(&userModel.User{}, &productModel.Product{})
	if err != nil {
		log.Fatal("Failed to Migrate into the table", err)
	}
	fmt.Println("Successfully connected to SQLite")
}
