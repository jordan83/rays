package main

import (
	"math"
)

type RGBColor struct {
	R, G, B float64
}

func NewRGBColor() *RGBColor {
	return &RGBColor{0, 0, 0}
}

func (c *RGBColor) MaxToOne() *RGBColor {
	max := math.Max(c.R, math.Max(c.G, c.B))
	if max > 1.0 {
		return c.DivideBy(max)
	}
	return c
}

func (c *RGBColor) DivideBy(val float64) *RGBColor {
	return &RGBColor{c.R / val, c.G / val, c.B / val}
}

func Black() *RGBColor {
	return NewRGBColor()
}

func White() *RGBColor {
	return &RGBColor{1, 1, 1}
}

func Red() *RGBColor {
	return &RGBColor{1, 0, 0}
}
