package main

import (
	. "fmt"
	. "project/workspace/sjm-poc-db/libs/db-conn-sqlsrv"
	. "project/workspace/sjm-poc-db/libs/db-conn-mongodb"
)

func main() {
	SqlSrv_T()
	MongoDb_T()
	Println()
}

