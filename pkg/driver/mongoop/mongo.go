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

// SendCollections mongoDB Connector
func (m *MongoConfig) SendCollections() ([]byte, error) {
	res := []byte{}
	colls := []string{}

	client, err := mongo.NewClient(options.Client().ApplyURI(m.URI))
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	// create test data
	// collection := client.Database("test").Collection("col")

	// for i := 0; i < 1000; i++ {
	// 	data := []interface{}{
	// 		bson.M{
	// 			"index": i,
	// 			"data":  "test",
	// 			"host":  "localhost",
	// 			"bing":  "test",
	// 		},
	// 	}

	// 	opts := options.InsertMany().SetOrdered(false)
	// 	res, err := collection.InsertMany(ctx, data, opts)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Printf("inserted documents with IDs %v\n", res.InsertedIDs)
	// }

	db := client.Database("lsqt_db")

	result, err := db.ListCollectionNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range result {
		colls = append(colls, v)
		// coll := db.Collection(v)

		// cursor, err := coll.Find(ctx, bson.M{})
		// if err != nil {
		// 	return nil, err
		// }

		// var results []interface{}
		// if err = cursor.All(ctx, &results); err != nil {
		// 	log.Fatal(err)
		// }

		// res, err = json.Marshal(results)
		// if err != nil {
		// 	return nil, err
		// }

	}

	fmt.Println(colls)

	coll := db.Collection(colls[0])

	cursor, err := coll.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	var results []bson.M
	if err = cursor.All(ctx, &results); err != nil {
		log.Fatal(err)
	}

	res, err = json.Marshal(results)
	if err != nil {
		return nil, err
	}

	store := []DayAccount{}
	err = json.Unmarshal(res, &store)
	if err != nil {
		return nil, err
	}

	fmt.Printf("%+v", store[0])

	return res, nil
}
