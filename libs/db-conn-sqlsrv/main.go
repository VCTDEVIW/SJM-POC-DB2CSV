package project

import (
    "database/sql"
    . "fmt"
    "log"
    "encoding/json"
    _"io/ioutil"
    "os"
    "runtime"

    _ "github.com/denisenkom/go-mssqldb"
)

// getDBConnection establishes a connection to the SQL Server database.
func getDBConnection(dsn string, port int, username string, password string, dbName string) (*sql.DB, error) {
    connString := Sprintf("sqlserver://%s:%s@%s:%d?database=%s;sslmode=disable", username, password, dsn, port, dbName)
    db, err := sql.Open("sqlserver", connString)
    if err != nil {
        return nil, err
    }
    return db, nil
}


func read(db *sql.DB, sqlQuery string, callback_flowFunc func(*sql.Rows) error) {
    rows, err := db.Query(sqlQuery)
    if err != nil {
        log.Fatal("Error reading records: ", err.Error())
    }
    defer rows.Close()

    for rows.Next() {
        if err := callback_flowFunc(rows); err != nil {
            log.Fatal("Error processing row: ", err.Error())
        }
    }

    if err := rows.Err(); err != nil {
        log.Fatal("Error iterating rows: ", err.Error())
    }
}

func sqlsrv_job1(rows *sql.Rows) (string, error) {
    return "", nil
}


func GetCurrentWorkingDirectory() (string, error) {
    cwd, err := os.Getwd()
    if err != nil {
        return "", err
    }
    return cwd, nil
}

func GetOsPathSlash() string {
	switch runtime.GOOS {
	case "windows":
		return "\\"
	case "linux":
		return "/"
	default:
		return "/"
	}
}

func main() {
    // workspace\sjm-poc-db
    VAR_thisPath, _ := GetCurrentWorkingDirectory()
    VAR_pathPrefix := VAR_thisPath + GetOsPathSlash() + "workspace" + GetOsPathSlash() + "sjm-poc-db" + GetOsPathSlash()
    VAR_configFile := VAR_pathPrefix + "simple_config.json"

    file, err := os.Open(VAR_configFile)
    if err != nil {
        Println("Error opening file:", err)
        return
    }
    defer file.Close()

    // Create a variable of type Config
    var config Config

    // Decode the JSON into the variable
    decoder := json.NewDecoder(file)
    err = decoder.Decode(&config)
    if err != nil {
        Println("Error decoding JSON:", err)
        return
    }

    // func getDBConnection(dsn string, port int, username string, password string, dbName string) (*sql.DB, error)
    sqlsrv_username := config.SQLSrv.SQLDBUsername
    sqlsrv_password := config.SQLSrv.SQLDBPassword
    sqlsrv_dsn := config.SQLSrv.SQLDBHost
    sqlsrv_port := config.SQLSrv.SQLDBPort
    sqlsrv_database := config.SQLSrv.SQLDBDbname
    //sqlsrv_stname := config.SQLSrv.SQLDBStname

    db, err := getDBConnection(sqlsrv_dsn, sqlsrv_port, sqlsrv_username, sqlsrv_password, sqlsrv_database)
    if err != nil {
        log.Fatal("Error connecting to the database:", err)
    }
    defer db.Close()

    // Test the connection
    err = db.Ping()
    if err != nil {
        log.Fatalf("Error connecting to the database: %s", err.Error())
    }
    Println("Connected to the database!")

}

