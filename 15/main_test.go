package main

import "testing"

const n = 1 << 7

func BenchmarkRunes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		createHugeString(n)
	}
}
