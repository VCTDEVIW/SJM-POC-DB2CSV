package main

import (
	. "fmt"
    _"encoding/json"
    _"io/ioutil"
    "net/http"
	. "project/workspace/sjm-poc-db/libs"
)

func mongodbJobHandler(w http.ResponseWriter, r *http.Request) {
	proc := Init_LoadMetadata()
	proc.Init_RunTask()
	proc.MongoDB_RunQuery()
	Fprintf(w, "Handling MongoDB Job...")
}

func sqlsrvJobHandler(w http.ResponseWriter, r *http.Request) {
	proc := Init_LoadMetadata()
	proc.Init_RunTask()
	proc.SqlSrv_RunQuery()
    Fprintf(w, "Handling SQL Server Job...")
}

func main() {
    //http.HandleFunc("/view", viewConfig())
    http.HandleFunc("/mongodb-job", mongodbJobHandler)
    http.HandleFunc("/sqlsrv-job", sqlsrvJobHandler)

    Println("Server is listening on port 8080...")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        Println("Error starting server:", err)
    }
}