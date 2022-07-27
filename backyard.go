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

		// Release the flies!
		for _, fly := range b.Flies {
			animateFly(w, h, s, fly)
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
	s.SetContent(0, 0, '#', nil, style)
	s.SetContent(0, h-1, '#', nil, style)
	s.SetContent(w-1, 0, '#', nil, style)
	s.SetContent(w-1, h-1, '#', nil, style)
}

func animateFly(w, h int, s tcell.Screen, fly *Fly) {
	runes := fly.Draw()

	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(fly.color)
	for i := 0; i < len(runes); i++ {
		s.SetContent(fly.x+i, fly.y, runes[i], nil, defStyle)
	}

	fly.EdgeCheck(w, h)
	fly.Move()
}
