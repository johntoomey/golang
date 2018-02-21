package main

import "fmt"

func fib(n int) {
	last, current := 0, 1
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", current)
		last, current = current, last+current
	}
}

func main() {
	fib(10)
}
