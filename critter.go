package main

import "github.com/gdamore/tcell/v2"

type Critter interface {
	Draw() []rune
	Move()
	CheckBounds(w, h int)
	GetColor() tcell.Color
	GetPos() (int, int)
}
