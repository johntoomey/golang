package main

import "fmt"

var doors [100]bool // how to do this with a struct/type

func main() {

	for step := 1; step < 100; step++ {
		for i := 0; i < 100; i += step {
			//fmt.Printf("%d ", i)
			doors[i] = !doors[i]
		}
		//fmt.Println("next")
	}
	fmt.Println(doors)
}
