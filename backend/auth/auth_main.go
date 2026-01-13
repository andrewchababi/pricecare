package auth

import (
	"log"

	"github.com/andrewchababi/pricecare/backend/database"
	"github.com/andrewchababi/pricecare/backend/models"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func Login(username string, password string) (models.User, string) {
	user := database.GetUserByUsername(username)
	if user.UserType == models.UserTypeNone {
		log.Printf("Login attempt failed because username '%s' does not exist", username)
		return models.NullUser(), ""
	}

	if !isPasswordCorrect(password, user.HashedPassword) {
		log.Printf("Login attempt failed because password '%s' is invalid for user '%s'", password, username)
		return models.NullUser(), ""
	}

	loginToken := database.CreateLoginToken(user.Id)
	log.Printf("User '%s' logged in", user.Username)

	return user, loginToken
}

func Logout(loginToken string) {
	log.Println("Logout auth function called [+]")
	user := getUserFromLoginToken(loginToken)

	log.Println("user from login token success [+] " + user.Username)

	if user.UserType != models.UserTypeNone {
		log.Printf("User '%s' logged out", user.Username)
	}

	loginTokenId, err := bson.ObjectIDFromHex(loginToken)
	log.Println("login token Id" + loginToken)
	if err != nil {
		log.Printf("Failed to decode loginToken in database.DeleteLoginToken(): %v", err)
		return
	}

	database.DeleteLoginToken(loginTokenId)
}

func Register(username string, password string, role models.UserType) {
	hashedPassword := getHashedPassword(password)

	user := models.User{
		Id:             bson.NewObjectID(),
		Username:       username,
		HashedPassword: hashedPassword,
		UserType:       role,
	}

	database.CreateUser(user)
}
