package main

import (
	"ebookreader/utils"
	"os"
)

func main() {
	dir, _ := os.Getwd()

	// var baseDir string
	// var err error

	// baseDir, err = os.UserHomeDir()

	utils.ProcessDirectory(dir)
}
