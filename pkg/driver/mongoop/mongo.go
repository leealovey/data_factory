package mongoop

import (
	"context"
	"encoding/json"
	"fmt"
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

	db := client.Database("test")

	result, err := db.ListCollectionNames(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)

	for _, v := range result {
		coll := db.Collection(v)

		cursor, err := coll.Find(ctx, bson.M{})
		if err != nil {
			return nil, err
		}

		var results []interface{}
		if err = cursor.All(context.TODO(), &results); err != nil {
			log.Fatal(err)
		}

		rawjson, err := json.Marshal(results)
		if err != nil {
			return nil, err
		}

		fmt.Println(string(rawjson))
		// for _, r := range results {
		// 	fmt.Println(r)
		// }
	}

	return client, nil
}
