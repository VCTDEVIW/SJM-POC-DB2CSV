package project

import (
	. "fmt"
	"runtime"
	_"runtime/debug"
)

func DebugGoRuntimeInfo() {
	// Print Go version
	Println("Go Version:", runtime.Version())

	// Print number of goroutines
	Println("Number of Goroutines:", runtime.NumGoroutine())

	// Print OS and architecture
	Println("Operating System:", runtime.GOOS)
	Println("Architecture:", runtime.GOARCH)

	// Print memory statistics
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	Println("Memory Allocated (Bytes):", memStats.Alloc)
	Println("Total Allocated Memory (Bytes):", memStats.TotalAlloc)
	Println("System Memory (Bytes):", memStats.Sys)
	Println("Number of GC Cycles:", memStats.NumGC)
}

