package sessions

import (
	"fmt"
	"net/http"
	"time"

	"github.com/nikovacevic/commonwealth/models"
	uuid "github.com/satori/go.uuid"
)

// Session is a user's session
type Session struct {
	ID         uuid.UUID `json:"id"`
	UserID     uint      `json:"user_id"`
	CreateDate time.Time `json:"create_date"`
}

// DBUser defines User table
// UserID: User
var DBUser = map[uint]models.User{}

// SessionID: Session
var dbSession = map[string]Session{}

// GetUser attempts to find User's Session based on the session cookie in the
// Request. Returns the User if successful, nil if unsuccessful.
func GetUser(w http.ResponseWriter, r *http.Request) *models.User {
	var user models.User
	var session Session

	// Get session cookie. Return nil if it does not exist.
	cookie, err := r.Cookie("sid")
	if err != nil {
		return nil
	}

	// Retrieve Session from DB
	// TODO implement BoltDB
	session, ok := dbSession[cookie.Value]
	if !ok {
		return nil
	}

	// Retrieve User from Session
	// TODO implement BoltDB
	user, ok = DBUser[session.UserID]
	if !ok {
		panic(fmt.Errorf("Session %v exists for non-existent User ID %v", session.ID, session.UserID))
	}

	return &user
}

// NewSession creates and stores a session cookie (sid)
func NewSession(w http.ResponseWriter, userID uint) *Session {
	id := uuid.NewV4()
	cookie := &http.Cookie{
		Name:  "sid",
		Value: id.String(),
	}
	http.SetCookie(w, cookie)

	session := Session{
		ID:         id,
		UserID:     userID,
		CreateDate: time.Now(),
	}
	dbSession[id.String()] = session

	return &session
}
