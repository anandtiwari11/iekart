package dependencyinjection

import (
	"github.com/anandtiwari11/IEKart-go/routers"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func LoadDependencies() *fx.App {
	return fx.New(
		fx.Provide(func() *gin.Engine {
			return gin.Default()
		}),
		ProductModule,
		UserModule,
		fx.Invoke(routers.RegisterUserRoutes),
		fx.Invoke(routers.RegisterProductRoutes),
		fx.Invoke(bootstrap),
	)
}
