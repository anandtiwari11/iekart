package daointerfaces

import (
	productModel "github.com/anandtiwari11/IEKart-go/products/models"
	userModel "github.com/anandtiwari11/IEKart-go/user/models"
)

type IProductDao interface {
	CreateProduct(product *productModel.Product, user *userModel.User) error
	GetProductDetailsByProductId(productId uint) (*productModel.Product, error)
	DeleteProductByID(productId uint) error
	BookProduct(productId uint, user *userModel.User) error
	BuyProduct(productId uint, user *userModel.User) error
}