package main

import (
	"flag"
)

type flags struct {
	Debug bool
	Args  []string
}

var global flags

func parseFlags() flags {
	flag.BoolVar(&global.Debug, "d", false, "set to true for debug display")
	flag.Parse()
	global.Args = flag.Args()
	return global
}

func debugF() bool {
	return global.Debug
}

func argsF() []string {
	return global.Args
}
