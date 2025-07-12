package files

import (
	"log"
	"os"
	"strings"
	"sync"
)

func FileScanner(dirPath string, recursive bool, textFilesChan chan []string, wg *sync.WaitGroup) {
	defer wg.Done()
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatalf("Error reading directory: %v", err)
	}
	textFiles := make([]string, 0)
	for _, entry := range entries {
		if entry.IsDir() && recursive {
			directory := dirPath + "/" + entry.Name()
			wg.Add(1)
			go FileScanner(directory, recursive, textFilesChan, wg)
		} else {
			fileNameParts := strings.Split(entry.Name(), ".")
			fileType := fileNameParts[len(fileNameParts)-1]
			if fileType == "txt" {
				filePath := dirPath + "/" + entry.Name()
				textFiles = append(textFiles, filePath)
			}
		}
	}
	textFilesChan <- textFiles
}
