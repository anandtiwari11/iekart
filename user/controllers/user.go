package controllers

import (
	"net/http"
	"strconv"

	userModel "github.com/anandtiwari11/IEKart-go/user/models"
	"github.com/gin-gonic/gin"
)

func (userController *UserController) Signup(c *gin.Context) {
    var input userModel.User
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
        return
    }

    if user, _ := userController.UserService.FindUserByUsername(input.Username); user != nil {
        c.JSON(http.StatusConflict, gin.H{"error": "user already exists"})
        return
    }

    if user, _ := userController.UserService.FindUserByEmail(input.Email); user != nil {
        c.JSON(http.StatusConflict, gin.H{"error": "user with this email already exists"})
        return
    }

    if err := userController.UserService.CreateUser(&input); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create a new user"})
        return
    }
    c.JSON(http.StatusCreated, gin.H{"message": input})
}


func (userController *UserController) Login(c *gin.Context) {
	var input userModel.LoginUser
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : "invalid request"})
		return
	}
    tokenString := userController.UserService.GenerateToken(input)
    if tokenString == "user not found" {
        c.JSON(http.StatusNotFound, gin.H{"error" : "user not found"})
        return
    }
    c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600 * 72, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"message" : "Login Successfull",
	})
}

func (userController *UserController) Logout(c *gin.Context) {
    c.SetCookie("Authorization", "", -1, "", "", false, true)
    c.JSON(http.StatusOK, gin.H{"message": "successfully logged out"})
}

func (userController *UserController) GetInfo(c *gin.Context) {
    user, exists := c.Get("user")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized access"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"user": user})
}

func (userController *UserController) GetAllProductOfTheUser(c *gin.Context) {
    userIdStr := c.Param("id")
    userId, err := strconv.ParseUint(userIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}
    product, err := userController.UserService.GetAllProductOfTheUser(uint(userId))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error" : err.Error()})
    }
    c.JSON(http.StatusOK, gin.H{"message" : product})
}

func (userController *UserController) GetBookedProducts(c *gin.Context) {
    userIdStr := c.Param("id")
    userId, err := strconv.ParseUint(userIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}
    product, err := userController.UserService.GetBookedProducts(uint(userId))
    c.JSON(http.StatusOK, gin.H{"message" : product})
}
