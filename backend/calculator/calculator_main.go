package calculator

import (
	"log"
	"sort"

	"github.com/andrewchababi/pricecare/backend/database"
	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	marginalOverhead     = 1.00
	marginalTestDiscount = 0.5
	perReqDonation       = 5.00
	revSharePerReq       = 0.5
)

func CalculatePanelPrice(testIds []string) float64 {

	analyses := database.GetAnalysesFromTestIds(testIds)
	sort.Slice(analyses, func(i, j int) bool {
		return analyses[i].ListPrice > analyses[j].ListPrice
	})

	totalPrice := 0.00
	totalVariableCost := 0.0

	anchor := analyses[0]
	totalPrice += float64(anchor.ListPrice)
	totalVariableCost += anchor.ReagentCost + marginalOverhead

	for _, test := range analyses[1:] {
		targetPrice := float64(test.ListPrice) * marginalTestDiscount
		floorPrice := (test.ReagentCost + marginalOverhead) * 3.0
		finalPrice := max(targetPrice, floorPrice)

		totalPrice += finalPrice
	}

	totalPrice += perReqDonation + revSharePerReq

	return totalPrice
}

func convertHexToObjectId(ids []string) []bson.ObjectID {
	bsonIds := make([]bson.ObjectID, 0, len(ids))

	for _, id := range ids {
		objectID, err := bson.ObjectIDFromHex(id)
		if err != nil {
			// Handle invalid ID - skip it or return error
			log.Printf("invalid bson id")
			continue
		}
		bsonIds = append(bsonIds, objectID)
	}

	return bsonIds
}
