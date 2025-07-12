package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/alecthomas/kong"
	"github.com/prasannakumar414/finder/cli"
)

func main() {
	var cmd cli.Command
	ctx := kong.Parse(&cmd)
	switch strings.Split(ctx.Command(), " ")[0] {
	case "list":
		fmt.Println("list command is working")
		fmt.Println(cmd.List.Path)
		fmt.Println(cmd.List.Recursive)
		dir, _ := os.Getwd()
		// Print the current working directory
		fmt.Println("Current working directory:", dir)
	}
}
