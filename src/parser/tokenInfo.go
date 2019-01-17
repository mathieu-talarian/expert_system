package parser

import (
	"regexp"
)

type TokenInfo struct {
	Type  int
	Regex *regexp.Regexp
}

type TokenType struct {
	Type    int
	Content *string
}

func (t *TokenInfo) isMatch(str *string) *TokenType {
	if t.Regex.Match([]byte(*str)) {
		if pos := t.Regex.FindIndex([]byte(*str)); pos != nil {
			if pos[0] == 0 {
				if content := substring(str, pos[0], pos[1]); content != nil {
					return &TokenType{t.Type, content}
				}
			}
		}
	}
	return nil
}
