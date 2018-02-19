package main

import "fmt"

func permute(str string, l, r int) {

	runes := []rune(str)

	if l == r {
		// add to map to remove duplicates?
		fmt.Println(str)
	} else {
		for i := l; i <= r; i++ {
			runes[l], runes[i] = runes[i], runes[l]
			permute(string(runes), l+1, r)
		}
	}
}

func anagrams(str string) {
	permute(str, 0, len(str)-1)
}

func main() {
	anagrams("asdfde")
}
