package routes

import (
	"github.com/anandtiwari11/event-trigger/controller"
	"github.com/gin-gonic/gin"
)

func RegisterTriggerRoutes(r *gin.Engine, triggerController *controller.TriggerController) {
	r.POST("/trigger", triggerController.ProcessTrigger)
	r.GET("/allTriggers", triggerController.GetAllTriggers)
	r.PUT("/updateTrigger", triggerController.UpdateTrigger)
}