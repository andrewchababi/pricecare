package database

import (
	"context"
	"log"

	"github.com/andrewchababi/pricecare/backend/config"
	"github.com/andrewchababi/pricecare/backend/models"
)

func CreateAnalysis(analysis models.Analysis) {
	ctx, cancel := context.WithTimeout(context.Background(), config.DatabaseTimeoutDuration)
	defer cancel()
	_, err := analysesCollection.InsertOne(ctx, analysis)

	if err != nil {
		log.Printf("Failed to insert analysis into database: %v", err)
	}
}
