package main

import (
	. "fmt"
    _"encoding/json"
    _"io/ioutil"
    "net/http"
	. "project/workspace/sjm-poc-db/libs"
)

func init() {
    InitConfigFile()
}

func mongodbJobHandler(w http.ResponseWriter, r *http.Request) {
	proc := Init_LoadMetadata()
	proc.Init_RunTask()

    use_token := proc.GetAccessToken()
    if use_token != "" {
        get_token := r.URL.Query().Get("token")
        if get_token != use_token {
            Fprintf(w, Sprintf("Invalid API access token: ?token=%s", get_token))
        }
        proc.MongoDB_RunQuery()
        Fprintf(w, "Handling MongoDB Job...")
    }

	proc.MongoDB_RunQuery()
	Fprintf(w, "Handling MongoDB Job...")
}

func sqlsrvJobHandler(w http.ResponseWriter, r *http.Request) {
	proc := Init_LoadMetadata()
	proc.Init_RunTask()

    use_token := proc.GetAccessToken()
    if use_token != "" {
        get_token := r.URL.Query().Get("token")
        if get_token != use_token {
            Fprintf(w, Sprintf("Invalid API access token: ?token=%s", get_token))
        }
        proc.MongoDB_RunQuery()
        Fprintf(w, "Handling SQL Server Job...")
    }

	proc.SqlSrv_RunQuery()
    Fprintf(w, "Handling SQL Server Job...")
}

func main() {
    proc := Init_LoadMetadata()
    proc.Init_RunTask()

    //http.HandleFunc("/view", viewConfig())
    mongodb_uri := proc.GetMongoDBUri()
    sqlsrv_uri := proc.GetSqlSrvUri()

    http.HandleFunc(Sprintf("/%s", mongodb_uri), mongodbJobHandler)
    http.HandleFunc(Sprintf("/%s", sqlsrv_uri), sqlsrvJobHandler)

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        msg := "Kong-Vastcom x SJM POC (MicroService Tool for DB2CSV Fetcher)"
        Fprintf(w, msg)
    })

    port := proc.GetApiWebPort()
    Printf("Server is listening on port %d ...", port)
    if err := http.ListenAndServe(Sprintf(":%d", port), nil); err != nil {
        Println("Error starting server:", err)
    }
}