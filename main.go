package main

import (
	. "fmt"
	"sync"
	"log"

	. "project/workspace/sjm-poc-db/libs"
)

func init() {
	Println()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	proc := Init_LoadMetadata()
	proc.Init_RunTask()

	Println()
	log.Println("Task begin...\n")

	go func(wg *sync.WaitGroup) {
		defer func() {
			Println()
			log.Println("SqlSrv task completed.\n")
			wg.Done()
		}()

		Println()
		log.Println("SqlSrv task running...\n")
		proc.SqlSrv_RunQuery()
	}(&wg)

	go func(wg *sync.WaitGroup) {
		defer func() {
			Println()
			log.Println("MongoDB task completed.\n")
			wg.Done()
		}()

		Println()
		log.Println("MongoDB task running...\n")
		proc.MongoDB_RunQuery()
	}(&wg)

	wg.Wait()
	Println()
    log.Println("All task workers finished.")
}

