package main

import (
	"testing"
)

func BenchmarkNewFly(b *testing.B) {
	flies := make([]*Fly, b.N)
	for i := 0; i < b.N; i++ {
		flies[i] = NewFly(100, 100)
	}
}

func BenchmarkFly(b *testing.B) {
	maxWidth := 1000
	maxHeight := 1000
	fly := NewFly(maxWidth, maxHeight)

	for i := 0; i < b.N; i++ {
		fly.Draw()
		fly.EdgeCheck(maxWidth, maxHeight)
		fly.Move()
	}
}
