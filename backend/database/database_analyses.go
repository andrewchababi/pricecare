package database

import (
	"context"
	"log"

	"github.com/andrewchababi/pricecare/backend/config"
	"github.com/andrewchababi/pricecare/backend/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func CreateAnalysis(analysis models.Analysis) {
	ctx, cancel := context.WithTimeout(context.Background(), config.DatabaseTimeoutDuration)
	defer cancel()
	_, err := analysesCollection.InsertOne(ctx, analysis)

	if err != nil {
		log.Printf("Failed to insert analysis into database: %v", err)
	}
}

func GetAnalyses() []models.Analysis {
	ctx, cancel := context.WithTimeout(context.Background(), config.DatabaseTimeoutDuration)
	defer cancel()

	findOptions := options.Find().SetSort(bson.D{{"testId", 1}})

	cursor, err := analysesCollection.Find(ctx, bson.D{}, findOptions)
	if err != nil {
		log.Printf("Failed to get analysis by id: %v", err)
		return []models.Analysis{}
	}
	defer cursor.Close(ctx)

	var analyses []models.Analysis
	err = cursor.All(ctx, &analyses)
	if err != nil {
		log.Printf("Failed to decode analyses in database.GetAnalyses(): %v", err)
		return []models.Analysis{}
	}

	return analyses
}

func GetAnalysesFromTestIds(testIds []string) []models.Analysis {
	ctx, cancel := context.WithTimeout(context.Background(), config.DatabaseTimeoutDuration)
	defer cancel()

	filter := bson.M{"testId": bson.M{"$in": testIds}}
	findOptions := options.Find().SetSort(bson.D{{"testId", 1}})

	cursor, err := analysesCollection.Find(ctx, filter, findOptions)
	if err != nil {
		log.Printf("Failed to get analyses by id: %v", err)
		return []models.Analysis{}
	}
	defer cursor.Close(ctx)

	var analyses []models.Analysis
	if err = cursor.All(ctx, &analyses); err != nil {
		return []models.Analysis{}
	}

	return analyses
}
