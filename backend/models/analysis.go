package models

import "go.mongodb.org/mongo-driver/v2/bson"

type product interface {
	GetReagentCost() float64
	GetListPrice() int
}

type Analysis struct {
	Id          bson.ObjectID `bson:"_id,omitempty" json:"id"`
	TestId      string        `bson:"testId" json:"testId"`
	Name        string        `bson:"name" json:"name"`
	ReagentCost float64       `bson:"reagentCost" json:"reagentCost"`
	ListPrice   int           `bson:"listPrice" json:"listPrice"`
}

func (a Analysis) GetReagentCost() float64 {
	return a.ReagentCost
}

func (a Analysis) GetListPrice() int {
	return a.ListPrice
}
