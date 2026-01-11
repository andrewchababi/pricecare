package database

import (
	"context"
	"log"

	"github.com/andrewchababi/pricecare/backend/config"
	"github.com/andrewchababi/pricecare/backend/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"golang.org/x/crypto/bcrypt"
)

// TEMP only used to populate database
func populateDatabase() {
	populateUsers()
	populateAnalyses()

	log.Println("Database has been populated")
}

// TEMP only used to populate database
func populateAnalyses() {
	ctx, cancel := context.WithTimeout(context.Background(), config.DatabaseTimeoutDuration)
	defer cancel()
	analysesCollection.DeleteMany(ctx, bson.M{})

	// populateLocation("235 Barclay Boulevard", "3235 Av. Barclay, Montréal, QC H3S 1K2", 20)
	populateAnalysis("THC", "test", 30, 1.2)
}

// TEMP only used to populate database
func populateAnalysis(testId string, name string, listPrice int, cost float64) {
	analysis := models.Analysis{
		Id:          bson.NewObjectID(),
		TestId:      testId,
		Name:        name,
		ListPrice:   listPrice,
		ReagentCost: cost,
	}

	CreateAnalysis(analysis)
}

// TEMP only used to populate database
func populateUsers() {
	ctx, cancel := context.WithTimeout(context.Background(), config.DatabaseTimeoutDuration)
	defer cancel()
	usersCollection.DeleteMany(ctx, bson.M{})

	populateUser("staff", "staff123", models.UserTypeStaff)
}

// TEMP only used to populate database
func populateUser(username string, password string, userType models.UserType) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 10)

	user := models.User{
		Id:             bson.NewObjectID(),
		Username:       username,
		HashedPassword: string(hashedPassword),
		UserType:       userType,
	}

	CreateUser(user)
}
