package routers

import (
	"github.com/anandtiwari11/IEKart-go/products/controllers"
	"github.com/anandtiwari11/IEKart-go/routers/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterProductRoutes(engine *gin.Engine, productController *controllers.ProductController) {
	productRoutes := engine.Group("/product")
	productRoutes.POST("/add", middleware.RequireAuth, productController.CreateProduct)
	productRoutes.GET("/info/:id", middleware.RequireAuth, productController.GetInfo)
	productRoutes.DELETE("/delete/:id", middleware.RequireAuth, productController.DeleteProduct)
	productRoutes.PATCH("/book/:id", middleware.RequireAuth, productController.BookProduct)
	productRoutes.DELETE("/buy/:id", middleware.RequireAuth, productController.BuyProduct)
}
