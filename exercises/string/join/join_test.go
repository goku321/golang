package main

import (
	"strconv"
	"testing"
	"time"
)

func BenchmarkJoinString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		joinString("string", strconv.Itoa(1234), "true", time.Now().String())
	}
}
