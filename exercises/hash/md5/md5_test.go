package main

import "testing"

func BenchmarkGetBytes(b *testing.B) {
	p := person{"yoda", 209}

	for i := 0; i < b.N; i++ {
		getBytes(p)
	}
}