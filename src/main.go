package main

import (
	"ciphgoer/cli"
	"flag"
	"fmt"
)

var guiFlag bool
var filepath string

func init() {
	flag.BoolVar(&guiFlag, "gui", false, "Run ciphgoer with a graphical user interface. STILL A WIP (Default: false)")
	flag.StringVar(&filepath, "file", "", "File that should be used instead of user input.")
	flag.Parse()
}

func main() {
	if !guiFlag {
		if filepath != "" {
			cli.StartCli(filepath)
		} else {
			cli.StartCli()
		}
	} else {
		fmt.Println("Starting GUI...")
	}
}
