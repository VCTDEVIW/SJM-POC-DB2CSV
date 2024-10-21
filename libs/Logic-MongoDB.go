package project

import (
    "context"
	"encoding/json"
	//"encoding/csv"
    . "fmt"
    "log"
    //"os"
	_"time"

	//"go.mongodb.org/mongo-driver/bson"
    _ "go.mongodb.org/mongo-driver/mongo"
    _ "go.mongodb.org/mongo-driver/mongo/options"
)

func MongoDB_Job1(data map[string]interface{}) {
	jsonDoc, _ := json.MarshalIndent(data, "", "  ")
	Printf("Document found: %s\n", jsonDoc)
}

func (load *META_Global) MongoDB_Test() {
	//username := os.Getenv("MONGO_USERNAME")
    //password := os.Getenv("MONGO_PASSWORD")

	host := load.LoadConfig.MongoDB.Host
	port := load.LoadConfig.MongoDB.Port
	username := load.LoadConfig.MongoDB.Username
	password := load.LoadConfig.MongoDB.Password


    // Connect to MongoDB
    client, err := MongoDB_Conn(host, port, username, password)
    if err != nil {
        log.Fatal(err)
    }
    defer client.Disconnect(context.TODO())

	/*
	JsonQuery := `
	{
        "text1": {
            "$exists": true,
            "$ne": ""
        }
    }`
	*/

	//JsonQuery := `{ "text1": "hello" }`
	JsonQuery := load.LoadConfig.Options.MongoDBEmbedJSON
	Filter, _ := MongoDB_ParseQueryFromString(JsonQuery)

	DbName := load.LoadConfig.MongoDB.DBName        // Change the database name as needed
	ColtName := load.LoadConfig.MongoDB.Collection // Change the collection name as needed
	Schema := client.Database(DbName).Collection(ColtName)

    // Perform the read operation with the callback
	if err := MongoDB_Read(Schema, Filter, MongoDB_Job1); err != nil {
		log.Fatal(err)
	}
}

func (load *META_Global) MongoDB_RunQuery() {
		//username := os.Getenv("MONGO_USERNAME")
    //password := os.Getenv("MONGO_PASSWORD")

	host := load.LoadConfig.MongoDB.Host
	port := load.LoadConfig.MongoDB.Port
	username := load.LoadConfig.MongoDB.Username
	password := load.LoadConfig.MongoDB.Password


    // Connect to MongoDB
    client, err := MongoDB_Conn(host, port, username, password)
    if err != nil {
        log.Fatal(err)
    }
    defer client.Disconnect(context.TODO())

	/*
	JsonQuery := `
	{
        "text1": {
            "$exists": true,
            "$ne": ""
        }
    }`
	*/

	//JsonQuery := `{ "text1": "hello" }`
	JsonQuery := load.LoadConfig.Options.MongoDBEmbedJSON
	Filter, _ := MongoDB_ParseQueryFromString(JsonQuery)

	DbName := load.LoadConfig.MongoDB.DBName        // Change the database name as needed
	ColtName := load.LoadConfig.MongoDB.Collection // Change the collection name as needed
	Schema := client.Database(DbName).Collection(ColtName)

    // Perform the read operation with the callback
	cursor, err := Schema.Find(nil, Filter) // No context used
	if err != nil {
		log.Println(err)
	}
	defer cursor.Close(nil)

	for cursor.Next(nil) {
		var result map[string]interface{}
		if err := cursor.Decode(&result); err != nil {
			log.Println(err)
		}
		// Call the provided callback with the result
		jsonDoc, _ := json.MarshalIndent(result, "", "  ")
		Printf("Document found: %s\n", jsonDoc)
	}

	if err := cursor.Err(); err != nil {
		log.Println(err)
	}

}

