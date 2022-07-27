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

		for _, fly := range b.Flies {
			runes := fly.Draw()

			defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(fly.color)
			s.SetContent(fly.x, fly.y, runes[0], nil, defStyle)
			s.SetContent(fly.x+1, fly.y, runes[1], nil, defStyle)

			// Draw border at corners
			s.SetContent(0, 0, '#', nil, borderStyle)
			s.SetContent(0, h-1, '#', nil, borderStyle)
			s.SetContent(w-1, 0, '#', nil, borderStyle)
			s.SetContent(w-1, h-1, '#', nil, borderStyle)

			fly.EdgeCheck(w, h)
			fly.Move()
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
