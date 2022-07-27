package main

import (
	"math/rand"

	"github.com/AvraamMavridis/randomcolor"
	"github.com/gdamore/tcell/v2"
)

// Fly is a digital representation of a fly
// with position, velocity, color, and animation information
type Fly struct {
	x     int
	y     int
	vx    int
	vy    int
	color tcell.Color
	frame int
}

// Draw gives life to our little fly friend
func (f *Fly) Draw() []rune {
	frames := []string{
		"\\/",
		"--",
		"/\\",
		"--",
	}

	return []rune(frames[f.frame])
}

// Move *ahem* moves a fly into their next position
func (f *Fly) Move() {
	f.x += f.vx
	f.y += f.vy

	if f.frame < 3 {
		f.frame++
	} else {
		f.frame = 0
	}
}

// EdgeCheck ensures our fly won't fly beyond the constraints
// of the space they occupy
func (f *Fly) EdgeCheck(maxWidth, maxHeight int) {
	f.vx = getVelocity(f.x, maxWidth-len(f.Draw()))
	f.vy = getVelocity(f.y, maxHeight-1)
}

// NewFly returns a, you guessed it, pointer to a new fly
func NewFly(w, h int) *Fly {
	return &Fly{
		x:     rand.Intn(w - 1),
		y:     rand.Intn(h),
		vx:    1,
		vy:    1,
		color: tcell.GetColor(randomcolor.GetRandomColorInHex()),
		frame: rand.Intn(4),
	}
}

// getVelocity determines, from a given 1-dimensional position, if
// the direction should change or if we can fly chaotically
func getVelocity(pos, limit int) int {
	if pos <= 0 {
		return 1
	} else if pos >= limit {
		return -1
	} else {
		return chaos()
	}
}

// chaos introduces chaotic/unpredictable movement
func chaos() int {
	switch rand.Intn(3) {
	case 0:
		return 1
	case 1:
		return 0
	default:
		return -1
	}
}
