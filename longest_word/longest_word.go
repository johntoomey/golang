package main

import (
	"regexp"
	"strings"
)

// LongestWord something something
func LongestWord(sen string) string {

	reg, _ := regexp.Compile("[^a-zA-Z0-9 ]+")
	stripped := reg.ReplaceAllString(sen, "")

	longest := 0
	longestIndex := 0

	for index, str := range strings.Fields(stripped) {
		if len(str) > longest {
			longest = len(str)
			longestIndex = index
		}
	}

	return strings.Fields(stripped)[longestIndex]
}
