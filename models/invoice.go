package models

import (
	"time"

	"github.com/nikovacevic/money"
)

// Invoice is an order
type Invoice struct {
	ID           uint64    `json:"id"`
	CreateDate   time.Time `json:"create_date"`
	CreateUserID uint64    `json:"create_user_id"`
	DueDate      time.Time `json:"due_date"`
	Total        money.USD `json:"total"`
}
