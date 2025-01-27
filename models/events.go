package models

import (
	"time"
)

type Event struct {
	ID            uint      `gorm:"primary_key" json:"id"`
	ExecutionTime time.Time `json:"execution_time"`
	Message       string    `json:"message"`
	State         string    `json:"state"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
