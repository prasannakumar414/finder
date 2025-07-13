package files_test

import (
	"fmt"
	"os"
	"sync"
	"testing"

	"github.com/alecthomas/assert/v2"
	"github.com/prasannakumar414/finder/files"
	"github.com/prasannakumar414/finder/utils"
)

func TestFileScanner(t *testing.T) {
	defer os.RemoveAll("test-directory")
	utils.TestEnvironmentSetup()
	var wg sync.WaitGroup
	testChan := make(chan []string)
	testFiles := make([]string, 0)
	wg.Add(1)
	go files.FileScanner("test-directory", true, testChan, &wg)
	go func() {
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
	for _, testFilePath := range testFiles {
		fileName := utils.GetFileName(testFilePath)
		assert.Equal(t, "test-file.txt", fileName)
	}
}
