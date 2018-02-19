package main

import "fmt"

func closesttozero(nums []int) int {

	closestPos := 1<<63 - 1
	closestNeg := -1 << 63
	hasNeg := false

	for _, n := range nums {
		if n == 0 {
			return 0
		}

		if n > 0 && n < closestPos {
			closestPos = n
		}

		if n < 0 && n > closestNeg {
			closestNeg = n
			hasNeg = true
		}
	}

	if closestPos > -closestNeg && hasNeg == true {
		return closestNeg
	}
	return closestPos
}

func main() {

	n := []int{-8, -3, -56, -3}

	fmt.Println(closesttozero(n))
}
