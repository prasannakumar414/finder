package files_test

import (
	"os"
	"sync"
	"testing"

	"github.com/alecthomas/assert/v2"
	"github.com/prasannakumar414/finder/files"
	"github.com/prasannakumar414/finder/utils"
)

func TestWordCounter(t *testing.T) {
	defer os.RemoveAll("test-directory")
	utils.TestEnvironmentSetup()
	var wg sync.WaitGroup
	wordMap := make(map[string]int)
	testChan := make(chan map[string]int)
	wg.Add(1)
	go files.WordCounter("test-directory/test-file.txt", testChan, &wg)
	select {
	case words := <-testChan:
		for k, v := range words {
			wordMap[k] += v
		}
	}
	wg.Wait()
	expected := make(map[string]int)
	expected["Hello"] = 2
	expected["Namaste"] = 1
	expected["Prasanna"] = 1
	expected["Kumar"] = 1
	assert.Equal(t, expected, wordMap)
}
