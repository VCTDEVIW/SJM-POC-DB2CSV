package project

import (
	_ "fmt"
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
	LoadConfig *JsonConfig
)

