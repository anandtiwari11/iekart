package main

import (
	"log"

	"github.com/anandtiwari11/IEKart-go/initializers"
	"github.com/anandtiwari11/IEKart-go/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	initializers.ConnectDB()

	router := gin.Default()
	router.Use(func(ctx *gin.Context) {
		ctx.Set("library", initializers.DB)
		ctx.Next()
	})
	routers.AddUserRoutes(router)
	log.Fatal(router.Run(":8080"))
}