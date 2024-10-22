package project

import (
    "encoding/json"
    . "fmt"
)

// Config represents the entire configuration structure.
type JsonConfig struct {
    Options struct {
        API_WebPort                 int    `json:"API_WebPort"`
        API_SqlJob_URI              string `json:"API_SqlJob_URI"`
        API_MongoDBJob_URI          string `json:"API_MongoDBJob_URI"`
        API_AccessGetToken          string `json:"API_AccessGetToken"`
        SqlSrvOutputFile            string `json:"SqlSrv_output_file"`
        MongoDBOutputFile           string `json:"mongodb_output_file"`
        MongoDBUseJSONQueryFile     string `json:"mongodb_use_jsonQuery_file"`
        MongoDBUseEmbedJSONQuery    string `json:"mongodb_use_embed_jsonQuery"`
        MongoDBEmbedJSON            string `json:"mongodb_embed_json"`
    } `json:"options"`
    SqlSrv struct {
        Host     string `json:"sql-db-host"`
        Port     int    `json:"sql-db-port"`
        Username string `json:"sql-db-username"`
        Password string `json:"sql-db-password"`
        DBName   string `json:"sql-db-dbName"`
        Table    string `json:"sql-db-table"`
        Query    string `json:"sql-db-query"`
    } `json:"SqlSrv"`
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
    config := JsonConfig{}
    
    // Populate the config structure
    config.Options.API_WebPort = 8080
    config.Options.API_SqlJob_URI = "sqlsrv-job"
    config.Options.API_MongoDBJob_URI = "mongodb-job"
    config.Options.API_AccessGetToken = "AbcD@1234"
    config.Options.SqlSrvOutputFile = "SqlSrv_output.csv"
    config.Options.MongoDBOutputFile = "mongodb_output.csv"
    config.Options.MongoDBUseJSONQueryFile = "yes"
    config.Options.MongoDBUseEmbedJSONQuery = "yes"
    config.Options.MongoDBEmbedJSON = "{ \"text1\":{ \"$exists\": true, \"$ne\": \"\" } }"
    
    config.SqlSrv.Host = "10.10.10.20"
    config.SqlSrv.Port = 1433
    config.SqlSrv.Username = "sa"
    config.SqlSrv.Password = "P@ssw0rd"
    config.SqlSrv.DBName = "test_db"
    config.SqlSrv.Table = "table1"
    config.SqlSrv.Query = ""
    
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

func GenInitConfigFile() JsonConfig {
	config := JsonConfig{}
    
    // Populate the config structure
	config.Options.API_WebPort = 8080
	config.Options.API_SqlJob_URI = "sqlsrv-job"
    config.Options.API_MongoDBJob_URI = "mongodb-job"
	config.Options.API_AccessGetToken = "AbcD@1234"
    config.Options.SqlSrvOutputFile = "SqlSrv_output.csv"
    config.Options.MongoDBOutputFile = "mongodb_output.csv"
    config.Options.MongoDBUseJSONQueryFile = "yes"
    config.Options.MongoDBUseEmbedJSONQuery = "yes"
    config.Options.MongoDBEmbedJSON = "{ \"text1\":{ \"$exists\": true, \"$ne\": \"\" } }"
    
    config.SqlSrv.Host = "10.10.10.20"
    config.SqlSrv.Port = 1433
    config.SqlSrv.Username = "sa"
    config.SqlSrv.Password = "P@ssw0rd"
    config.SqlSrv.DBName = "test_db"
    config.SqlSrv.Table = "table1"
    config.SqlSrv.Query = ""
    
    config.MongoDB.Host = "10.10.10.20"
    config.MongoDB.Port = "27017"
    config.MongoDB.Username = "root"
    config.MongoDB.Password = "1234"
    config.MongoDB.DBName = "test_db"
    config.MongoDB.Collection = "info"

	return config
}

