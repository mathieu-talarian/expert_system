package parser

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"gopkg.in/karalabe/cookiejar.v1/collections/deque"
)

// Parser struct
type Parser struct {
	Index  int
	Tokens *TokenTypes
	Stack  *deque.Deque

	*ParseResult
	*TokenInfos
	// Save old state on parser if i have to process further
	oldState *int
}

func NewParser(index int, p *ParseResult, t *TokenTypes, s *deque.Deque) *Parser {
	return &Parser{
		Index:       index,
		ParseResult: p,
		Tokens:      t,
		Stack:       s,
		TokenInfos:  splitIntoTokens(),
		oldState:    nil,
	}
}

func fileToString(file *os.File) (str string) {
	buf := bytes.NewBuffer(nil)
	buf.ReadFrom(file)
	file.Close()
	return buf.String()
}

func (t *TokenTypes) SearchForUnknown() error {
	for _, v := range *t {
		if v.Type == Unknown {
			return fmt.Errorf("issue, unknown character %s", *v.Content)
		}
	}
	return nil
}

// ParseFile func
func ParseFile(file *os.File) error {
	toParse := fileToString(file)
	entities := splitIntoTokens().split(&toParse)
	entities.print()
	if err := entities.SearchForUnknown(); err != nil {
		log.Fatal(err)
	}
	entities.appendEndline()
	parser := NewParser(0, NewParseResult(), entities, deque.New())
	parser.Process()
	return nil
}
