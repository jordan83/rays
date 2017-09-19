package main

type RGBColor struct {
	R, G, B float64
}

func NewRGBColor() *RGBColor {
	return &RGBColor{0, 0, 0}
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
