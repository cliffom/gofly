package main

import (
	"math/rand"
	"testing"
	"time"
)

// BenchmarkNewFly tests creating N flies
func BenchmarkNewFly(b *testing.B) {
	flies := make([]*Fly, b.N)
	for i := 0; i < b.N; i++ {
		flies[i] = NewFly(100, 100)
	}
}

// BenchmarkFly tests animating a single fly N times
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

// BenchmarkFlies tests animating N flies [500,1000) times
func BenchmarkFlies(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	numFlies := rand.Intn(500) + 500
	maxWidth := 1000
	maxHeight := 1000

	flies := make([]*Fly, numFlies)
	for i := 0; i < numFlies; i++ {
		flies[i] = NewFly(maxWidth, maxHeight)
	}

	for i := 0; i < b.N; i++ {
		for _, fly := range flies {
			fly.Draw()
			fly.EdgeCheck(maxWidth, maxHeight)
			fly.Move()
		}
	}
}
