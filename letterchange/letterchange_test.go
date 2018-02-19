package main

import "testing"

var changeTests = []struct {
	n        string
	expected string
}{
	{"abcde", "bcdEf"},
	{"", ""},
}

func TestLetterChange(t *testing.T) {
	for _, test := range changeTests {
		output := LetterChanges(test.n)
		if output != test.expected {
			t.Errorf("asdf")
		}
	}
}
