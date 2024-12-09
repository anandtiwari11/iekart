package daointerfaces

import "github.com/anandtiwari11/IEKart-go/user/models"

type IUserDAO interface {
	FindUserByUsername(username string) (*models.User, error)
	FindUserByEmail(username string) (*models.User, error)
	CreateUser(user *models.User) error
}