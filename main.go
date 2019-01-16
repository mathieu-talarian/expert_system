package main

import (
	"log"
	"os"
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
	parseFlags()
	if debugF() {
		printError("Debug on")
	}
	if len(argsF()) == 0 {
		log.Fatal("Not enough arguments")
	} else if (len(argsF())) == 2 {
		printError("Too much arguments")
	}
	// open File
	file, err := os.Open(argsF()[0])
	if err != nil {
		printError(err)
	}
	parseFile(file)
}
