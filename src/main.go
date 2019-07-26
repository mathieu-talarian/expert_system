package main

import (
	"github.com/mathmoul/expert_system/src/esfmt"
	"github.com/mathmoul/expert_system/src/filehandler"
	"github.com/mathmoul/expert_system/src/flags"
	"github.com/mathmoul/expert_system/src/parser"
)

func main() {
	flags.ParseFlags()
	if flags.DebugF() {
		esfmt.LogAndExit("Debug on")
	}

	if len(flags.Args) == 0 {
		esfmt.LogAndExit("Not enough arguments")
	} else if (len(flags.Args)) == 2 {
		esfmt.LogAndExit("Too much arguments")
	}

	if err := filehandler.OpenFile(flags.Args[0]); err != nil {
		esfmt.LogAndExit(err.Error())
	}

	if _, err := parser.Start(filehandler.OpenedFile.ToString()); err != nil {
		esfmt.LogAndExit(err.Error())
	}

	if error := filehandler.OpenedFile.Close(); error != nil {
		esfmt.LogAndExit(error.Error())
	}
}
