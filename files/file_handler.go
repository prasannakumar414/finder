package files

import (
	"sync"

	"github.com/prasannakumar414/finder/models"
)

// one place that sends data to our main in required type
func FilesDirectoryHandler(
	path string,
	recursive bool,
	wordCount bool,
	lineCountChan chan models.LineCount,
	wordCountChan chan map[string]int,
	textFilesChan chan []string,
	wg *sync.WaitGroup,
) {
	textFiles := make([]string, 0)
	wg.Add(1)
	go FileScanner(path, recursive, textFilesChan, wg)
	go func() {
		for {
			select {
			case filePaths := <-textFilesChan:
				textFiles = append(textFiles, filePaths...)
				for _, filePath := range filePaths {
					wg.Add(1)
					go LineCounter(filePath, lineCountChan, wg)
					if wordCount {
						wg.Add(1)
						go WordCounter(filePath, wordCountChan, wg)
					}
				}
			}
		}
	}()
}
