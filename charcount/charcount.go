package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	b, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("err")
	}
	str := string(b)

	count := make(map[rune]int)
	runes := ".,_!£$%^&*(~{+'"

	// Add all of the characters to the map with value 0
	for _, r := range runes {
		count[r] = 0
	}

	// Loop over the string and count the values
	for _, r := range str {
		if _, ok := count[r]; ok {
			count[r]++
		}
	}

	// Print all of the counts
	for k, v := range count {
		fmt.Printf("%q %d\n", k, v)
	}
}
