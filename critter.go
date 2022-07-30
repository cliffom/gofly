package main

type Critter interface {
	Draw() []rune
	Move()
	UpdateVelocity(w, h int)
	GetColor() Color
	GetPos() (x, y int)
}
