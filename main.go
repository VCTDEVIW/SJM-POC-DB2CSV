package main

import (
	. "fmt"

	. "project/workspace/sjm-poc-db/libs/configFile"
	. "project/workspace/sjm-poc-db/libs/db-conn-sqlsrv"
	. "project/workspace/sjm-poc-db/libs/db-conn-mongodb"
)

func init() {
	Config_T()
	SqlSrv_T()
	MongoDb_T()
}

func main() {
	Println()



	MockJson()

}

