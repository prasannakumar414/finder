package files

import (
	"bufio"
	"io"
	"log"
	"os"
)

func LineCounter(filePath string)  {
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
	
}
