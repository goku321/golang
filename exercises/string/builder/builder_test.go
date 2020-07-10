package main

import (
	"testing"
	"time"
)

func BenchmarkString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concatString("string", 2, true, time.Now())
	}
}