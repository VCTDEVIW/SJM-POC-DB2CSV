package project

import (
	"runtime"
	"os"
)

func GetOsPathSlash() string {
	switch runtime.GOOS {
	case "windows":
		return "\\"
	case "linux":
		return "/"
	default:
		return "/"
	}
}

func GetCurrentWorkingDirectory() (string, error) {
    cwd, err := os.Getwd()
    if err != nil {
        return "", err
    }
    return cwd, nil
}

