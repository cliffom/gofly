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
	if f.x <= 1 || f.x >= maxWidth-1 {
		f.vx *= -1
	} else {
		f.vx = Chaos(f.vx)
	}

	if f.y <= 0-1 || f.y >= maxHeight-1 {
		f.vy *= -1
	} else {
		f.vy = Chaos(f.vy)
	}
}

// NewFly returns a, you guessed it, pointer to a new fly
func NewFly(w, h int) *Fly {
	return &Fly{
		x:     rand.Intn(w),
		y:     rand.Intn(h),
		vx:    1,
		vy:    1,
		color: tcell.GetColor(randomcolor.GetRandomColorInHex()),
		frame: rand.Intn(4),
	}
}

// Chaos introduces chaotic/unpredictable movement
func Chaos(d int) int {
	switch rand.Intn(2) {
	case 0:
		return d * 1
	default:
		return d * -1
	}
}
