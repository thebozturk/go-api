package configs

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client {
	// create new mongo client
	client, err := mongo.NewClient(options.Client().ApplyURI(EnvMongoURI()))

	// if error creating client, log error and exit
	if err != nil {
		log.Fatalln(err)
	}

	// create context with timeout of 10 seconds
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	// if error connecting, log error and exit
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalln(err)
	}
	return client
}

// DB is the exported variable that holds the connection to the database
var DB *mongo.Client = ConnectDB()

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	// get collection from database
	collection := client.Database("go-api").Collection(collectionName)
	return collection
}
