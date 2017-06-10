package sessions

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/boltdb/bolt"
	jwt "github.com/dgrijalva/jwt-go"
)

// SessionHandler receives session-related functions, embedding the session
// database connection pool
type SessionHandler struct {
	db *bolt.DB
}

// Session is a user's session
type Session struct {
	ID         string    `json:"id"`
	UserID     uint64    `json:"user_id"`
	CreateDate time.Time `json:"create_date"`
}

var sess *SessionHandler

func init() {
	// Open session DB
	db, err := bolt.Open("boltdb/session.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	sess = &SessionHandler{db: db}

	// Create the session bucket if it doesn't already exist
	err = sess.db.Update(func(tx *bolt.Tx) error {
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

// GetSessionHandler returns the initialize instance of the SessionHandler
func GetSessionHandler() *SessionHandler {
	return sess
}

// GetUserID attempts to retrieve a User's ID form the Request's Session.
// Returns 0 if not found.
func (sess *SessionHandler) GetUserID(r *http.Request) uint64 {
	var session *Session

	// Get session cookie. Return nil if it does not exist.
	cookie, err := r.Cookie("sid")
	if err != nil {
		return 0
	}

	// Retrieve session from DB
	err = sess.db.View(func(tx *bolt.Tx) error {
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
		return 0
	}

	return session.UserID
}

// NewSession creates and stores a session cookie (sid)
func (sess *SessionHandler) NewSession(w http.ResponseWriter, userID uint64) *Session {
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

	err = sess.db.Update(func(tx *bolt.Tx) error {
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
