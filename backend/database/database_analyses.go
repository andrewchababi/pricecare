package database

import (
	"context"
	"log"

	"github.com/andrewchababi/pricecare/backend/config"
	"github.com/andrewchababi/pricecare/backend/models"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func CreateAnalysis(analysis models.Analysis) {
	ctx, cancel := context.WithTimeout(context.Background(), config.DatabaseTimeoutDuration)
	defer cancel()
	_, err := analysesCollection.InsertOne(ctx, analysis)

	if err != nil {
		log.Printf("Failed to insert analysis into database: %v", err)
	}
}

func GetAnalysesFromTestIds(testIds []string) []models.Analysis {
	ctx, cancel := context.WithTimeout(context.Background(), config.DatabaseTimeoutDuration)
	defer cancel()

	filter := bson.M{"testId": bson.M{"$in": testIds}}

	cursor, err := analysesCollection.Find(ctx, filter)
	if err != nil {
		return []models.Analysis{}
	}
	defer cursor.Close(ctx)

	var analyses []models.Analysis
	if err = cursor.All(ctx, &analyses); err != nil {
		return []models.Analysis{}
	}

	return analyses
}
