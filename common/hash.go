package common

import (
	"golang.org/x/crypto/bcrypt"
)

// Hashes string password with default cost (may be changed for security).
// Function panics if password cannot be generated.
func MustHashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hash)
}
