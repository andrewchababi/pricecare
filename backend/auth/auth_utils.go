package auth

import (
	"log"

	"github.com/andrewchababi/pricecare/backend/database"
	"github.com/andrewchababi/pricecare/backend/models"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/v2/bson"
	"golang.org/x/crypto/bcrypt"
)

const BCRYPT_DIFFICULTY_FACTOR = 10

func getHashedPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), BCRYPT_DIFFICULTY_FACTOR)
	if err != nil {
		log.Printf("Failed to hash password: %v", err)
		return ""
	}
	return string(bytes)
}

func isPasswordCorrect(password string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func generateLoginToken() string {
	id := uuid.New()
	return id.String()
}

func getUserFromLoginToken(loginToken string) models.User {
	if loginToken == "" {
		return models.NullUser()
	}

	userId := database.GetUserIdFromLoginToken(loginToken)

	if userId == bson.NilObjectID {
		return models.NullUser()
	}

	user := database.GetUserById(userId)

	return user
}
