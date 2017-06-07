package models

import (
	"time"
)

// User is a user
type User struct {
	ID           uint64    `json:"id"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Organization string    `json:"organization"`
	Email        string    `json:"email"`
	Phone        string    `json:"phone"`
	PasswordHash string    `json:"password_hash"`
	CreateDate   time.Time `json:"create_date"`
}
