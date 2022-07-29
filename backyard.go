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

	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(critter.GetColor())
	x, y := critter.GetPos()
	for i := 0; i < len(runes); i++ {
		s.SetContent(x+i, y, runes[i], nil, defStyle)
	}

	critter.UpdateVelocity(w, h)
	critter.Move()
}
