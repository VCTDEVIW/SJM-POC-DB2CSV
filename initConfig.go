package main

import (
	. "fmt"
	"os"
	"encoding/json"
	. "project/workspace/sjm-poc-db/libs"
)

func main() {
	filename := ConfigFilename

    // Check if the file exists
    if _, err := os.Stat(filename); os.IsNotExist(err) {
        // File does not exist, create it
        file, err := os.Create(filename)
        if err != nil {
            Println("Error creating file:", err)
            return
        }
        defer file.Close()

        // Write the default config with pretty print
        data, err := json.MarshalIndent(GenInitConfigFile(), "", "  ")
        if err != nil {
            Println("Error marshaling JSON:", err)
            return
        }

        if _, err := file.Write(data); err != nil {
            Println("Error writing to file:", err)
        }
        Println("Created config.json with default settings.")
    } else {
        // File exists, check if it's empty
        fileInfo, err := os.Stat(filename)
        if err != nil {
            Println("Error checking file:", err)
            return
        }

        if fileInfo.Size() == 0 {
            // File is empty, write the default config with pretty print
            file, err := os.OpenFile(filename, os.O_WRONLY, 0644)
            if err != nil {
                Println("Error opening file:", err)
                return
            }
            defer file.Close()

            data, err := json.MarshalIndent(GenInitConfigFile(), "", "  ")
            if err != nil {
                Println("Error marshaling JSON:", err)
                return
            }

            if _, err := file.Write(data); err != nil {
                Println("Error writing to file:", err)
            }
            Println("config.json was empty, wrote default settings to it.")
        } else {
            Println("config.json already exists and is not empty.")
        }
    }
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

