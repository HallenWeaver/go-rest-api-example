package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Character struct {
	ID      primitive.ObjectID `bson:_id,omitempty`
	OwnerId string             `bson:"owner_id"`
	Name    string             `bson:"name"`
	Age     int                `bson:"age"`
}
