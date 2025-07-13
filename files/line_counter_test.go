package files_test

import (
	"fmt"
	"os"
	"sync"
	"testing"

	"github.com/alecthomas/assert/v2"
	"github.com/prasannakumar414/finder/files"
	"github.com/prasannakumar414/finder/models"
	"github.com/prasannakumar414/finder/utils"
)

func TestLineCounter(t *testing.T) {
	defer os.RemoveAll("test-directory")
	utils.TestEnvironmentSetup()
	var wg sync.WaitGroup
	var lineCountModel models.LineCount
	testChan := make(chan models.LineCount)
	wg.Add(1)
	go files.LineCounter("test-directory/test-file.txt", testChan, &wg)
	select {
	case lineCount := <-testChan:
		fmt.Println("got this linecoutn", lineCount)
		lineCountModel = models.LineCount{
			FilePath:  lineCount.FilePath,
			LineCount: lineCount.LineCount,
		}
	}
	wg.Wait()
	fmt.Println("line count model is ", lineCountModel)
	assert.Equal(t, "test-file.txt", utils.GetFileName(lineCountModel.FilePath))
	assert.Equal(t, 4, lineCountModel.LineCount)
}
