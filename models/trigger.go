package models

import (
	"encoding/json"
	"time"
)

type Trigger struct {
	ID            uint            `gorm:"primary_key" json:"id"`
	Name          string          `gorm:"not null" json:"name"`               // Descriptive name for the trigger
	Type          string          `gorm:"not null" json:"type"`               // Type of trigger: "scheduled" or "api"
	Interval      uint            `json:"interval,omitempty"`                 // Interval for recurring triggers (e.g., "10m", "1h")
	ExecutionTime time.Time       `json:"execution_time,omitempty"`           // Exact time for non-recurring triggers
	Endpoint      string          `json:"endpoint,omitempty"`                 // Endpoint for API triggers
	Message       string          `json:"message,omitempty"`                  // message for scheduled type
	Payload       json.RawMessage `gorm:"type:json" json:"payload,omitempty"` // JSON payload for API triggers
	IsRecurring   bool            `json:"is_recurring"`                       // Whether the trigger is recurring
	CreatedAt     time.Time       `json:"created_at"`                         // Timestamp for when the trigger was created
	UpdatedAt     time.Time       `json:"updated_at"`                         // Timestamp for when the trigger was last updated
}
