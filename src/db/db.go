package db

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	MongoDB *mongo.Database
}

func DBConnection() (*DB, error) {
	errEnv := godotenv.Load(".env")
	if errEnv != nil {
		log.Printf("%s", errEnv)
	}

	uri := os.Getenv("URI")

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		return nil, err
	}
	mdb := client.Database("book-store-dev")

	return &DB{
		mdb,
	}, nil
}
