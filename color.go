package main

import "github.com/AvraamMavridis/randomcolor"

type Color struct {
	HexValue string
}

// String returns the string representation of a color
func (c Color) String() string {
	return c.HexValue
}

// NewRandomColor returns a random color in hex format
func NewRandomColor() Color {
	return Color{
		HexValue: randomcolor.GetRandomColorInHex(),
	}
}
