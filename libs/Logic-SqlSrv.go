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

    // Execute the query and pass the callback
    if err := SqlSrv_Read(db, sqlsrv_query, SqlSrv_Job2_GenCsv); err != nil {
        log.Fatal(err)
    }

    // Write records to CSV
	if err := writeRecords(writer, records); err != nil {
		Println("Error writing records to file:", err)
		return
	}

    Println("CSV file created successfully.")
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

func SqlSrv_Job2_GenCsv(rows []any) {
    //fixInitSlice := []string{" "}

    for _, value := range rows {
        //Printf("%v\t", *(value.(*any))) // Print each value in the row
        
        //scanResult := Sprintf( "%v\t", *(value.(*any)) )
        scanResult := Sprintf( "%v", *(value.(*any)) )
        Printf(scanResult)
    }
    Println()
}

