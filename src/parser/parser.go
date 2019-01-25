package parser

import (
	"bytes"
	"os"

	lane "gopkg.in/oleiade/lane.v1"
)

// Parser struct
type Parser struct {
	index int

	*ParseResult
	// rules ParseResult
	tokens *TokenTypes
	stack  *lane.Deque
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
	entities := splitIntoTokens().split(&toParse).print()
	*entities = append(*entities, &TokenType{EndLine, func() (s *string) {
		*s = "\n"
		return
	}()})
	parser := &Parser{
		index:       0,
		ParseResult: NewParseResult(),
		tokens:      entities,
		stack:       lane.NewDeque(),
	}
	carryOn := true
}
