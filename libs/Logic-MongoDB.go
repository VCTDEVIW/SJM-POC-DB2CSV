package project

import (
    "context"
	"encoding/json"
    . "fmt"
    "log"
    _"os"
	_"time"
)

func MongoDB_Job1(data map[string]interface{}) {
	jsonDoc, _ := json.MarshalIndent(data, "", "  ")
	Printf("Document found: %s\n", jsonDoc)
}

func (load *META_Global) MongoDB_Test() {
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

