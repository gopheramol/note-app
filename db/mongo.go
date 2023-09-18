package db

import (
	"context"
	"log"
	"time"

	"github.com/gopheramol/notesapp/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ConnectDB(config *config.Config) (client *mongo.Client, err error) {
	clientOptions := options.Client().ApplyURI(config.DBSource)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Println("database connection error", err)
		return nil, err
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Println("err", err)
		return
	}
	log.Println("Successfully connected and pinged.")
	return client, nil
}
