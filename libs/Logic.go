package project

import (
	. "fmt"
    "log"
    _ "encoding/json"
    _"io/ioutil"

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

