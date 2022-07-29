package main

import (
	"math/rand"

	"github.com/AvraamMavridis/randomcolor"
	"github.com/gdamore/tcell/v2"
)

var frames = []string{
	">\u25CF<",
	"-\u25CF-",
}

// Fly is a digital representation of a fly
// with position, velocity, color, and animation information
type Fly struct {
	x      int
	y      int
	vx     int
	vy     int
	color  tcell.Color
	frame  int
	frames []string
}

// Draw gives life to our little fly friend
func (f *Fly) Draw() []rune {
	return []rune(f.frames[f.frame])
}

// UpdateVelocity updates a fly's movement based on the
// width and height of the area it occupies. If the next movement
// would result in going beyond the boundaries, a fly will change
// direction. Otherwise let the flight be random (chaos)
func (f *Fly) UpdateVelocity(w, h int) {
	getVel := func(pos, max int) int {
		if pos <= 0 {
			return 1
		} else if pos >= max {
			return -1
		}
		return chaos()
	}

	f.vx = getVel(f.x, w-len(f.Draw()))
	f.vy = getVel(f.y, h-1)
}

// Move *ahem* moves a fly into their next position
func (f *Fly) Move() {
	f.x += f.vx
	f.y += f.vy

	if f.frame < len(f.frames)-1 {
		f.frame++
	} else {
		f.frame = 0
	}
}

// GetColor returns a fly's color
func (f *Fly) GetColor() tcell.Color {
	return f.color
}

// GetPos returns a fly's current position
func (f *Fly) GetPos() (x, y int) {
	return f.x, f.y
}

// NewFly returns a, you guessed it, pointer to a new fly
func NewFly(w, h int) *Fly {
	return &Fly{
		x:      rand.Intn(w - 1),
		y:      rand.Intn(h),
		vx:     rand.Intn(2),
		vy:     rand.Intn(2),
		color:  tcell.GetColor(randomcolor.GetRandomColorInHex()),
		frame:  rand.Intn(len(frames)),
		frames: frames,
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
