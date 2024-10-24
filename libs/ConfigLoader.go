package project

import (
	. "fmt"
	"os"
	"encoding/json"
)

func FetchConfigFile(local_DebugMode int) string {
	if local_DebugMode == 1 {
		Println("Debug mode is set to ON [=1].")

		VAR_thisPath, _ := GetCurrentWorkingDirectory()
		VAR_pathPrefix := VAR_thisPath + GetOsPathSlash()
		VAR_configFile := VAR_pathPrefix + ConfigFilename
		return VAR_configFile
	} else {
		Println("Debug mode is set to OFF [=0].")

		return ConfigFilename
	}
}

func (load *META_Global) ReadConfigFile() {
	// Open the JSON file
    file, err := os.Open(FetchConfigFile(load.DebugMode))
    if err != nil {
        Println("Error opening config.json:", err)
        return
    }
    defer file.Close()

    // Create an instance of Config
    var Config JsonConfig

    // Decode the JSON data into the config struct
    decoder := json.NewDecoder(file)
    err = decoder.Decode(&Config)
    if err != nil {
        Println("Error decoding JSON:", err)
        return
    }

    // Print the loaded configuration
    Printf("Loaded configuration: %+v\n", Config)
    load.LoadConfig = &Config
}

func InitConfigFile() {
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

