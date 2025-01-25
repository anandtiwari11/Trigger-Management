package routes

import (
	"github.com/anandtiwari11/event-trigger/controller"
	"github.com/gin-gonic/gin"
)

func RegisterEventRoutes(r *gin.Engine, triggerController *controller.TriggerController) {
	r.GET("/getEvents", triggerController.GetAllEvents)
	r.PUT("/updateEvents",triggerController.UpdateEvent)
	r.DELETE("/deleteEvent", triggerController.DeleteEvent)
}