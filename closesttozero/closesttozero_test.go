package main

import "testing"

var closesttozeroTests = []struct {
	n        []int
	expected int
}{
	{[]int{2, 3, 4}, 2},
	{[]int{-2, -3, -4}, -2},
	{[]int{2, 3, 0}, 0},
	{[]int{-2, 3, 4}, -2},
}

func TestClosesttozero(t *testing.T) {
	for _, test := range closesttozeroTests {
		output := closesttozero(test.n)
		if output != test.expected {
			t.Errorf("asdf")
		}
	}
}
