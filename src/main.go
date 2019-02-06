package main

import (
	"log"
	"os"

	"github.com/mathmoul/expert_system/src/flags"
	"github.com/mathmoul/expert_system/src/parser"
)

var color = map[string]string{
	"ANSI_COLOR_RED":     "\x1b[31m",
	"ANSI_COLOR_GREEN":   "\x1b[32m",
	"ANSI_COLOR_YELLOW":  "\x1b[33m",
	"ANSI_COLOR_BLUE":    "\x1b[34m",
	"ANSI_COLOR_MAGENTA": "\x1b[35m",
	"ANSI_COLOR_CYAN":    "\x1b[36m",
	"ANSI_COLOR_RESET":   "\x1b[0m",
}

func printError(v interface{}) {
	log.Println(color["ANSI_COLOR_RED"], "Error =>", color["ANSI_COLOR_RESET"], v)
}

func main() {
	flags.ParseFlags()
	if flags.DebugF() {
		printError("Debug on")
	}
	if len(flags.Args) == 0 {
		log.Fatal("Not enough arguments")
	} else if (len(flags.Args)) == 2 {
		printError("Too much arguments")
	}
	// open File
	file, err := os.Open(flags.Args[0])
	if err != nil {
		printError(err)
	}
	if err := parser.ParseFile(file); err != nil {
		printError(err)
	}
}
