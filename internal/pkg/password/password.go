// Package password provides helpers for securely hashing and verifying
// user passwords using bcrypt instead of storing/comparing them in plaintext.
package password

import "golang.org/x/crypto/bcrypt"

// Hash returns a bcrypt hash of the given plaintext password, suitable for
// storage in the database.
func Hash(plaintext string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(plaintext), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

// Matches performs a constant-time comparison between a bcrypt hash and a
// plaintext candidate password. It returns true only if they match.
func Matches(hashed, plaintext string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plaintext))
	return err == nil
}
