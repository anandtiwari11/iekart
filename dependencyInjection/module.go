package dependencyinjection

import (
	"log"

	productControllers "github.com/anandtiwari11/IEKart-go/products/controllers"
	productDaoImpl "github.com/anandtiwari11/IEKart-go/products/daoImpl"
	productDaointerface "github.com/anandtiwari11/IEKart-go/products/daoInterfaces"
	productService "github.com/anandtiwari11/IEKart-go/products/services"
	productServiceinterface "github.com/anandtiwari11/IEKart-go/products/servicesInterfaces"
	userController "github.com/anandtiwari11/IEKart-go/user/controllers"
	userDaoimpl "github.com/anandtiwari11/IEKart-go/user/daoImpl"
	userDaointerfaces "github.com/anandtiwari11/IEKart-go/user/daoInterfaces"
	userServiceinterfaces "github.com/anandtiwari11/IEKart-go/user/serviceInterfaces"
	userServices "github.com/anandtiwari11/IEKart-go/user/services"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

var ProductModule = fx.Options(
	fx.Provide(
		fx.Annotate(
			productDaoImpl.NewProductDaoImpl,
			fx.As(new(productDaointerface.IProductDao)),
		),
	),
	fx.Provide(
		fx.Annotate(
			productService.NewProductService,
			fx.As(new(productServiceinterface.IProductIntf)),
		),
	),
	fx.Provide(productControllers.NewProductController),
)

var UserModule = fx.Options(
	fx.Provide(
		fx.Annotate(
			userDaoimpl.NewUserDAOImpl,
			fx.As(new(userDaointerfaces.IUserDAO)),
		),
	),
	fx.Provide(
		fx.Annotate(
			userServices.NewUserService,
			fx.As(new(userServiceinterfaces.IUserServiceInterface)),
		),
	),
	fx.Provide(userController.NewUserController),
)

func bootstrap(router *gin.Engine) {
	log.Println("Setting up the Gin server on port :8080...")
	go func() {
		if err := router.Run("0.0.0.0:8080"); err != nil {
			log.Fatalf("Failed to run server: %v", err)
		}
	}()
}