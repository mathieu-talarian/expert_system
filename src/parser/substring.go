package parser

import (
	"strings"
)

func substring(str *string, pos0 int, pos1 int) *string {
	var ret string
	end := pos0 + pos1
	if pos0 > len(*str) {
		return nil
	}
	ret = strings.TrimSpace((*str)[pos0:end])
	*str = (*str)[end:]
	return &ret
}
