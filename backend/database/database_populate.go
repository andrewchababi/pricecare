package database

import (
	"context"
	"log"

	"github.com/andrewchababi/pricecare/backend/config"
	"github.com/andrewchababi/pricecare/backend/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"golang.org/x/crypto/bcrypt"
)

func clearCollection(collection *mongo.Collection) {
	ctx, cancel := context.WithTimeout(context.Background(), config.DatabaseTimeoutDuration)
	defer cancel()

	collection.DeleteMany(ctx, bson.M{})
}

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

	// Map of testId -> [name, listPrice, reagentCost]
	tests := map[string][3]interface{}{
		"A1C":          {"Hemoglobin A1C", 49, 2.07},
		"ALB":          {"Albumin", 40, 0.36},
		"ALP":          {"Alkaline Phosphatase", 31, 0.24},
		"ALT":          {"Alanine Aminotransferase", 30, 0.14},
		"AST":          {"Aspartate Aminotransferase", 30, 0.09},
		"BILI_TOT":     {"Total Bilirubin", 31, 0.24},
		"BILI_DIR":     {"Direct Bilirubin", 30, 0.17},
		"CALCIUM":      {"Calcium", 31, 0.24},
		"CHOL":         {"Cholesterol", 48, 0.24},
		"CK":           {"Creatine Kinase", 39, 0.24},
		"CO2":          {"Carbon Dioxide", 34, 0.50},
		"CREAT":        {"Creatinine", 39, 1.01},
		"CRP":          {"C-Reactive Protein", 41, 1.21},
		"FERRITIN":     {"Ferritin", 41, 1.21},
		"FOLATE":       {"Folate", 40, 1.38},
		"FT3":          {"Free T3", 31, 0.94},
		"FT4":          {"Free T4", 31, 0.94},
		"GGT":          {"Gamma-Glutamyl Transferase", 31, 0.24},
		"GLUCOSE":      {"Glucose", 31, 0.24},
		"HDL":          {"HDL Cholesterol", 30, 0.77},
		"IRON":         {"Iron", 30, 0.15},
		"LDH":          {"Lactate Dehydrogenase", 37, 0.13},
		"LIPASE":       {"Lipase", 37, 0.55},
		"MAGNESIUM":    {"Magnesium", 31, 0.22},
		"PHOS":         {"Phosphorus", 30, 0.13},
		"PSA_TOT":      {"Total PSA", 42, 1.32},
		"PSA_FREE":     {"Free PSA", 43, 1.93},
		"TRIG":         {"Triglycerides", 32, 0.24},
		"TSH":          {"Thyroid Stimulating Hormone", 36, 0.71},
		"UREA":         {"Urea", 30, 0.17},
		"URIC_ACID":    {"Uric Acid", 30, 0.13},
		"VIT_B12":      {"Vitamin B12", 41, 1.21},
		"VIT_D":        {"Vitamin D", 59, 3.03},
		"CBC":          {"Complete Blood Count", 29, 0.50},
		"ELECTROLYTES": {"Electrolytes Panel", 46, 4.69},
		"HARMONY":      {"Harmony Prenatal Test", 1155, 289.00},
	}

	for testId, data := range tests {
		name := data[0].(string)
		listPrice := data[1].(int)
		reagentCost := data[2].(float64)
		populateAnalysis(testId, name, listPrice, reagentCost)
	}

	log.Println("Populated", len(tests), "analyses")
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
