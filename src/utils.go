package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	dir, _ := os.Getwd()

	// var baseDir string
	// var err error

	// baseDir, err = os.UserHomeDir()

	processDirectory(dir)
}

func setBaseDirectory() (string, error) {
	a, err := os.UserHomeDir()
	return a, err
}

func _decompressEpubFile(filePath string, file os.DirEntry) {

	fmt.Println("Path:", filePath)
	fmt.Println("File:", file.Name())

	BASEDIR, err := setBaseDirectory()

	zipReader, err := zip.OpenReader(filePath)

	if err != nil {
		log.Fatal("ePub file `", file.Name(), "` not openable!")
		log.Fatal(err)
	}

	for _, subFile := range zipReader.File {
		fmt.Println("Extracted", subFile.Name)
		dir, _ := strings.CutSuffix(file.Name(), ".epub")
		dir = strings.Join([]string{BASEDIR, "eBookReader", dir}, string(os.PathSeparator))

		saveExtractedFile(dir, subFile)
		fmt.Println(subFile.FileHeader)
	}
}

func saveExtractedFile(baseDirectory string, file *zip.File) error {
	var fileName string = file.Name
	var finalPath = filepath.Join(baseDirectory, fileName)
	fmt.Println(finalPath)

	if file.FileInfo().IsDir() {
		os.MkdirAll(finalPath, os.FileMode(os.O_CREATE))
		return nil
	} else {
		var pathParts []string = strings.Split(finalPath, string(os.PathSeparator))
		pathParts = pathParts[:len(pathParts)-1]

		var newDir = strings.Join(pathParts, string(os.PathSeparator))

		os.MkdirAll(newDir, os.FileMode(os.O_CREATE))
	}

	src, err := file.Open()

	if err != nil {
		log.Print("File opening failed!")
		log.Fatal(err)
		log.Print("\n")
	}

	dst, err := os.Create(finalPath)
	if err != nil {
		log.Print("File creation failed!")
		log.Fatal(err)
		log.Print("\n")
	}

	_, err = io.Copy(dst, src)

	if err != nil {
		log.Print("File copy failed!")
		log.Fatal(err)
		log.Print("\n")
	}
	return nil
}

func decompressEpubFile(filePath string, d os.DirEntry, err error) error {

	if strings.HasSuffix(d.Name(), ".epub") {
		_decompressEpubFile(filePath, d)
	}

	return nil
}

func processDirectory(directory string) error {
	fmt.Println("Walking dir `", directory, "` ...")
	return filepath.WalkDir(directory, decompressEpubFile)
}
