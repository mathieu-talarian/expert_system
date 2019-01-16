package main

import (
	"os"
	"regexp"

	lane "gopkg.in/oleiade/lane.v1"
)

const (
	OpenParenthesis = 0
	CloseParenthesis
	Not
	And
	Or
	Xor
	Implies
	If
	Comment
	EndLine
	Axiom
	Unknow
	InitialFacts
	Queries
)

func parseFile(file *os.File) {

}

type TokenInfo struct {
	Token int
	Regex *regexp.Regexp
}

type Parser struct {
	index int

	// rules ParseResult
	token map[string]string
	stack *lane.Deque
}

func splitIntoTokens(to string) []*TokenInfo {
	tokenTypes := []*TokenInfo{
		&TokenInfo{OpenParenthesis, regexp.MustCompile("(")},
		&TokenInfo{CloseParenthesis, regexp.MustCompile(")")},
		&TokenInfo{Not, regexp.MustCompile("!")},
		&TokenInfo{Ans}
	}
	return tokenTypes
}
