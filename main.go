package main

import (
	. "fmt"

	. "project/workspace/sjm-poc-db/libs"
)

func init() {
	Println()
}

func main() {
	proc := Init_LoadMetadata()
	proc.Init_RunTask()
}

