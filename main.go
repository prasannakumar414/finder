package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/alecthomas/kong"
	"github.com/prasannakumar414/finder/cli"
	"github.com/prasannakumar414/finder/files"
	"github.com/prasannakumar414/finder/models"
)

func main() {
	start := time.Now()
	var cmd cli.Command
	ctx := kong.Parse(&cmd)
	switch strings.Split(ctx.Command(), " ")[0] {
	case "list":
		path := cmd.List.Path
		if path == "" {
			var err error
			path, err = os.Getwd()
			if err != nil {
				log.Fatalf("Error when getting current directory: %v", err)
			}
		}
		var wg sync.WaitGroup
		stopChan := make(chan bool)
		lineCountChan := make(chan models.LineCount)
		wordCountMapChan := make(chan map[string]int)
		textFilesChan := make(chan []string)
		fileCount := 0
		wordCountsMap := make(map[string]int)
		files.FilesDirectoryHandler(
			path,
			cmd.List.Recursive,
			lineCountChan,
			wordCountMapChan,
			textFilesChan,
			&wg,
		)
		go func() {
			for {
				select {
				case lineCountModel := <-lineCountChan:
					pathParts := strings.Split(lineCountModel.FilePath, "/")
					fileName := pathParts[len(pathParts)-1]
					fmt.Println("File Path : ", lineCountModel.FilePath)
					fmt.Println("File Name : ", fileName)
					fmt.Println("Line Count : ", lineCountModel.LineCount)
					fmt.Println("-----------------------------")
					fileCount++
				case wordMap := <-wordCountMapChan:
					for k, v := range wordMap {
						wordCountsMap[k] += v
					}
				case <-stopChan:
					return
				}
			}
		}()
		wg.Wait()
		stopChan<-true
		close(textFilesChan)
		close(lineCountChan)
		close(wordCountMapChan)
		close(stopChan)
		if fileCount == 0 {
			fmt.Println("no files in the directory : ", path)
		} else {
			fmt.Println(fileCount)
			wordCounts := files.GetMostFrequentWordsFromMap(wordCountsMap)
			fmt.Println("Top ten most frequent words are : ")
			for _, word := range wordCounts {
				fmt.Println("word : ", word.Word)
				fmt.Println("count : ", word.Count)
				fmt.Println("--------------------------------")
			}
			fmt.Printf("scanned %d files in %d ms. \n", fileCount, (time.Now().Sub(start)).Milliseconds())
		}
	}
}
