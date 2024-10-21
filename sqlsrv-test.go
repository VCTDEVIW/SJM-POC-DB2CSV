package main

import (
    "database/sql"
    . "fmt"
    "log"
    "encoding/json"
    _"io/ioutil"
    "os"
    "runtime"
    "time"
    "context"

    _ "github.com/denisenkom/go-mssqldb"
)

// Define a struct to match the JSON structure
type Config struct {
    Configuration struct {
        OutputPath string `json:"outputPath"`
    } `json:"configuration"`
    SQLSrv struct {
        SQLDBHost     string `json:"sql-db-host"`
        SQLDBPort     int    `json:"sql-db-port"`
        SQLDBUsername string `json:"sql-db-username"`
        SQLDBPassword string `json:"sql-db-password"`
        SQLDBDbname   string `json:"sql-db-dbname"`
        SQLDBStname   string `json:"sql-db-stname"`
    } `json:"sqlsrv"`
    MongoDB interface{} `json:"mongodb"` // Empty object
    Misc    interface{} `json:"misc"`    // Empty object
}

// SqlSrv_Conn(...) establishes a connection to the SQL Server database.
// Connection timeout is set to 5 sec.
func SqlSrv_Conn(dsn string, port int, username string, password string, dbName string) (*sql.DB, error) {
    connString := Sprintf("sqlserver://%s:%s@%s:%d?database=%s;connection timeout=3", username, password, dsn, port, dbName)
    db, err := sql.Open("sqlserver", connString)

    SqlSrv_Ping(db)

    if err != nil {
        return nil, err
    }
    return db, nil
}

func SqlSrv_Ping(db *sql.DB) {
    VAR_timeoutSec := 3
    timeoutSec := time.Duration(VAR_timeoutSec)
    ctx, cancel := context.WithTimeout(context.Background(), timeoutSec *time.Second)
    defer cancel()

    if err := db.PingContext(ctx); err != nil {
        log.Fatalf("Error connecting to the database: %s", err.Error())
    }

    Println("Connected to the database!")
}

func SqlSrv_Read(db *sql.DB, sqlQuery string, callback_flowFunc func(*sql.Rows) error) {
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
    VAR_pathPrefix := VAR_thisPath + GetOsPathSlash()
    VAR_configFile := VAR_pathPrefix + "sample_config.json"

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

    db, err := SqlSrv_Conn(sqlsrv_dsn, sqlsrv_port, sqlsrv_username, sqlsrv_password, sqlsrv_database)
    if err != nil {
        log.Fatal("Error connecting to the database:", err)
    }
    defer db.Close()

    
    
}

