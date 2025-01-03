package services

import (
	"time"

	productModel "github.com/anandtiwari11/IEKart-go/products/models"
	daointerfaces "github.com/anandtiwari11/IEKart-go/user/daoInterfaces"
	userModel "github.com/anandtiwari11/IEKart-go/user/models"
	jwt "github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserDao daointerfaces.IUserDAO
}

func NewUserService(UserDao daointerfaces.IUserDAO) *UserService {
	return &UserService{
		UserDao: UserDao,
	}
}

func (userService *UserService) FindUserByUsername(username string) (*userModel.User, error) {
	return userService.UserDao.FindUserByUsername(username)
}
func (userService *UserService) FindUserByEmail(username string) (*userModel.User, error) {
	return userService.UserDao.FindUserByEmail(username)
}
func (userService *UserService) CreateUser(user *userModel.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return userService.UserDao.CreateUser(user)
}

func (userService *UserService) GenerateToken(login userModel.LoginUser) string {
	user, err := userService.FindUserByUsername(login.Username)
	if err != nil {
		return "user not found"
	}
	tokenString, err := userService.GenerateJWT(user.Email)
	if err != nil {
		return "failed to generate token"
	}
	return tokenString
}

func (userService *UserService) GenerateJWT(email string) (string, error) {
	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("secret-key"))
}

func (userService *UserService) GetAllProductOfTheUser(userId uint) (*[]productModel.Product, error) {
	return userService.UserDao.GetAllProductOfTheUser(userId)
}

func (userService *UserService) GetBookedProducts(userId uint) (*[]productModel.Product, error) {
	return userService.UserDao.GetBookedProducts(userId)
}