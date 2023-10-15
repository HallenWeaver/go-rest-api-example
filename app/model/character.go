package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Character struct {
	ID      primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	OwnerId string             `json:"ownerid,omitempty" bson:"ownerid,omitempty"`
	Name    string             `json:"name,omitempty" bson:"name,omitempty"`
	Age     int                `json:"age,omitempty" bson:"age,omitempty"`
}
