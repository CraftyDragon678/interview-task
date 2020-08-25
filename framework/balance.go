package framework

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

// GetMoney return amount by id
func GetMoney(id string) int {
	users := DB.Collection("users")
	if result := users.FindOne(context.TODO(), bson.M{"id": id}); result.Err() == nil {
		var user User
		result.Decode(&user)
		return user.Money
	}
	return -1
}

// GiveMoney give to user.
// if negative, take from user
func GiveMoney(id string, amount int) {
	users := DB.Collection("users")
	users.FindOneAndUpdate(context.TODO(), bson.M{"id": id}, bson.M{
		"$inc": bson.M{
			"money": amount,
		},
	})
}
