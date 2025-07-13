package files_test

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"testing"

	"github.com/alecthomas/assert/v2"
	"github.com/prasannakumar414/finder/files"
)

func TestFileScanner(t *testing.T) {
	defer os.RemoveAll("test-directory")
	err := os.Mkdir("test-directory", 0777)
	if err != nil {
		log.Fatalf("Error when creating directory for testing : %v", err)
	}
	data := []byte("Hello\nNamaste\nPrasanna\nKumar")
	err = os.WriteFile("test-directory/test-file.txt", data ,0777)
	if err != nil {
		log.Fatalf("Error when creating file : %v", err)
	}
	err = os.Mkdir("test-directory/test-sub-directory", 0777)
	if err != nil {
		log.Fatalf("Error when creating directory for testing : %v", err)
	}
	err = os.WriteFile("test-directory/test-sub-directory/test-file.txt", data ,0777)
	if err != nil {
		log.Fatalf("Error when creating file : %v", err)
	}
	var wg sync.WaitGroup
	testChan := make(chan []string)
	testFiles := make([]string, 0)
	wg.Add(1)
	go files.FileScanner("test-directory", true, testChan, &wg)
	go func () {
		for {
			select {
			case file := <-testChan:
				testFiles = append(testFiles, file...)
			}
		}
	}()
	wg.Wait()
	// should count the 2 files one in directory and subdirectory
	fmt.Println(testFiles)
	assert.Equal(t, 2, len(testFiles))
	// should match name of file
	for _, testFile := range testFiles {
		pathParts := strings.Split(testFile, "/")
		fileName := pathParts[len(pathParts) - 1]
		assert.Equal(t, "test-file.txt", fileName)
	}
}