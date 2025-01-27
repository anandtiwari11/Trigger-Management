package controller

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/anandtiwari11/event-trigger/constants"
	"github.com/anandtiwari11/event-trigger/models"
	"github.com/gin-gonic/gin"
)

func (triggerController *TriggerController) ProcessTrigger(c *gin.Context) {
	var input models.Trigger
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{constants.ERROR : constants.INVALID_INPUT})
		return
	}

	if input.Type == constants.API {
		if input.Endpoint == "" {
			c.JSON(http.StatusBadRequest, gin.H{constants.ERROR: constants.INVALID_INPUT})
			return
		}
		handleAPITrigger(triggerController, c, &input)
	} else if input.Type == constants.SCHEDULED {
		if input.ExecutionTime.IsZero() {
			input.ExecutionTime = time.Now()
		}
		c.JSON(http.StatusOK, input)
		err := handleScheduledTrigger(triggerController, &input)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{constants.ERROR: err.Error()})
			return
		}
	}
}

func handleScheduledTrigger(triggerController *TriggerController, input *models.Trigger) error {
	if err := triggerController.TriggerService.CreateNewTrigger(input); err != nil {
		return fmt.Errorf(constants.FAILED_TO_PROCESS_SCHEDULED_TRIGGER, err)
	}
    if err := triggerController.TriggerService.ProcessScheduledTrigger(input); err != nil {
        return fmt.Errorf(constants.FAILED_TO_PROCESS_SCHEDULED_TRIGGER, err)
    }
    if input.IsRecurring {
        input.ExecutionTime = input.ExecutionTime.Add(time.Duration(input.Interval) * time.Minute)
        if err := triggerController.TriggerService.UpdateTrigger(input); err != nil {
            return fmt.Errorf(constants.FAILED_TO_UPDATE_SCHEDULED_TRIGGER, err)
        }
    }
    
    return nil
}

func handleAPITrigger(triggerController *TriggerController, c *gin.Context, input *models.Trigger) {
	triggerController.TriggerService.CreateNewTrigger(input)
	log.Printf("Received input payload: %s", string(input.Payload))
	if err := triggerController.TriggerService.ProcessAPITrigger(input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{constants.ERROR: err})
	}
	c.JSON(http.StatusOK, gin.H{constants.STATUS: constants.SUCCESS})
}

func (triggerController *TriggerController) GetAllEvents(c *gin.Context) {
	events, err := triggerController.TriggerService.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{constants.ERROR: err})
	}
	c.JSON(http.StatusOK, gin.H{constants.EVENTS: events})
}

func (triggerController *TriggerController) UpdateEvent(c *gin.Context) {
	var input models.Event
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{constants.ERROR: constants.INVALID_INPUT})
		return
	}
	err := triggerController.TriggerService.UpdateEvent(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{constants.ERROR: err})
	}
	c.JSON(http.StatusOK, gin.H{constants.MESSAGE: constants.EVENT_UPDATED})
}

func (triggerController *TriggerController) DeleteEvent(c *gin.Context) {
	var input models.Event
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{constants.ERROR: constants.INVALID_INPUT})
		return
	}
	err := triggerController.TriggerService.DeleteEvent(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{constants.ERROR: err})
	}
	c.JSON(http.StatusOK, gin.H{constants.MESSAGE: constants.EVENT_UPDATED})
}

func (triggerController *TriggerController) GetAllTriggers(c *gin.Context) {
	triggers, err := triggerController.TriggerService.GetAllTriggers()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{constants.ERROR: err})
	}
	c.JSON(http.StatusOK, gin.H{constants.TRIGGERS: triggers})
}

func (triggerController *TriggerController) UpdateTrigger(c *gin.Context) {
	var input models.Trigger
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{constants.ERROR: constants.INVALID_INPUT})
		return
	}
	if err := triggerController.TriggerService.UpdateTrigger(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{constants.ERROR: constants.INVALID_INPUT})
		return
	}
}


func (triggerController *TriggerController) DeleteTrigger(c *gin.Context) {
	var input models.Trigger
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{constants.ERROR: constants.INVALID_INPUT})
		return
	}
	err := triggerController.TriggerService.DeleteTrigger(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{constants.ERROR: err})
	}
	c.JSON(http.StatusOK, gin.H{constants.MESSAGE: constants.EVENT_UPDATED})
}