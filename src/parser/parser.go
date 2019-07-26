package parser

import (
	"fmt"
	"strings"
)

func removeEmptyFields(slice []string) (result []string) {
	for _, v := range slice {
		if v != "" {
			result = append(result, v)
		}
	}
	return
}

func removeCommentedLine(slice []string) (result []string) {
	for _, v := range slice {
		if v != "#" {
			result = append(result, v)
		}
	}
	return slice
}

// Start parsing opened file
func Start(openedFileToString string) (ret *interface{}, err error) {
	trimmed := strings.TrimSpace(openedFileToString)

	split := strings.Split(trimmed, "\n")

	split = append(removeEmptyFields(split))
	split = append(removeCommentedLine(split))
	for _, v := range split {
		fmt.Println(string(v[0]))
	}
	return nil, nil
}
