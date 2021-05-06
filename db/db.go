package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/AskJag07/virtuoso-server/config"
)

func Init() *mongo.Client {

	MongoDBUri := config.GetVar("MONGODB_URI")

	client, err := mongo.NewClient(options.Client().ApplyURI(MongoDBUri))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return client

}
