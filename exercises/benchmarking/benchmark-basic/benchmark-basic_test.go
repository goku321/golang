package main

import (
	"fmt"
	"testing"
)

// BenchmarkHello xxx
func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("hello")
	}
}
