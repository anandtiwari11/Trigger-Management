package models

import (
	"encoding/json"
	"time"
)

type Event struct {
	ID        uint            `gorm:"primaryKey;autoIncrement" json:"id"`                        // Unique ID for the event
	Payload   json.RawMessage `gorm:"type:json" json:"payload,omitempty"`                        // JSON payload for API-based triggers
	Response  string          `gorm:"type:json" json:"response,omitempty"`                       // The API response for API-based triggers
	State     string          `gorm:"type:varchar(50)" json:"state"`                             // Event state: "active", "archived", "deleted"
	Timestamp time.Time       `gorm:"type:timestamp;default:current_timestamp" json:"timestamp"` // Timestamp when the event was triggered
}
