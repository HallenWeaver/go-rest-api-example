package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Email    string             `json:"email,omitempty" bson:"email,omitempty"`
	Username string             `json:"username,omitempty" bson:"username,omitempty"`
	Password string             `json:"password,omitempty" bson:"password,omitempty"`
	Role     UserRole           `json:"role,omitempty" bson:"role,omitempty"`
	IsActive bool               `json:"isactive,omitempty" bson:"isactive,omitempty"`
}

type UserRole string

const (
	Standard UserRole = "Standard"
	Admin             = "Admin"
)
