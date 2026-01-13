package database

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/andrewchababi/pricecare/backend/config"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

var (
	client                *mongo.Client
	database              *mongo.Database
	usersCollection       *mongo.Collection
	analysesCollection    *mongo.Collection
	loginTokensCollection *mongo.Collection
)

func init() {
	godotenv.Load()

	client = getClient()
	database = getDatabase()
	usersCollection = database.Collection(config.DatabaseCollectionUsers)
	analysesCollection = database.Collection(config.DatabaseCollectionAnalyses)
	loginTokensCollection = database.Collection(config.DatabaseCollectionLoginTokens)

	log.Println("Connected to database")

	// TEMP
	// populateDatabase()
	// clearCollection(loginTokensCollection)
}

func getEnvironmentvariable(name string) string {
	variable := os.Getenv(name)
	if variable == "" {
		log.Fatalf("Environment variable %s is not set", name)
	}
	return variable
}

func getClient() *mongo.Client {
	uri := getEnvironmentvariable(config.EnvVarDatabaseUrl)

	serverApi := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI(uri).SetServerAPIOptions(serverApi)
	client, err := mongo.Connect(clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to mongodb: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatalf("Failed to ping mongodb: %v", err)
	}

	return client
}

func getDatabase() *mongo.Database {
	databaseName := getEnvironmentvariable(config.EnvVarDatabaseName)
	return client.Database(databaseName)
}
