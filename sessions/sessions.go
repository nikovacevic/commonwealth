package sessions

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/boltdb/bolt"
	"github.com/nikovacevic/commonwealth/models"
	uuid "github.com/satori/go.uuid"
)

// Session is a user's session
type Session struct {
	ID         uuid.UUID `json:"id"`
	UserID     uint64    `json:"user_id"`
	CreateDate time.Time `json:"create_date"`
}

func init() {
	// Open session DB
	db, err := bolt.Open("boltdb/session.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create the session bucket if it doesn't already exist
	err = db.Update(func(tx *bolt.Tx) error {
		_, err = tx.CreateBucketIfNotExists([]byte("session"))
		if err != nil {
			return err
		}
		_, err = tx.CreateBucketIfNotExists([]byte("user"))
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}

// GetUser attempts to find User's Session based on the session cookie in the
// Request. Returns the User if successful, nil if unsuccessful.
func GetUser(w http.ResponseWriter, r *http.Request, db *bolt.DB) *models.User {
	var user *models.User
	var session *Session

	// Get session cookie. Return nil if it does not exist.
	cookie, err := r.Cookie("sid")
	if err != nil {
		return nil
	}

	// Retrieve session from DB
	err = db.View(func(tx *bolt.Tx) error {
		sessionBkt := tx.Bucket([]byte("session"))
		data := sessionBkt.Get([]byte(cookie.Value))
		if data == nil {
			return nil
		}
		session = &Session{}
		err = json.Unmarshal(data, session)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	if session == nil {
		return nil
	}

	// Retrieve User from Session
	// Retrieve session from DB
	err = db.View(func(tx *bolt.Tx) error {
		userBkt := tx.Bucket([]byte("user"))
		b := make([]byte, 8)
		binary.LittleEndian.PutUint64(b, session.UserID)
		data := userBkt.Get(b)
		if data == nil {
			return nil
		}
		user = &models.User{}
		err = json.Unmarshal(data, user)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	if user == nil {
		log.Fatal(fmt.Errorf("Session %v exists for non-existent User ID %v", session.ID, session.UserID))
	}

	return user
}

// NewSession creates and stores a session cookie (sid)
func NewSession(w http.ResponseWriter, userID uint64, db *bolt.DB) *Session {
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

	err := db.Update(func(tx *bolt.Tx) error {
		sessionBkt := tx.Bucket([]byte("session"))
		data, er := json.Marshal(session)
		if er != nil {
			return er
		}
		er = sessionBkt.Put([]byte(session.ID.String()), data)
		if er != nil {
			return er
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	return &session
}
