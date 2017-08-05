package models

import (
	"time"

	"github.com/nikovacevic/money"
)

// ProductOption is an option for a Product
type ProductOption struct {
	ID           uint64    `json:"id"`
	Cost         money.USD `json:"cost"`
	CreateDate   time.Time `json:"create_date"`
	CreateUserID uint64    `json:"create_user_id"`
	Description  string    `json:"description"`
	IsActive     bool      `json:"is_active"`
	Price        money.USD `json:"price"`
	ProductID    uint64    `json:"product_id"`
}
