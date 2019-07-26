package flags

import (
	"flag"
)

// Flags is q struct of our wanted flags
type Flags struct {
	Debug bool
}

// Args array of args
var Args []string

// Global is result of flags parsing
var Global Flags

// ParseFlags uses flag package to parse flags
func ParseFlags() Flags {
	flag.BoolVar(&Global.Debug, "d", false, "set to true for debug display")
	flag.Parse()
	Args = flag.Args()
	return Global
}

// DebugF return true if debug enabled
func DebugF() bool {
	return Global.Debug
}

// GetArgs getter for args
func GetArgs() []string {
	return Args
}
