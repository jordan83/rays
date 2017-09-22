package main

type ViewPlane struct {
	Hres, Vres                 int
	PixelSize, Gamma, InvGamma float32
	ShowOutOfGamut             bool
	NumSamples                 int
}

func (v ViewPlane) SetGamma(g float32) {
	v.Gamma = g
	v.InvGamma = 1.0 / g
}
