package utils

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// setBaseDirectory sets the base directory for the application.
// This is the directory at which extracted information is stored.
func SetBaseDirectory() (string, error) {
	a, err := os.UserHomeDir()
	return a, err
}

// _decompressEpubFile decompresses the .ePub file and stores the contents appropriately.
func _DecompressEpubFile(filePath string, file os.DirEntry) {

	fmt.Println("Path:", filePath)
	fmt.Println("File:", file.Name())

	BASEDIR, _ := SetBaseDirectory()

	zipReader, err := zip.OpenReader(filePath)

	if err != nil {
		log.Fatal("ePub file `", file.Name(), "` not openable!")
		log.Fatal(err)
	}

	for _, subFile := range zipReader.File {
		fmt.Println("Extracted", subFile.Name)
		dir, _ := strings.CutSuffix(file.Name(), ".epub")
		dir = strings.Join([]string{BASEDIR, "eBookReader", dir}, string(os.PathSeparator))

		SaveExtractedFile(dir, subFile)
		fmt.Println(subFile.FileHeader)
	}
}

// SaveExtractedFile saves the files extracted from the .ePub file in the correct directories.
func SaveExtractedFile(baseDirectory string, file *zip.File) error {
	var fileName string = file.Name
	var finalPath = filepath.Join(baseDirectory, fileName)
	fmt.Println(finalPath)

	if file.FileInfo().IsDir() {
		os.MkdirAll(finalPath, os.FileMode(os.O_CREATE))
		return nil
	} else {
		_, err := os.Open(finalPath)

		if !os.IsExist(err) {
			var pathParts []string = strings.Split(finalPath, string(os.PathSeparator))
			pathParts = pathParts[:len(pathParts)-1]

			var newDir = strings.Join(pathParts, string(os.PathSeparator))

			os.MkdirAll(newDir, os.FileMode(os.O_CREATE))
		} else {
			log.Println("File found in path!")
			return nil
		}
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

// decompressEpubFile decompresses the .ePub file and stores the contents appropriately.
// To be used with os directory walking process.
func DecompressEpubFile(filePath string, d os.DirEntry, err error) error {

	if strings.HasSuffix(d.Name(), ".epub") {
		_DecompressEpubFile(filePath, d)
	}

	return nil
}

func ProcessDirectory(directory string) error {
	fmt.Println("Walking dir `", directory, "` ...")
	return filepath.WalkDir(directory, DecompressEpubFile)
}
