package controller

import serviceinterface "github.com/anandtiwari11/event-trigger/serviceInterface"

type TriggerController struct {
	TriggerService serviceinterface.IServiceInterface
}

func NewTriggerController(triggerService serviceinterface.IServiceInterface) *TriggerController {
	return &TriggerController{
		TriggerService: triggerService,
	}
}