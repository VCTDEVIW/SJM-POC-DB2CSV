package project

import (
	. "fmt"
)

func Init_LoadMetadata() *META_Global {
	global_InitParam := InitParam

	inst := &META_Global {
		InitParam: global_InitParam,
	}

	return inst
}

func (load *META_Global) Init_ProgramBegin() {
	load.ConfigFilename = ConfigFilename
	load.DebugMode = DebugMode
	DebugGoRuntimeInfo()
	Println()
	Println(load.InitParam)
}

func (load *META_Global) Check_LoadConfig_Rule_1() {
	Println()
	Println("Registered SQL Server query:")
	Println(load.LoadConfig.SqlSrv.Query)
	Println("\nRegistered MongoDB query:")
	Println(load.LoadConfig.Options.MongoDBEmbedJSON)
	Println()
}

func (load *META_Global) PatchFetchCsvOutputFilenames() {
	SqlSrv_CsvFilename = load.LoadConfig.Options.SqlSrvOutputFile
	MongoDB_CsvFilename = load.LoadConfig.Options.MongoDBOutputFile
}

func (load *META_Global) GetMongoDBUri() string {
	return load.LoadConfig.Options.API_MongoDBJob_URI
}

func (load *META_Global) GetSqlSrvUri() string {
	return load.LoadConfig.Options.API_SqlJob_URI
}

func (load *META_Global) GetApiWebPort() int {
	return load.LoadConfig.Options.API_WebPort
}

func (load *META_Global) GetAccessToken() string {
	return load.LoadConfig.Options.API_AccessGetToken
}

func (load *META_Global) Init_RunTask() {
	load.Init_ProgramBegin()
	load.ReadConfigFile()
	load.Check_LoadConfig_Rule_1()
	load.PatchFetchCsvOutputFilenames()
	//load.SqlSrv_Test()
	//load.MongoDB_Test()
	/*
	load.SqlSrv_RunQuery()
	load.MongoDB_RunQuery()
	*/
}

