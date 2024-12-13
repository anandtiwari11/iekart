package servicesinterfaces

import (
	productModel "github.com/anandtiwari11/IEKart-go/products/models"
	userModel "github.com/anandtiwari11/IEKart-go/user/models"
)

type IProductIntf interface {
	CreateProduct(req *productModel.ProductReq, user *userModel.User) error
	GetProductDetailsByProductId(productId uint) (*productModel.Product, error)
	DeleteProductByID(productId uint, userId uint) error
	BookProduct(productId uint, user *userModel.User) error
	BuyProduct(productId uint, user *userModel.User) error
}