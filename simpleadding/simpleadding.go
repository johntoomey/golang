package main

import "fmt"

func SimpleAdding(num int) int {

	if num == 1 {
		return 1
	} else {
		return num + SimpleAdding(num-1)
	}

}

func main() {
	fmt.Println(SimpleAdding(4))
}
