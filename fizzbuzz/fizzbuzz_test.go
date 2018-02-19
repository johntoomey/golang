package main

import "testing"

func BenchmarkFizzbuzz(b *testing.B) {

	for i := 0; i < b.N; i++ {
		fizzbuzzB()
	}

}
