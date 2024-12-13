package routers

import (
	"github.com/anandtiwari11/IEKart-go/routers/middleware"
	"github.com/anandtiwari11/IEKart-go/user/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(engine *gin.Engine, userController *controllers.UserController) {
	authRoutes := engine.Group("/auth")
	authRoutes.POST("/signup", userController.Signup)
	authRoutes.POST("/login", userController.Login)
	authRoutes.GET("/getinfo", middleware.RequireAuth, userController.GetInfo)
	authRoutes.GET("/logout", middleware.RequireAuth, userController.Logout)
	authRoutes.GET("/userproducts/:id", middleware.RequireAuth, userController.GetAllProductOfTheUser)
	authRoutes.GET("/booked-products", middleware.RequireAuth, userController.GetBookedProducts)
}
