package services

import (
	"fmt"
	"time"

	daointerfaces "github.com/anandtiwari11/IEKart-go/products/daoInterfaces"
	productModel "github.com/anandtiwari11/IEKart-go/products/models"
	userModel "github.com/anandtiwari11/IEKart-go/user/models"
)

type ProductService struct {
	ProductDao daointerfaces.IProductDao
}

func NewProductService (ProductDao daointerfaces.IProductDao) *ProductService {
	return &ProductService{
		ProductDao: ProductDao,
	}
}

func (productService *ProductService) CreateProduct(req *productModel.ProductReq, user *userModel.User) error {
	newProduct := productModel.Product{
        Name:        req.Name,
        Description: req.Description,
        Price:       req.Price,
        SellerID:    user.ID,
        CreatedAt:   time.Now(),
        UpdatedAt:   time.Now(),
    }
	if err := productService.ProductDao.CreateProduct(&newProduct, user); err != nil {
		return fmt.Errorf("failed to create product: %v", err)
	}
	return nil
}

func (productService *ProductService) GetProductDetailsByProductId(productId uint) (*productModel.Product, error) {
	return productService.ProductDao.GetProductDetailsByProductId(productId)
}

func (productService *ProductService) DeleteProductByID(productId uint) error {
	return productService.ProductDao.DeleteProductByID(productId)
}

func (productService *ProductService) BookProduct(productId uint, user *userModel.User) error {
	return productService.ProductDao.BookProduct(productId, user)
}

func (productService *ProductService) BuyProduct(productId uint, user *userModel.User) error {
	if err := productService.ProductDao.BuyProduct(productId, user); err != nil {
		return fmt.Errorf("failed to finalize product purchase: %v", err)
	}
	return nil
}
