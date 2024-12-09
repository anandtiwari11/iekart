package daoimpl

import (
	"fmt"

	"github.com/anandtiwari11/IEKart-go/initializers"
	productModel "github.com/anandtiwari11/IEKart-go/products/models"
	userModel "github.com/anandtiwari11/IEKart-go/user/models"
)

type UserDAOImpl struct{}

func NewUserDAOImpl() *UserDAOImpl {
    return &UserDAOImpl{}
}

func (dao *UserDAOImpl) FindUserByUsername(username string) (*userModel.User, error) {
    var user userModel.User
    result := initializers.DB.Where("username = ? and is_active = ?", username, true).First(&user)
    if result.Error != nil {
        return nil, result.Error
    }
    return &user, nil
}

func (dao *UserDAOImpl) FindUserByEmail(email string) (*userModel.User, error) {
    var user userModel.User
    result := initializers.DB.Where("email = ? and is_active = ?", email, true).First(&user)
    if result.Error != nil {
        return nil, result.Error
    }
    return &user, nil
}

func (dao *UserDAOImpl) CreateUser(user *userModel.User) error {
    if err := initializers.DB.Create(&user).Error; err != nil {
        return err
    }
    return nil
}

func (dao *UserDAOImpl) GetAllProductOfTheUser(userId uint) (*[]productModel.Product, error) {
    var products []productModel.Product
    if err := initializers.DB.Where("id = ?", userId).Find(&products).Error; err != nil {
        return nil, fmt.Errorf("failed to fetch products for user ID %d: %w", userId, err)
    }
    return &products, nil
}