package auth

import (
	"log"

	"github.com/andrewchababi/pricecare/backend/config"
	"golang.org/x/crypto/bcrypt"
)

func GetHashedPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), config.AuthBcryptDifficultyFactor)
	if err != nil {
		log.Printf("Failed to hash password in auth.getHashedPassword(): %v", err)
		return ""
	}

	return string(bytes)
}

func isPasswordCorrect(password string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	return err == nil
}
