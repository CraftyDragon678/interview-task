package framework

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

// DB database
var DB *mongo.Database

// InitDB making connection
func InitDB() {
	var full string
	if full = os.Getenv("DB_FULL"); full == "" {
		full = fmt.Sprintf("mongodb://%s:%s@%s:%s/%s?authSource=admin",
			os.Getenv("DB_ID"), url.QueryEscape(os.Getenv("DB_PW")),
			os.Getenv("DB_HOST"), os.Getenv("DB_PORT"),
			os.Getenv("DB_DB"))
	}
	c, err := mongo.NewClient(options.Client().ApplyURI(full))

	if err != nil {
		log.Fatalln("Can't make db client")
	}
	client = c
	DB = client.Database("xsi")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = c.Connect(ctx)

	if err != nil {
		log.Fatalln("Can't connect db")
	}
}

// ExitDB disconnect
func ExitDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := client.Disconnect(ctx)

	if err != nil {
		log.Fatalln("[ERROR] in disconnect db")
	}
}
