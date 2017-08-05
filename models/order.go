package models

import (
	"time"
)

// Order is an order
type Order struct {
	ID           uint64    `json:"id"`
	InvoiceID    uint64    `json:"invoice_id"`
	CreateDate   time.Time `json:"create_date"`
	CreateUserID uint64    `json:"create_user_id"`
	UserID       uint64    `json:"user_id"`
}
