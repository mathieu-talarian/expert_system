package parser

import (
	"fmt"
	"regexp"
	"strings"
)

const (
	OpenParenthesis = iota
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
	Unknown
	InitialFacts
	Queries
)

type TokenInfos []*TokenInfo
type TokenTypes []*TokenType

func splitIntoTokens() *TokenInfos {
	return &TokenInfos{
		&TokenInfo{OpenParenthesis, regexp.MustCompile("\\(")},
		&TokenInfo{CloseParenthesis, regexp.MustCompile("\\)")},
		&TokenInfo{Not, regexp.MustCompile("!")},
		&TokenInfo{And, regexp.MustCompile("\\+")},
		&TokenInfo{Or, regexp.MustCompile("\\|")},
		&TokenInfo{Xor, regexp.MustCompile("\\^")},
		&TokenInfo{Implies, regexp.MustCompile("=>")},
		&TokenInfo{If, regexp.MustCompile("<=>")},
		&TokenInfo{EndLine, regexp.MustCompile("\n")},
		&TokenInfo{Axiom, regexp.MustCompile("[a-zA-Z]")},
		&TokenInfo{Comment, regexp.MustCompile("#.*")},
		&TokenInfo{InitialFacts, regexp.MustCompile("=")},
		&TokenInfo{Queries, regexp.MustCompile("\\?")},
		&TokenInfo{Unknown, regexp.MustCompile(".*")},
	}
}

func (t *TokenTypes) print() *TokenTypes {

	for _, v := range *t {
		fmt.Println(*v.Content, v.Type)
	}
	return t
}

func (t *TokenInfos) split(str *string) (tt *TokenTypes) {
	var e = []*TokenType{}
	for len(*str) > 0 {
		if (*str)[0] != '\n' {
			*str = strings.TrimSpace(*str)
		}
		e = append(e, t.matchToken(str))
	}
	return (*TokenTypes)(&e)
}

func (t *TokenInfos) matchToken(str *string) *TokenType {
	for _, v := range *t {
		if tk := v.isMatch(str); tk != nil {
			return tk
		}
	}
	return nil
}
