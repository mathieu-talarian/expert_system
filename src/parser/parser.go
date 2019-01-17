package parser

import (
	"bytes"
	"os"

	lane "gopkg.in/oleiade/lane.v1"
)

// Parser struct
type Parser struct {
	index int

	// rules ParseResult
	token map[string]string
	stack *lane.Deque
}

func fileToString(file *os.File) (str string) {
	buf := bytes.NewBuffer(nil)
	buf.ReadFrom(file)
	file.Close()
	return buf.String()
}

// ParseFile func
func ParseFile(file *os.File) {
	toParse := fileToString(file)
	splitIntoTokens().split(&toParse).print()
}
