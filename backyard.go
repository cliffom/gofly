package main

import (
	"time"

	"github.com/gdamore/tcell/v2"
)

type Backyard struct {
	Screen tcell.Screen
	Flies  []*Fly
}

func (b *Backyard) Simulate() {
	s := b.Screen

	for {
		s.Clear()
		width, height := s.Size()

		for _, fly := range b.Flies {
			runes := fly.Draw()

			defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(fly.color)
			s.SetContent(fly.x, fly.y, runes[0], nil, defStyle)
			s.SetContent(fly.x+1, fly.y, runes[1], nil, defStyle)

			fly.EdgeCheck(width, height)
			fly.Move()
		}
		s.Show()
		time.Sleep(250 * time.Millisecond)
	}
}
