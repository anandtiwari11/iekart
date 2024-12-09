package routers

import (
	"github.com/anandtiwari11/IEKart-go/products/controllers"
	daoimpl "github.com/anandtiwari11/IEKart-go/products/daoImpl"
	"github.com/anandtiwari11/IEKart-go/products/services"
	"github.com/anandtiwari11/IEKart-go/routers/middleware"
	"github.com/gin-gonic/gin"
)

func AddProductRoutes(router *gin.Engine) {
	productDAO := daoimpl.NewProductDaoImpl()
	productService := services.NewProductService(productDAO)
	productController := controllers.NewProductController(productService)
	product := router.Group("/product")
	{
		product.POST("/add", middleware.RequireAuth, productController.CreateProduct)
		product.GET("/info/:id", middleware.RequireAuth, productController.GetInfo)
		product.DELETE("/delete/:id", middleware.RequireAuth, productController.DeleteProduct)
		product.PATCH("/book/:id", middleware.RequireAuth, productController.BookProduct)
		product.DELETE("/buy/:id", middleware.RequireAuth, productController.BuyProduct)
	}
}