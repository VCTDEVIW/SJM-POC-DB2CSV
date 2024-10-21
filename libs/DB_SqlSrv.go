package project

import (
	. "fmt"
    "database/sql"
    "log"
    _"io/ioutil"
    "time"
    "context"

    _ "github.com/denisenkom/go-mssqldb"
)

// QueryHandler is a function type that processes each row of the query result.
//type QueryHandler func(row []any)
type QueryHandler func(row []any)

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

