package model

type Character struct {
	Id      string `json:"id"`
	OwnerId string `json:"ownerId"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
}
