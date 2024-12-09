package daoimpl

import (
	"github.com/anandtiwari11/IEKart-go/initializers"
	"github.com/anandtiwari11/IEKart-go/user/models"
)

type UserDAOImpl struct{}

func NewUserDAOImpl() *UserDAOImpl {
    return &UserDAOImpl{}
}

func (dao *UserDAOImpl) FindUserByUsername(username string) (*models.User, error) {
    var user models.User
    result := initializers.DB.Where("username = ? and is_active = ?", username, true).First(&user)
    if result.Error != nil {
        return nil, result.Error
    }
    return &user, nil
}

func (dao *UserDAOImpl) FindUserByEmail(email string) (*models.User, error) {
    var user models.User
    result := initializers.DB.Where("email = ? and is_active = ?", email, true).First(&user)
    if result.Error != nil {
        return nil, result.Error
    }
    return &user, nil
}

func (dao *UserDAOImpl) CreateUser(user *models.User) error {
    if err := initializers.DB.Create(&user).Error; err != nil {
        return err
    }
    return nil
}
