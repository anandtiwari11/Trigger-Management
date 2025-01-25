package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/anandtiwari11/event-trigger/constants"
	"github.com/anandtiwari11/event-trigger/models"
	"github.com/gin-gonic/gin"
)

func (triggerController *TriggerController) ProcessTrigger(c *gin.Context) {
	var input models.Trigger
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if input.Type == constants.API {
		if input.Endpoint == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Trigger"})
			return
		}
		handleAPITrigger(triggerController, c, &input)
	} else if input.Type == constants.SCHEDULED {
		if input.ExecutionTime.IsZero() {
			input.ExecutionTime = time.Now()
		}
		handleScheduledTrigger(triggerController, c, &input)
	}
}

func handleScheduledTrigger(triggerController *TriggerController, c *gin.Context, input *models.Trigger) {
	istLocation, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		log.Printf("Error Finding Time Zone")
	}
	log.Printf("Received input payload: %s", string(input.Message))
	apiResponse := models.Event{
		Payload:   input.Payload,
		Response:  string(input.Message),
		State:     constants.ACTIVE,
		Timestamp: time.Now(),
	}
	input.ExecutionTime = input.ExecutionTime.In(istLocation)
	if err := triggerController.TriggerService.CreateNewTrigger(input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to store Trigger",
			"details": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "data": apiResponse})
}

func handleAPITrigger(triggerController *TriggerController, c *gin.Context, input *models.Trigger) {
	client := &http.Client{
		Timeout: 20 * time.Second,
	}
	log.Printf("Received input payload: %s", string(input.Payload))
	var payloadJSON interface{}
	if err := json.Unmarshal(input.Payload, &payloadJSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload format", "details": err.Error()})
		return
	}
	request, err := http.NewRequest(http.MethodPost, input.Endpoint, strings.NewReader(string(input.Payload)))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to create API request",
			"details": err.Error(),
		})
		return
	}
	request.Header.Set("Content-Type", "application/json")
	response, err := client.Do(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to call API",
			"details": err.Error(),
		})
		return
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read API response"})
		return
	}
	log.Printf("API Response: %s", string(body))

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "API call returned non-2xx status code",
			"details": string(body),
		})
		return
	}
	apiResponse := models.Event{
		Payload:   input.Payload,
		Response:  string(body),
		State:     constants.ACTIVE,
		Timestamp: time.Now(),
	}
	log.Printf("Attempting to store API response: %+v", apiResponse)
	if err := triggerController.TriggerService.CreateNewEvent(&apiResponse); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to store API response",
			"details": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "data": apiResponse})
}

func (triggerController *TriggerController) GetAllEvents(c *gin.Context) {
	events, err := triggerController.TriggerService.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : err})
	}
	c.JSON(http.StatusOK, gin.H{"events" : events})
}

func (triggerController *TriggerController) UpdateEvent(c *gin.Context) {
	var input models.Event
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	err := triggerController.TriggerService.UpdateEvent(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : err})
	}
	c.JSON(http.StatusOK, gin.H{"message" : "event updated"})
}

func (triggerController *TriggerController) DeleteEvent(c *gin.Context) {
	var input models.Event
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	err := triggerController.TriggerService.DeleteEvent(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : err})
	}
	c.JSON(http.StatusOK, gin.H{"message" : "event updated"})
}

func (triggerController *TriggerController) GetAllTriggers(c *gin.Context) {
	triggers, err := triggerController.TriggerService.GetAllTriggers()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : err})
	}
	c.JSON(http.StatusOK, gin.H{"triggers" : triggers})
}


func (triggerController *TriggerController) UpdateTrigger(c *gin.Context) {
	var input models.Trigger
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := triggerController.TriggerService.UpdateTrigger(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
}

func (triggerController *TriggerController) TriggerUpdate() {
	
}