package serviceinterfaces

import (
	productModel "github.com/anandtiwari11/IEKart-go/products/models"
	userModel "github.com/anandtiwari11/IEKart-go/user/models"
)

type IUserServiceInterface interface {
	FindUserByUsername(username string) (*userModel.User, error)
	FindUserByEmail(username string) (*userModel.User, error)
	CreateUser(user *userModel.User) error
	GenerateToken(user userModel.LoginUser) string
	GenerateJWT(email string) (string, error)
	GetAllProductOfTheUser(userId uint) (*[]productModel.Product, error)
	GetBookedProducts(userId uint) (*[]productModel.Product, error)
}
