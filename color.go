package main

import "github.com/AvraamMavridis/randomcolor"

type Color struct {
	HexValue string
}

func (c Color) String() string {
	return c.HexValue
}

func NewRandomColor() Color {
	return Color{
		HexValue: randomcolor.GetRandomColorInHex(),
	}
}
