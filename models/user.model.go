package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Address struct {
	State   string `json:"state" bson:"state"`
	City    string `json:"city" bson:"city"`
	PinCode int    `json:"pincode" bson:"pincode" `
}

type User struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name    string             `json:"name" bson:"name"`
	Age     int                `json:"age" bson:"age"`
	Address Address            `json:"address" bson:"address"`
}
