package main

import (
	"fmt"
	"unicode"
)

func LetterChanges(str string) string {

	runes := []rune(str)

	for i, r := range runes {
		if unicode.IsLetter(r) {
			if r == 'z' {
				runes[i] = 'a'
			} else {
				runes[i] = r + 1
			}
		}
	}

	for i, r := range runes {
		if unicode.IsLetter(r) {
			for _, v := range "aeiou" {
				if r == v {
					runes[i] = unicode.ToUpper(r)
				}
			}
		}
	}

	return string(runes)
}

func main() {

	// do not modify below here, readline is our function
	// that properly reads in the input for you
	// fmt.Println(LetterChanges("hello*3"))
	fmt.Println(LetterChanges("fun times!"))

}
