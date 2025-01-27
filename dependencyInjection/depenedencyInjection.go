package dependencyinjection

import (
	"github.com/anandtiwari11/event-trigger/initializers"
	"github.com/anandtiwari11/event-trigger/routes"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)


func LoadDependencies() *fx.App {
	return fx.New(
		fx.Provide(func() *gin.Engine {
			router:= gin.Default()
			router.Use(CORSMiddleware())
			return router
		}),
		TriggerModule,
		fx.Invoke(EventCronJob),
		fx.Invoke(routes.RegisterTriggerRoutes),
		fx.Invoke(routes.RegisterEventRoutes),
		fx.Invoke(bootstrap),
		fx.Invoke(initializers.ConnectDB),
		fx.Invoke(initializers.InitScheduler),
	)
}