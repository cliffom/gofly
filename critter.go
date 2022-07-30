package main

type Critter interface {
	Draw() []rune
	Move()
	UpdateVelocity(w, h int)
	GetColor() string
	GetPos() (x, y int)
}
