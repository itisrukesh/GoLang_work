package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB mongo.Database

//connection establishment.
func ConnectMongo() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, errtimeout := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	fmt.Println("Connected to MongoDB...")
	DB = *client.Database("form")
	fmt.Println(errtimeout)

}

func init() {
	ConnectMongo() //calling it.
}
