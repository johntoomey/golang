package main

import "fmt"

/*
func fizzbuzzA() {
	for i := 0; i <= 100; i++ {

		if i%15 == 0 {
			fmt.Println("fizzbuzz")
		} else if i%3 == 0 {
			fmt.Println("fizz")
		} else if i%5 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
	}
}
*/

func fizzbuzzB() {
	for i := 1; i <= 100; i++ {
		result := ""
		if i%3 == 0 {
			result += "Fizz"
		}
		if i%5 == 0 {
			result += "Buzz"
		}
		if result != "" {
			fmt.Println(result)
			continue
		}
		fmt.Println(i)
	}
}
