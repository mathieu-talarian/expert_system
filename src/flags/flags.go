package flags

import (
	"flag"
)

type Flags struct {
	Debug bool
}

var Args []string

var Global Flags

func ParseFlags() Flags {
	flag.BoolVar(&Global.Debug, "d", false, "set to true for debug display")
	flag.Parse()
	Args = flag.Args()
	return Global
}

func DebugF() bool {
	return Global.Debug
}

func GetArgs() []string {
	return Args
}
