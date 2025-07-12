package files

import (
	"bufio"
	"io"
	"log"
	"os"
	"sync"

	"github.com/prasannakumar414/finder/models"
)

func LineCounter(filePath string, a chan (models.LineCount), wg *sync.WaitGroup) {
	defer wg.Done()
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error when opening file: %v", err)
	}
	reader := bufio.NewReader(file)
	lineCount := 0
	for {
		_, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("Error when reading line: %v", err)
		}
		lineCount++
	}
	a <- models.LineCount{FilePath: filePath, LineCount: lineCount}
}
