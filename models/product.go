package models

import (
	"time"

	"github.com/nikovacevic/money"
)

// Product is a product
type Product struct {
	ID           uint64    `json:"id"`
	Cost         money.USD `json:"cost"`
	CreateDate   time.Time `json:"create_date"`
	CreateUserID uint64    `json:"create_user_id"`
	Description  string    `json:"description"`
	IsActive     bool      `json:"is_active"`
	Name         string    `json:"name"`
	Price        money.USD `json:"price"`
}
