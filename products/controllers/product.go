package controllers

import (
	"net/http"
	"strconv"

	productModel "github.com/anandtiwari11/IEKart-go/products/models"
	userModel "github.com/anandtiwari11/IEKart-go/user/models"
	"github.com/gin-gonic/gin"
)

func (productController *ProductController) CreateProduct(c *gin.Context) {
	var product productModel.ProductReq
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, exist := c.Get("user")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized user"})
		return
	}
	authenticatedUser, ok := user.(userModel.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user information"})
		return
	}
	if err := productController.ProductService.CreateProduct(&product, &authenticatedUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create product"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "product created succesfully"})
}

func (productController *ProductController) GetInfo(c *gin.Context) {
	productIdStr := c.Param("id")
	productId, err := strconv.ParseUint(productIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product ID"})
		return
	}
	product, err := productController.ProductService.GetProductDetailsByProductId(uint(productId))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"product details": product})
}

func (productController *ProductController) DeleteProduct(c *gin.Context) {
	productIdStr := c.Param("id")
	productId, err := strconv.ParseUint(productIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	if err := productController.ProductService.DeleteProductByID(uint(productId)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "failed to delete product1"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "product deleted succesfully"})
}

func (productController *ProductController) BookProduct(c *gin.Context) {
	productIdParam := c.Param("id")
	productId, err := strconv.ParseUint(productIdParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product ID"})
		return
	}
	user, exist := c.Get("user")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized user"})
		return
	}
	authenticatedUser, ok := user.(userModel.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve user information"})
		return
	}
	if err := productController.ProductService.BookProduct(uint(productId), &authenticatedUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "product booked successfully"})
}

func (productController *ProductController) BuyProduct(c *gin.Context) {
	productIdParam := c.Param("id")
	productId, err := strconv.ParseUint(productIdParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product ID"})
		return
	}
	user, exist := c.Get("user")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized user"})
		return
	}
	authenticatedUser, ok := user.(userModel.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve user information"})
		return
	}
	if err := productController.ProductService.BuyProduct(uint(productId), &authenticatedUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "product purchased successfully"})
}