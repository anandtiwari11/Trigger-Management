package dependencyinjection

import (
	"github.com/anandtiwari11/event-trigger/routes"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)


func LoadDependencies() *fx.App {
	return fx.New(
		fx.Provide(func() *gin.Engine {
			return gin.Default()
		}),
		TriggerModule,
		fx.Invoke(routes.RegisterTriggerRoutes),
		fx.Invoke(routes.RegisterEventRoutes),
		fx.Invoke(bootstrap),
	)
}