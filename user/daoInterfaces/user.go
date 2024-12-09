package daointerfaces

import (
	productModel "github.com/anandtiwari11/IEKart-go/products/models"
	userModel "github.com/anandtiwari11/IEKart-go/user/models"
)

type IUserDAO interface {
	FindUserByUsername(username string) (*userModel.User, error)
	FindUserByEmail(username string) (*userModel.User, error)
	CreateUser(user *userModel.User) error
	GetAllProductOfTheUser(userId uint) (*[]productModel.Product, error)
}