package daoimpl

import (
	"fmt"

	"github.com/anandtiwari11/IEKart-go/initializers"
	productModel "github.com/anandtiwari11/IEKart-go/products/models"
	userModel "github.com/anandtiwari11/IEKart-go/user/models"
	"gorm.io/gorm"
)

type daoImpl struct{}

func NewProductDaoImpl() *daoImpl {
	return &daoImpl{}
}

func (dao *daoImpl) CreateProduct(product *productModel.Product, user *userModel.User) error {
	return initializers.DB.Transaction(func(tx *gorm.DB) error {
		product.SellerID = user.ID
		if err := tx.Create(product).Error; err != nil {
			return fmt.Errorf("failed to create product: %v", err)
		}
		return nil
	})
}

func (dao *daoImpl) GetProductDetailsByProductId(productId uint) (*productModel.Product, error) {
	var product productModel.Product
	if err := initializers.DB.Where("id = ?", productId).First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("product with id %s not found", string(productId))
		}
		return nil, err
	}
	return &product, nil
}

func (dao *daoImpl) DeleteProductByID(productId uint) error {
	return initializers.DB.Transaction(func(tx *gorm.DB) error {
		var product productModel.Product
		if err := tx.First(&product, "id = ?", productId).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return fmt.Errorf("product with ID %d not found", productId)
			}
			return fmt.Errorf("failed to find product: %v", err)
		}
		if err := tx.Delete(&product).Error; err != nil {
			return fmt.Errorf("failed to delete product: %v", err)
		}
		return nil
	})
}

func (dao *daoImpl) BookProduct(productId uint, user *userModel.User) error {
	return initializers.DB.Transaction(func(tx *gorm.DB) error {
		var product productModel.Product
		if err := tx.First(&product, "id = ?", productId).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return fmt.Errorf("product with ID %d not found", productId)
			}
			return fmt.Errorf("failed to retrieve product: %v", err)
		}
		if product.BookedByID != nil {
			return fmt.Errorf("product with ID %d is already booked", productId)
		}
		if product.BuyerID != nil {
			return fmt.Errorf("product with ID %d has already been sold", productId)
		}
		product.BookedByID = &user.ID
		if err := tx.Save(&product).Error; err != nil {
			return fmt.Errorf("failed to book product: %v", err)
		}
		return nil
	})
}

func (dao *daoImpl) BuyProduct(productId uint, user *userModel.User) error {
	return initializers.DB.Transaction(func(tx *gorm.DB) error {
		var product productModel.Product
		if err := tx.First(&product, "id = ?", productId).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return fmt.Errorf("product with ID %d not found", productId)
			}
			return fmt.Errorf("failed to retrieve product: %v", err)
		}
		if product.BookedByID == nil || *product.BookedByID != user.ID {
			return fmt.Errorf("you can only buy a product you have booked")
		}
		if product.BuyerID != nil {
			return fmt.Errorf("product with ID %d has already been purchased", productId)
		}
		product.BuyerID = &user.ID
		if err := tx.Delete(&product).Error; err != nil {
			return fmt.Errorf("failed to finalize purchase and delete product: %v", err)
		}
		return nil
	})
}