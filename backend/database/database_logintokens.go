package database

import (
	"context"
	"errors"
	"log"

	"github.com/andrewchababi/pricecare/backend/config"
	"github.com/andrewchababi/pricecare/backend/models"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func GetLoginTokens() []models.LoginToken {
	ctx, cancel := context.WithTimeout(context.Background(), config.DatabaseTimeoutDuration)
	defer cancel()

	cursor, err := loginTokensCollection.Find(ctx, bson.D{})
	if err != nil {
		log.Printf("Failed to get login tokens in database.GetLoginTokens(): %v", err)
		return []models.LoginToken{}
	}
	defer cursor.Close(ctx)

	var loginTokens []models.LoginToken
	err = cursor.All(ctx, &loginTokens)
	if err != nil {
		log.Printf("Failed to get login tokens in database.GetLoginTokens(): %v", err)
		return []models.LoginToken{}
	}

	return loginTokens
}

func GetUserIdFromLoginToken(loginToken string) bson.ObjectID {
	loginTokenId, err := bson.ObjectIDFromHex(loginToken)
	if err != nil {
		return bson.NilObjectID
	}

	ctx, cancel := context.WithTimeout(context.Background(), config.DatabaseTimeoutDuration)
	defer cancel()

	var loginTokenObject models.LoginToken
	err = loginTokensCollection.FindOne(ctx, bson.M{"_id": loginTokenId}).Decode(&loginTokenObject)
	if err != nil {
		return bson.NilObjectID
	}

	return loginTokenObject.UserId
}

func CreateLoginToken(userId bson.ObjectID) string {
	ctx, cancel := context.WithTimeout(context.Background(), config.DatabaseTimeoutDuration)
	defer cancel()

	loginToken := models.LoginToken{
		UserId: userId,
	}

	result, err := loginTokensCollection.InsertOne(ctx, loginToken)
	if err != nil {
		log.Printf("Failed to create login token in database.CreateLoginToken(): %v", err)
		return ""
	}

	insertedId, ok := result.InsertedID.(bson.ObjectID)
	if !ok {
		err = errors.New("expected type of result.InsertedID to be bson.ObjectID")
		log.Printf("Failed to decode created login token in database.CreateLoginToken(): %v", err)
		return ""
	}

	return insertedId.Hex()
}

func DeleteLoginTokens(loginTokenIds []bson.ObjectID) {
	ctx, cancel := context.WithTimeout(context.Background(), config.DatabaseTimeoutDuration)
	defer cancel()

	_, err := loginTokensCollection.DeleteMany(ctx, bson.M{"_id": bson.M{"$in": loginTokenIds}})
	if err != nil {
		log.Printf("Failed to delete login tokens in database.DeleteLoginTokens(): %v", err)
	}
}

func DeleteLoginToken(loginTokenId bson.ObjectID) {
	log.Println("database delete login token function called [+]")
	ctx, cancel := context.WithTimeout(context.Background(), config.DatabaseTimeoutDuration)
	defer cancel()

	_, err := loginTokensCollection.DeleteOne(ctx, bson.M{"_id": loginTokenId})
	if err != nil {
		log.Printf("Failed to delete login token in database.DeleteLoginToken(): %v", err)
	}
}

func DeleteLoginTokensByUserId(userId bson.ObjectID) {
	ctx, cancel := context.WithTimeout(context.Background(), config.DatabaseTimeoutDuration)
	defer cancel()

	filter := bson.M{"userId": userId}

	result, err := loginTokensCollection.DeleteMany(ctx, filter)
	if err != nil {
		log.Printf("Failed to delete login tokens in database.DeleteLoginTokensByUserId(): %v", err)
		return
	}

	log.Printf("Deleted '%d' login tokens for user '%s'", result.DeletedCount, userId.Hex())
}
