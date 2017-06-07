package auth

import "golang.org/x/crypto/bcrypt"

// HashPassword hashes a password. Alias of bcrypt.GenerateFromPassword().
func HashPassword(password []byte) ([]byte, error) {
	return bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
}

// CheckPassword checks a given password against a hashed password, returning
// an error if the two do not match. Alias of bcrypt.CompareHashAndPassword().
func CheckPassword(hash, password []byte) error {
	return bcrypt.CompareHashAndPassword(hash, password)
}
