package project

import (
    "context"
	"encoding/json"
    . "fmt"
    _"log"
    _"os"
	_"time"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func MongoDB_Conn(host, port, username, password string) (*mongo.Client, error) {
    var uri string
    
    // Check if username and password are set
    if username != "" && password != "" {
        uri = Sprintf("mongodb://%s:%s@%s:%s", username, password, host, port)
    } else {
		uri = Sprintf("mongodb://%s:%s", host, port)	// No authentication
    }

    // Set client options
    clientOptions := options.Client().ApplyURI(uri)

    // Connect to MongoDB
    client, err := mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
        return nil, err
    }

    // Check the connection
    if err := client.Ping(context.TODO(), nil); err != nil {
        return nil, err
    }
    Println("Connected to MongoDB!")
    return client, nil
}

func MongoDB_Read(collection *mongo.Collection, filter interface{}, callback func(map[string]interface{})) error {
	cursor, err := collection.Find(nil, filter) // No context used
	if err != nil {
		return err
	}
	defer cursor.Close(nil)

	for cursor.Next(nil) {
		var result map[string]interface{}
		if err := cursor.Decode(&result); err != nil {
			return err
		}
		// Call the provided callback with the result
		callback(result)
	}

	if err := cursor.Err(); err != nil {
		return err
	}

	return nil
}

func MongoDB_ParseQueryFromString(query string) (bson.M, error) {
	var filter bson.M
	if err := json.Unmarshal([]byte(query), &filter); err != nil {
		return nil, err
	}

	return filter, nil
}

