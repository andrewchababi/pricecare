package models

import "go.mongodb.org/mongo-driver/v2/bson"

type LoginToken struct {
	Id     bson.ObjectID `bson:"_id,omitempty" json:"id"`
	UserId bson.ObjectID `bson:"userId"        json:"userId"`
}
