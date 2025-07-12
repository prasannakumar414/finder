package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/alecthomas/kong"
	"github.com/prasannakumar414/finder/cli"
	"github.com/prasannakumar414/finder/files"
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
		textFiles := files.FileScanner(path, cmd.List.Recursive)
		if len(textFiles) == 0 {
			fmt.Println("No text files found in directory : ", path)
			return
		}
		for _, file := range textFiles {
			fmt.Println(file)
		}
	}
}
