package services

import (
	"context"
	"database/sql"
	"log"

	"github.com/nikovacevic/commonwealth/models"
)

// UserService is ...
type UserService struct {
	db *sql.DB
}

// UserNotFound is ...
type UserNotFound error

// NewUser creates and returns a new instance of the UserService
func NewUser(db *sql.DB) *UserService {
	return &UserService{db}
}

// ByEmail ...
func (us *UserService) ByEmail(email string) (*models.User, error) {
	return nil, *new(UserNotFound)
}

// ByID ...
func (us *UserService) ByID(id uint64) (*models.User, error) {
	user := &models.User{}
	ctx := context.Background()
	row := us.db.QueryRowContext(
		ctx,
		`SELECT   u.first_name,
		          u.last_name,
			  u.email,
			  u.phone,
			  u.organization
		FROM      users AS u
		WHERE     u.id = $1;`,
		id,
	)
	err := row.Scan(
		&(user.FirstName),
		&(user.LastName),
		&(user.Email),
		&(user.Phone),
		&(user.Organization),
	)
	if err != nil {
		return nil, *new(UserNotFound)
	}

	return user, nil
}

// Create ...
func (us *UserService) Create(user *models.User) (*models.User, error) {
	// TODO validate

	ctx := context.Background()
	stmt, err := us.db.PrepareContext(ctx, "INSERT INTO users (first_name, last_name, email, phone, organization, password_hash) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id")
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.ExecContext(ctx, user.FirstName, user.LastName, user.Email, user.Phone, user.Organization, user.PasswordHash)
	if err != nil {
		log.Fatal(err)
	}

	// TODO scan INSERTED ID into user.ID

	return user, nil
}
