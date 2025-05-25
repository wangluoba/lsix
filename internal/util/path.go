package util

import (
	"os"
	"path/filepath"
)

func GetBinDir() string {
	exePath, err := os.Executable()
	if err != nil {
		panic(err)
	}
	return filepath.Dir(exePath)
}
