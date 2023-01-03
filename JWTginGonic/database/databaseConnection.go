package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv" //loads env variables from a .env file
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// create a func that returns us a mongo client
func DBinstance() *mongo.Client {
	err := godotenv.Load(".env") //Load will read your env file(s) and load them into ENV for this process.
	if err != nil {
		log.Fatal("Error loading the .env file")
	}

	MongoDB := os.Getenv("MONGODB_URL") //Getenv retrieves the value of the environment variable named by the key.

	client, err := mongo.NewClient(options.Client().ApplyURI(MongoDB)) //NewClient creates a new client to connect to a deployment specified by the uri.
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()            //Canceling this context releases resources associated with it, so code should call cancel as soon as the operations running in this Context complete
	err = client.Connect(ctx) //connect to the client using context
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	return client
}

var Client *mongo.Client = DBinstance() //variable client of type *mnogo.Client

func OpenCollectiion(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("cluster0").Collection(collectionName)
	return collection
}
