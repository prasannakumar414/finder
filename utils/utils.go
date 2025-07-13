package utils

import (
	"log"
	"os"
	"strings"
)

// Creates Test Directory for testing purposes
func TestEnvironmentSetup() {
	err := os.Mkdir("test-directory", 0777)
	if err != nil {
		log.Fatalf("Error when creating directory for testing : %v", err)
	}
	data := []byte("Hello\nNamaste\nPrasanna\nKumar")
	err = os.WriteFile("test-directory/test-file.txt", data, 0777)
	if err != nil {
		log.Fatalf("Error when creating file : %v", err)
	}
	err = os.Mkdir("test-directory/test-sub-directory", 0777)
	if err != nil {
		log.Fatalf("Error when creating directory for testing : %v", err)
	}
	err = os.WriteFile("test-directory/test-sub-directory/test-file.txt", data, 0777)
	if err != nil {
		log.Fatalf("Error when creating file : %v", err)
	}
}

func GetFileName(filePath string) string {
	pathParts := strings.Split(filePath, "/")
	fileName := pathParts[len(pathParts)-1]
	return fileName
}
