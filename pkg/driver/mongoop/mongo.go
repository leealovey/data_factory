package mongoop

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoConfig mongodb configuration
type MongoConfig struct {
	URI string
}

// Conn mongoDB Connector
func (m *MongoConfig) Conn() (*mongo.Client, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(m.URI))
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	collection := client.Database("test").Collection("col")

	res, err := collection.InsertOne(ctx, bson.M{"hello": "world"})
	if err != nil {
		return nil, err
	}
	id := res.InsertedID
	log.Println(id)

	return client, nil
}
