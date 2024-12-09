package routers

import (
	"github.com/anandtiwari11/IEKart-go/routers/middleware"
	"github.com/anandtiwari11/IEKart-go/user/controllers"
	daoimpl "github.com/anandtiwari11/IEKart-go/user/daoImpl"
	"github.com/anandtiwari11/IEKart-go/user/services"
	"github.com/gin-gonic/gin"
)

func AddUserRoutes(router *gin.Engine) {
	userDAO := daoimpl.NewUserDAOImpl()                 
	userService := services.NewUserService(userDAO)  
	userController := controllers.NewUserController(userService)
	auth := router.Group("/auth")
	{
		auth.POST("/signup", userController.Signup)
		auth.POST("/login", userController.Login)
		auth.GET("/getinfo", middleware.RequireAuth, userController.GetInfo)
		auth.GET("/logout", middleware.RequireAuth, userController.Logout)
	}
}
