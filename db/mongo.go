package db

import (
	"context"
	"log"

	"github.com/gopheramol/notesapp/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB(config *config.Config) (client *mongo.Client, err error) {
	clientOptions := options.Client().ApplyURI(config.DBSource)
	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}
	log.Println("Connected to MongoDB!")
	return client, nil
}
