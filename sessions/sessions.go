package sessions

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/boltdb/bolt"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/nikovacevic/commonwealth/models"
)

// Session is a user's session
type Session struct {
	ID         string    `json:"id"`
	UserID     uint64    `json:"user_id"`
	CreateDate time.Time `json:"create_date"`
}

func init() {
	// Open session DB
	bdb, err := bolt.Open("boltdb/session.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer bdb.Close()

	// Create the session bucket if it doesn't already exist
	err = bdb.Update(func(tx *bolt.Tx) error {
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
func GetUser(w http.ResponseWriter, r *http.Request, bdb *bolt.DB, db *sql.DB) *models.User {
	var user *models.User
	var session *Session

	// Get session cookie. Return nil if it does not exist.
	cookie, err := r.Cookie("sid")
	if err != nil {
		return nil
	}

	// Retrieve session from DB
	err = bdb.View(func(tx *bolt.Tx) error {
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

	ctx := context.Background()
	user = &models.User{}
	// Retrieve User from Session
	row := db.QueryRowContext(ctx, "SELECT u.id, u.first_name, u.last_name, u.email, u.phone, u.organization, u.create_date FROM users AS u WHERE id = $1;", session.UserID)
	err = row.Scan(&(user.ID), &(user.FirstName), &(user.LastName), &(user.Email), &(user.Phone), &(user.Organization), &(user.CreateDate))
	if err != nil {
		log.Fatal(err)
	}

	if user == nil {
		log.Fatal(fmt.Errorf("Session %v exists for non-existent User ID %v", session.ID, session.UserID))
	}

	return user
}

// NewSession creates and stores a session cookie (sid)
func NewSession(w http.ResponseWriter, userID uint64, bdb *bolt.DB) *Session {
	secret := []byte("7D72E7C2E630EE763C2E7A1AEB3F2035A0227E8C66C1F3EFC64")
	token := jwt.New(jwt.SigningMethodHS256)
	signedToken, err := token.SignedString(secret)
	if err != nil {
		log.Fatal(err)
	}

	cookie := &http.Cookie{
		Name:    "sid",
		Value:   signedToken,
		Expires: time.Now().Add(30 * time.Minute),
	}
	http.SetCookie(w, cookie)

	session := Session{
		ID:         signedToken,
		UserID:     userID,
		CreateDate: time.Now(),
	}

	err = bdb.Update(func(tx *bolt.Tx) error {
		sessionBkt := tx.Bucket([]byte("session"))
		data, er := json.Marshal(session)
		if er != nil {
			return er
		}
		er = sessionBkt.Put([]byte(signedToken), data)
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
