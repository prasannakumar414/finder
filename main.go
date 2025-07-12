package main

import (
	"fmt"

	"github.com/alecthomas/kong"
	"github.com/prasannakumar414/finder/cli"
)

func main() {
	ctx := kong.Parse(&cli.Command)
	switch ctx.Command() {
	case "list <path>":
		fmt.Println("list command is working")
	}
}
