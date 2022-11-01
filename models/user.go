package models

type Address struct {
	stage   string `json: "stage" bson: "state"`
	city    string `json: "city" bson: "city"`
	pinCode int    `json: "pinCode" bson: "pinCode" `
}

type User struct {
	id      string  `json: "id" bson: "user_id"`
	name    string  `json: "name" bson:  "user_name"`
	age     int     `json: "age" bson: "user_age"`
	address Address `json: "address" bson: "user_address"`
}
