package main

import "fmt"

func ScaleBalancing(scale []int, weights []int) string {
	fmt.Println(scale)

	if scale[0] == scale[1] {
		fmt.Println("they are equal")
	}

	for _, num := range weights {
		if scale[0]+num == scale[1] || scale[0] == scale[1]+num {
			fmt.Println(num)
		}
	}

	for i, num_1 := range weights {
		for j, num_2 := range weights {

			if i == j {
				continue
			}

			if scale[0]+num_1 == scale[1]+num_2 || scale[0]+num_2 == scale[1]+num_1 {
				fmt.Println(num_1, num_2)
			}
		}
	}

	return ""
}

func main() {

	s := []int{3, 4}
	b := []int{1, 2, 7, 7}

	fmt.Println(ScaleBalancing(s, b))
}
