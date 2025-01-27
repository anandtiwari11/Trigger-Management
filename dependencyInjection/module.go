package dependencyinjection

import (
	"log"
	"os"

	"github.com/anandtiwari11/event-trigger/controller"
	jobs "github.com/anandtiwari11/event-trigger/cronJobs"
	"github.com/anandtiwari11/event-trigger/dao"
	daointerface "github.com/anandtiwari11/event-trigger/daoInterface"
	"github.com/anandtiwari11/event-trigger/service"
	serviceinterface "github.com/anandtiwari11/event-trigger/serviceInterface"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

var TriggerModule =  fx.Options(
	fx.Provide(
		fx.Annotate(
			dao.NewTriggerDaoImpl,
			fx.As(new(daointerface.ITriggerDao)),
		),
	),
	fx.Provide(
		fx.Annotate(
			service.NewTriggerService,
			fx.As(new(serviceinterface.IServiceInterface)),
		),
	),
	fx.Provide(controller.NewTriggerController),
)

func bootstrap(router *gin.Engine) {
	log.Println("Setting up the Gin server on port :8080...")
	go func() {
		port := os.Getenv("PORT")
		if port == "" {
			port = "8080"
		}
		if err := router.Run("0.0.0.0:8080"); err != nil {
			log.Fatalf("Failed to run server: %v", err)
		}
	}()
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Authorization, X-Requested-With, Accept, Origin")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, PATCH, DELETE, OPTIONS")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func EventCronJob() {
	triggerDAO := &dao.TriggerDaoImpl{}
	eventLifecycleJob := &jobs.EventLifecycleJob{
		TriggerDAO: triggerDAO,
	}
	go eventLifecycleJob.Run()
}