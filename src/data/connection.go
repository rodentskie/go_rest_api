package db

import (
	"context"
	"fmt"
	env "go-rest-api/src/functions"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Start() {
	fmt.Println("Starting Db...")

	uri := env.GetEnv("MONGO_URI", "mongodb://localhost/go_test")
	fmt.Println(uri)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}
