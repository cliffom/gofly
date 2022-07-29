package main

import "github.com/gdamore/tcell/v2"

type Critter interface {
	Draw() []rune
	Move()
	UpdateVelocity(w, h int)
	GetColor() tcell.Color
	GetPos() (x, y int)
}
