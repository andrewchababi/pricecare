package calculator

import (
	"sort"

	"github.com/andrewchababi/pricecare/backend/database"
)

const (
	marginalOverhead     = 1.00
	marginalTestDiscount = 0.5
	perReqDonation       = 5.00
	revSharePerReq       = 0.5
)

func CalculatePanelPrice(ids []string) float64 {
	tests, _ := database.GetMultipleTestsByID(ids)
	sort.Slice(tests, func(i, j int) bool {
		return tests[i].ListPrice > tests[j].ListPrice
	})

	totalPrice := 0.00
	totalVariableCost := 0.0

	anchor := tests[0]
	totalPrice += float64(anchor.ListPrice)
	totalVariableCost += anchor.ReagentCost + marginalOverhead

	for _, test := range tests[1:] {
		targetPrice := float64(test.ListPrice) * marginalTestDiscount
		floorPrice := (test.ReagentCost + marginalOverhead) * 3.0
		finalPrice := max(targetPrice, floorPrice)

		totalPrice += finalPrice
	}

	totalPrice += perReqDonation + revSharePerReq

	return totalPrice
}
