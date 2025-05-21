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
	// getFocusDirectories()
	var num int = 10
	var address *int = &num

	fmt.Println("Value using pointer:", *(address))
	fmt.Println("Value using variable:", num)

	dir, _ := os.Getwd()

	processDirectory(dir)
}

func getFocusDirectories() {
	// homeDirectory, _ := os.UserHomeDir()
	// r, _ := os.ReadDir(homeDirectory)

	// fmt.Println(r[1].Name())
	var dir string = "/files"
	if []rune(dir)[0] == '/' {
		fmt.Println("Starts with slash")
	}
	extractFileNamesInDirectory(dir)
}

func extractEpubFileName(file os.DirEntry) (name string) {
	// var file zip.ReadCloser = zip.OpenReader(fileName)

	name = file.Name()

	return name
}

func extractEpubFileProperties(file os.DirEntry) map[string]string {
	var properties = make(map[string]string)

	properties["fileName"] = extractEpubFileName(file)

	return properties
}

func extractFileNamesInDirectory(directory string) {
	var currentDirectory string
	var err error

	// Establish present directory
	currentDirectory, err = os.Getwd()

	fmt.Println(currentDirectory)

	if err != nil {
		log.Fatal(err, "encountered")
	}
	// Change working directory to directory variable
	os.Chdir(directory)

	// Extract files from present directory
	var files []os.DirEntry

	files, err = os.ReadDir(currentDirectory + directory)

	if err != nil {
		log.Fatal("Encountered ", err)
	}

	for i := 0; i < len(files); i++ {
		file := files[i]
		fileName := extractEpubFileName(file)
		fmt.Println(file.Info())
		// decompressEpubFile(file)

		log.Print(fileName)
	}

	// Restore directory to original directory
	os.Chdir(currentDirectory)
	return
}

func decompressEpubFile(filePath string, file os.DirEntry) {
	// var zipReader *zip.ReadCloser
	// var err error

	fmt.Println("Path:", filePath)
	fmt.Println("File:", file.Name())

	zipReader, err := zip.OpenReader(filePath)

	if err != nil {
		log.Fatal("ePub file `", file.Name(), "` not openable!")
		log.Fatal(err)
	}

	for _, subFile := range zipReader.File {
		fmt.Println("Extracted", subFile.Name)
		dir, _ := strings.CutSuffix(filePath, ".epub")

		saveExtractedFile(dir, subFile)
		fmt.Println(subFile.FileHeader)
	}
}

func saveExtractedFile(directory string, file *zip.File) error {
	var fileName string = file.Name
	var finalPath = filepath.Join(directory, fileName)
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

	// src, err := os.Open(fileName)
	if err != nil {
		log.Print("File opening failed!")
		log.Fatal(err)
		log.Print("\n")
	}

	// var pathParts []string = strings.Split(finalPath, string(os.PathSeparator))
	// pathParts = pathParts[:len(pathParts)-1]

	// var newDir = strings.Join(pathParts, string(os.PathSeparator))

	// os.MkdirAll(newDir, os.FileMode(os.O_CREATE))

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

func walkDecompressEpubFile(path string, d os.DirEntry, err error) error {
	// var zipReader *zip.ReadCloser
	// var err error

	// if filepath.Ext(d.Name()) == ".epub" {
	// 	decompressEpubFile(path, d)
	// }

	if strings.HasSuffix(d.Name(), ".epub") {
		decompressEpubFile(path, d)
	}

	return nil
}

func processDirectory(directory string) {
	fmt.Println("Walking dir `", directory, "` ...")
	filepath.WalkDir(directory, walkDecompressEpubFile)
}
