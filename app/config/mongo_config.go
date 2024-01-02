package config

import (
	"context"
	"time"

	"github.com/rezaif79-ri/iris-api-101/app/util"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func OpenMongoBookDB() (*mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Get mongo connection from ENV
	mongoHost := util.GetEnv("MONGO_HOST", "localhost")
	mongoPort := util.GetEnv("MONGO_PORT", "27027")
	mongoProtocol := util.GetEnv("MONGO_PROTOCOL", "mongodb")
	mongoCreds := util.GetEnv("MONGO_CREDS", "")

	connUri := mongoProtocol + "://" + mongoCreds + mongoHost + ":" + mongoPort
	clientOptions := options.Client()
	clientOptions.ApplyURI(connUri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	return client.Database("books_database"), nil
}
