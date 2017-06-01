package models

import (
	"time"
)

// User is a user
type User struct {
	ID           uint      `json:"id"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Organization string    `json:"organization"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"password_hash"`
	CreateDate   time.Time `json:"create_date"`
}