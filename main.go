package main

import (
	_ "fmt"
	_ "project/workspace/sjm-poc-db/libs/configFile"
	. "project/workspace/sjm-poc-db/libs/db-conn-sqlsrv"
)

func main() {
	Sqlsrv_T()
}

