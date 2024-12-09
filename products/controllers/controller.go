package controllers

import servicesinterfaces "github.com/anandtiwari11/IEKart-go/products/servicesInterfaces"



type ProductController struct {
	ProductService servicesinterfaces.IProductIntf
}

func NewProductController (productService servicesinterfaces.IProductIntf) *ProductController {
	return &ProductController{
		ProductService: productService,
	}
}