package main

import (
    "context"
	"encoding/json"
    . "fmt"
    "log"
    _"os"
	_"time"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	//username := os.Getenv("MONGO_USERNAME")
    //password := os.Getenv("MONGO_PASSWORD")

	host := "10.10.10.20"
	port := "27017"
	username := "root"
	password := "1234"


    // Connect to MongoDB
    client, err := MongoDB_Conn(host, port, username, password)
    if err != nil {
        log.Fatal(err)
    }
    defer client.Disconnect(context.TODO())

	JsonQuery := `
	{
        "text1": {
            "$exists": true,
            "$ne": ""
        }
    }`

	//JsonQuery := `{ "text1": "hello" }`
	Filter, _ := MongoDB_ParseQueryFromString(JsonQuery)

	DbName := "test_db"        // Change the database name as needed
	ColtName := "info" // Change the collection name as needed
	Schema := client.Database(DbName).Collection(ColtName)

    // Perform the read operation with the callback
	if err := MongoDB_Read(Schema, Filter, MongoDB_Job1); err != nil {
		log.Fatal(err)
	}
}


// Parse query from HEREDOC-like string
func MongoDB_ParseQueryFromString(query string) (bson.M, error) {
	var filter bson.M
	if err := json.Unmarshal([]byte(query), &filter); err != nil {
		return nil, err
	}

	return filter, nil
}


// connectToMongo establishes a connection to the MongoDB server.
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

func MongoDB_Job1(data map[string]interface{}) {
	jsonDoc, _ := json.MarshalIndent(data, "", "  ")
	Printf("Document found: %s\n", jsonDoc)
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
