package project

import (
    "encoding/json"
    . "fmt"
)

// Config represents the entire configuration structure.
type Config struct {
    Options struct {
        SQLSrvOutputFile            string `json:"sqlsrv_output_file"`
        MongoDBOutputFile           string `json:"mongodb_output_file"`
        MongoDBUseJSONQueryFile     string `json:"mongodb_use_jsonQuery_file"`
        MongoDBUseEmbedJSONQuery    string `json:"mongodb_use_embed_jsonQuery"`
        MongoDBEmbedJSON            string `json:"mongodb_embed_json"`
    } `json:"options"`
    SQLSrv struct {
        Host     string `json:"sql-db-host"`
        Port     int    `json:"sql-db-port"`
        Username string `json:"sql-db-username"`
        Password string `json:"sql-db-password"`
        DBName   string `json:"sql-db-dbName"`
        Table    string `json:"sql-db-table"`
        Query    string `json:"sql-db-query"`
    } `json:"sqlsrv"`
    MongoDB struct {
        Host       string `json:"mgdb-host"`
        Port       string `json:"mgdb-port"`
        Username   string `json:"mgdb-username"`
        Password   string `json:"mgdb-password"`
        DBName     string `json:"mgdb-dbName"`
        Collection string `json:"mgdb-collection"`
    } `json:"mongodb"`
}

func MockJson() {
    // Example of how to create an instance of Config
    config := Config{}
    
    // Populate the config structure
    config.Options.SQLSrvOutputFile = "sqlsrv_output.csv"
    config.Options.MongoDBOutputFile = "mongodb_output.csv"
    config.Options.MongoDBUseJSONQueryFile = "yes"
    config.Options.MongoDBUseEmbedJSONQuery = "yes"
    config.Options.MongoDBEmbedJSON = "{ \"text1\":{ \"$exists\": true, \"$ne\": \"\" } }"
    
    config.SQLSrv.Host = "10.10.10.20"
    config.SQLSrv.Port = 1433
    config.SQLSrv.Username = "sa"
    config.SQLSrv.Password = "P@ssw0rd"
    config.SQLSrv.DBName = "test_db"
    config.SQLSrv.Table = "table1"
    config.SQLSrv.Query = ""
    
    config.MongoDB.Host = "10.10.10.20"
    config.MongoDB.Port = "27017"
    config.MongoDB.Username = "root"
    config.MongoDB.Password = "1234"
    config.MongoDB.DBName = "test_db"
    config.MongoDB.Collection = "info"

    // Example of marshaling to JSON
    jsonData, err := json.MarshalIndent(config, "", "  ")
    if err != nil {
        Println("Error marshaling to JSON:", err)
        return
    }

    Println(string(jsonData))
}

