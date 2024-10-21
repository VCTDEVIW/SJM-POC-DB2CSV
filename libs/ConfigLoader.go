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

