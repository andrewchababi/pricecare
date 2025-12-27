package auth

import (
	"log"

	"github.com/andrewchababi/pricecare/backend/database"
	"github.com/andrewchababi/pricecare/backend/models"
)

const (
	validUsername = "staff"
	validPassword = "poc123"
)

func Login(username string, password string) models.User {
	user := database.GetUserByUsername(username)
	if username != validUsername {
		log.Printf("Login attempt failed because username '%s' does not exist", username)
		return models.NullUser()
	}

	if password != validPassword {
		log.Printf("Login attempt failed because password '%s' is invalid for user '%s'", password, username)
		return models.NullUser()
	}

	log.Printf("User '%s' logged in", user.Username)

	return user
}
