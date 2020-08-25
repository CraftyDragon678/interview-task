package framework

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// User is user
type User struct {
	ID    string
	Money int `bson:"money"`
}

// CheckUserExist return if user exist in database
func CheckUserExist(id string) bool {
	users := DB.Collection("users")
	result := users.FindOne(context.TODO(), bson.M{"id": id})
	return result.Err() != mongo.ErrNoDocuments
}

// AddUser add user in users collection
func AddUser(id string) {
	users := DB.Collection("users")
	users.InsertOne(context.TODO(), User{ID: id})
}
