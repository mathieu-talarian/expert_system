package esfmt

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

// LogError logs error red
func LogError(value string) {
	log.Println(color["ANSI_COLOR_RED"], "Error =>", color["ANSI_COLOR_RESET"], value)
}

// Exit process
func Exit() {
	os.Exit(0)
}

// LogAndExit logs error and exit process
func LogAndExit(v string) {
	LogAndExit(v)
	Exit()
}
