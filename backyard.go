package main

import (
	"time"

	"github.com/gdamore/tcell/v2"
)

type Backyard struct {
	Screen    tcell.Screen
	Flies     []*Fly
	Frametime time.Duration
}

func (b *Backyard) Simulate() {
	s := b.Screen
	borderStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorRed)

	for {
		s.Clear()
		w, h := s.Size()

		// Draw border
		drawBorder(w, h, s, borderStyle)

		// The flies shall fly!
		for _, fly := range b.Flies {
			animateCritter(w, h, s, fly)
		}
		s.Show()
		time.Sleep(b.Frametime * time.Millisecond)
	}
}

// SpeedUp decreases the delay between frames, thus speeding up the
// simulation
func (b *Backyard) SpeedUp() bool {
	minFrametime := 100
	step := 25
	if b.Frametime >= time.Duration(minFrametime+step) {
		b.Frametime -= time.Duration(step)
		return true
	}

	return false
}

// SpeedDown increases the delay between frames, thus slowing down the
// simulation
func (b *Backyard) SpeedDown() bool {
	maxFrametime := 2000
	step := 25
	if b.Frametime <= time.Duration(maxFrametime-step) {
		b.Frametime += time.Duration(step)
	}

	return false
}

// NewBackyard returns a backyard for our flies to play in
func NewBackyard(s tcell.Screen, f []*Fly, t time.Duration) *Backyard {
	return &Backyard{
		Screen:    s,
		Flies:     f,
		Frametime: t,
	}
}

func drawBorder(w, h int, s tcell.Screen, style tcell.Style) {
	s.SetContent(0, 0, '\u25E4', nil, style)
	s.SetContent(0, h-1, '\u25E3', nil, style)
	s.SetContent(w-1, 0, '\u25E5', nil, style)
	s.SetContent(w-1, h-1, '\u25E2', nil, style)
}

func animateCritter(w, h int, s tcell.Screen, critter Critter) {
	runes := critter.Draw()

	color := tcell.GetColor(critter.GetColor().String())
	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(color)
	x, y := critter.GetPos()
	for i := 0; i < len(runes); i++ {
		s.SetContent(x+i, y, runes[i], nil, defStyle)
	}

	critter.UpdateVelocity(w, h)
	critter.Move()
}
