package project

import (
	. "fmt"
)

type META_Global struct {
	InitParam bool
	ConfigFilename string
	DebugMode int
	LoadConfig *JsonConfig
}

var (
	InitParam bool = true
	ConfigFilename string = "config.json"
	DebugMode int = 1
)

func Init_LoadMetadata() *META_Global {
	global_InitParam := InitParam

	inst := &META_Global {
		InitParam: global_InitParam,
	}

	return inst
}

func (load *META_Global) Init_ProgramBegin() {
	Println(load.InitParam)
	load.ConfigFilename = ConfigFilename
	load.DebugMode = DebugMode
}

func (load *META_Global) Check_LoadConfig_Rule_1() {
	Println()
	Println(load.LoadConfig.SqlSrv.Query)
	Println(load.LoadConfig.Options.MongoDBEmbedJSON)
	Println()
}

func (load *META_Global) Init_RunTask() {
	load.Init_ProgramBegin()
	load.ReadConfigFile()
	load.Check_LoadConfig_Rule_1()
}


