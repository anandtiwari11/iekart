package serviceinterfaces

import "github.com/anandtiwari11/IEKart-go/user/models"

type IUserServiceInterface interface {
	FindUserByUsername(username string) (*models.User, error)
	FindUserByEmail(username string) (*models.User, error)
	CreateUser(user *models.User) error
	GenerateToken(user models.LoginUser) string
	GenerateJWT(email string) (string, error)
}
