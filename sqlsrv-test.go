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

// QueryHandler is a function type that processes each row of the query result.
type QueryHandler func(row []any)

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
    MongoDB any `json:"mongodb"` // Empty object
    Misc    any `json:"misc"`    // Empty object
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

func SqlSrv_Read(db *sql.DB, query string, handler QueryHandler) error {
    rows, err := db.Query(query)
    if err != nil {
        return Errorf("error executing query: %v", err)
    }
    defer rows.Close()

    // Get column names and prepare a slice for the row data
    columns, err := rows.Columns()
    if err != nil {
        return Errorf("error getting columns: %v", err)
    }

    // Create a slice to hold the values for each row
    values := make([]any, len(columns))

    // Iterate through the result set
    for rows.Next() {
        // Prepare the slice for the current row
        for i := range values {
            values[i] = new(any) // Allocate memory for each column
        }

        // Scan the row into the values slice
        if err := rows.Scan(values...); err != nil {
            return Errorf("error scanning row: %v", err)
        }

        // Call the handler for each row
        handler(values)
    }

    // Check for errors during iteration
    if err := rows.Err(); err != nil {
        return Errorf("error iterating over rows: %v", err)
    }

    return nil
}

func SqlSrv_Job1(rows []any) {
    for _, value := range rows {
        //Printf("%v\t", *(value.(*any))) // Print each value in the row
        scanResult := Sprintf( "%v", *(value.(*any)) )
        Printf(scanResult)
    }
    Println()
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

    var sqlsrv_query string

    sqlsrv_query = `use test_db; select top(10) [rid],[text1],[text2],[updated_at] from table1;`
    sqlsrv_query = `use test_db; select top(10) * from table1 order by [text1] desc;`

    // Execute the query and pass the callback
    if err := SqlSrv_Read(db, sqlsrv_query, SqlSrv_Job1); err != nil {
        log.Fatal(err)
    }
}

