package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/alecthomas/kong"
	"github.com/prasannakumar414/finder/cli"
	"github.com/prasannakumar414/finder/files"
	"github.com/prasannakumar414/finder/models"
)

func main() {
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
		textFiles := make([]string, 0)
		textFilesChan := make(chan []string)
		var wg sync.WaitGroup
		wg.Add(1)
		go files.FileScanner(path, cmd.List.Recursive, textFilesChan, &wg)
		go func() {
			for {
				select {
				case files := <-textFilesChan:
					textFiles = append(textFiles, files...)
				}
			}
		}()
		wg.Wait()
		if len(textFiles) == 0 {
			fmt.Println("No text files found in directory : ", path)
			return
		}
		a := make(chan models.LineCount)
		for _, file := range textFiles {
			wg.Add(1)
			go files.LineCounter(file, a, &wg)
		}
		go func() {
			for {
				select {
				case msg := <-a:
					fmt.Println("File Path : ", msg.FilePath)
					fmt.Println("Line Count : ", msg.LineCount)
					fmt.Println("-----------------------------")
				}
			}
		}()
		wg.Wait()
		close(a)
		fmt.Printf("printed line count for %d files. \n", len(textFiles))
	}
}
