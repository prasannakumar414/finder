package files

import (
	"log"
	"os"
	"strings"
)

func FileScanner(dirPath string, recursive bool) []string {
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatalf("Error reading directory: %v", err)
	}
	textFiles := make([]string, 0)
	for _, entry := range entries {
		if entry.IsDir() && recursive {
			directory := dirPath + "/" + entry.Name()
			textFiles = append(textFiles, FileScanner(directory, recursive)...)
		} else {
			fileNameParts := strings.Split(entry.Name(), ".")
			fileType := fileNameParts[len(fileNameParts)-1]
			if fileType == "txt" {
				filePath := dirPath + "/" + entry.Name()
				textFiles = append(textFiles, filePath)
			}
		}
	}
	return textFiles
}
