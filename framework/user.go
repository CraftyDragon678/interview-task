package framework

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

// User is user
type User struct {
	// ID is discord user id
	ID string
}

// CheckUserExist return if user exist in database
func CheckUserExist(id string) bool {
	users := DB.Collection("users")
	result := users.FindOne(context.TODO(), User{ID: id})
	return result.Err() != mongo.ErrNoDocuments
}

// AddUser add user in users collection
func AddUser(id string) {
	users := DB.Collection("users")
	users.InsertOne(context.TODO(), User{ID: id})
}
