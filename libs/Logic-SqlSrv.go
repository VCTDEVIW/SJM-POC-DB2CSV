package project

import (
	. "fmt"
    "log"
    _ "encoding/json"
    _"io/ioutil"

    "encoding/csv"
	"os"

    _ "github.com/denisenkom/go-mssqldb"
)

func SqlSrv_Job1(rows []any) {
    for _, value := range rows {
        //Printf("%v\t", *(value.(*any))) // Print each value in the row
        scanResult := Sprintf( "%v", *(value.(*any)) )
        Printf(scanResult)
    }
    Println()
}

func (load *META_Global) SqlSrv_Test() {
    // func SqlSrv_Conn(dsn string, port int, username string, password string, dbName string) (*sql.DB, error)
    sqlsrv_username := load.LoadConfig.SqlSrv.Username
    sqlsrv_password := load.LoadConfig.SqlSrv.Password
    sqlsrv_dsn := load.LoadConfig.SqlSrv.Host
    sqlsrv_port := load.LoadConfig.SqlSrv.Port
    sqlsrv_database := load.LoadConfig.SqlSrv.DBName
    //sqlsrv_table := load.LoadConfig.SqlSrv.Table

    db, err := SqlSrv_Conn(sqlsrv_dsn, sqlsrv_port, sqlsrv_username, sqlsrv_password, sqlsrv_database)
    if err != nil {
        log.Fatal("Error connecting to the database:", err)
    }
    defer db.Close()

    var sqlsrv_query string

	/*
    sqlsrv_query = `use test_db; select top(10) [rid],[text1],[text2],[updated_at] from table1;`
    sqlsrv_query = `use test_db; select top(10) * from table1 order by [text1] desc;`
	*/

	sqlsrv_query = load.LoadConfig.SqlSrv.Query

    // Execute the query and pass the callback
    if err := SqlSrv_Read(db, sqlsrv_query, SqlSrv_Job1); err != nil {
        log.Fatal(err)
    }
}

func (load *META_Global) SqlSrv_RunQuery() {
    // func SqlSrv_Conn(dsn string, port int, username string, password string, dbName string) (*sql.DB, error)
    sqlsrv_username := load.LoadConfig.SqlSrv.Username
    sqlsrv_password := load.LoadConfig.SqlSrv.Password
    sqlsrv_dsn := load.LoadConfig.SqlSrv.Host
    sqlsrv_port := load.LoadConfig.SqlSrv.Port
    sqlsrv_database := load.LoadConfig.SqlSrv.DBName
    //sqlsrv_table := load.LoadConfig.SqlSrv.Table

    db, err := SqlSrv_Conn(sqlsrv_dsn, sqlsrv_port, sqlsrv_username, sqlsrv_password, sqlsrv_database)
    if err != nil {
        log.Fatal("Error connecting to the database:", err)
    }
    defer db.Close()


    outputFilename := SqlSrv_CsvFilename

    // Create a new CSV file
	file, err := createFile(outputFilename)
	if err != nil {
		Println("Error creating file:", err)
		return
	}
	defer file.Close()
    
    // Create a CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

    /*
    records := [][]string{
		{"Name", "Age", "City"},
		{"张三", "30", "北京"},
	}
    */
    var records [][]string


	/*
    sqlsrv_query = `use test_db; select top(10) [rid],[text1],[text2],[updated_at] from table1;`
    sqlsrv_query = `use test_db; select top(10) * from table1 order by [text1] desc;`
	*/
    var sqlsrv_query string
	sqlsrv_query = load.LoadConfig.SqlSrv.Query

    // Execute the query
    rows, err := db.Query(sqlsrv_query)
    if err != nil {
        log.Println("error executing query: %v", err)
    }
    defer rows.Close()

    // Write column headers
    columns, err := rows.Columns()
    if err != nil {
        log.Fatalf("Error getting columns: %v", err)
    }
    if err := writer.Write(columns); err != nil {
        log.Fatalf("Error writing headers to CSV: %v", err)
    }

    // Iterate through the result set
    for rows.Next() {
        // Create a slice to hold the values
        values := make([]interface{}, len(columns))
        valuePtrs := make([]interface{}, len(columns))
        for i := range values {
            valuePtrs[i] = &values[i]
        }

        // Scan the row into the pointers
        if err := rows.Scan(valuePtrs...); err != nil {
            log.Fatalf("Error scanning row: %v", err)
        }

        // Convert the values to strings
        stringValues := make([]string, len(columns))
        for i, val := range values {
            if val != nil {
                stringValues[i] = Sprintf("'%v'", val)
            } else {
                stringValues[i] = ""
            }
        }

        // Write the row to CSV
        if err := writer.Write(stringValues); err != nil {
            log.Fatalf("Error writing row to CSV: %v", err)
        }
    }

    // Check for errors during iteration
    if err := rows.Err(); err != nil {
        log.Println("error iterating over rows: %v", err)
    }


    // Write records to CSV
	if err := writeRecords(writer, records); err != nil {
		Println("Error writing records to file:", err)
		return
	}

    Println("CSV file created from SqlSrv successfully.")
}

// createFile opens the file for writing, creating it if it does not exist or truncating it if it does.
func createFile(filename string) (*os.File, error) {
	// Open the file with O_RDWR (read/write), O_CREATE (create if not exists), and O_TRUNC (truncate to zero if exists)
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return nil, err
	}
	// Write BOM for UTF-8
	file.WriteString("\xEF\xBB\xBF")
	return file, nil
}

// writeRecords writes the given records to the CSV writer.
func writeRecords(writer *csv.Writer, records [][]string) error {
	for _, record := range records {
		if err := writer.Write(record); err != nil {
			return err
		}
        Println(record)
	}
	return nil
}

