package database

import (
	"context"
	"log"
	"time"

	"github.com/andrewchababi/pricecare/backend/models"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func GetUserByUsername(username string) models.User {
	var user models.User

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := usersCollection.FindOne(ctx, bson.M{"username": username}).Decode(&user)

	if err != nil {
		return models.NullUser()
	}

	return user
}

func GetUserById(userId bson.ObjectID) models.User {
	var user models.User

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := usersCollection.FindOne(ctx, bson.M{"_id": userId}).Decode(&user)

	if err != nil {
		return models.NullUser()
	}

	return user
}

func CreateUser(user models.User) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := usersCollection.InsertOne(ctx, user)

	if err != nil {
		log.Printf("Failed to insert user into database: %v", err)
	}
}
