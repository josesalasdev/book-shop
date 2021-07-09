package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// Encrypt func.
func Encrypt(word string) string {
	secret := []byte(word)

	// Hashing the secret with the default cost of 10
	hashedSecret, err := bcrypt.GenerateFromPassword(secret, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hashedSecret)
}

// CompareHash func.
func CompareHash(wordHashed string, wordNotHashed string) error {
	return bcrypt.CompareHashAndPassword([]byte(wordHashed), []byte(wordNotHashed))
}
