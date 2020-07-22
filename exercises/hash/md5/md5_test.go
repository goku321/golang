package main

import "testing"

func BenchmarkGetBytes(b *testing.B) {
	p := person{"yoda", 2009}

	for i := 0; i < b.N; i++ {
		getBytes(p)
	}
}

func BenchmarkGetBytesSprintF(b *testing.B) {
	p := person{"yoda", 2009}

	for i := 0; i < b.N; i++ {
		getBytesSprintF(p)
	}
}