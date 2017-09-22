package main

import (
	"fmt"
)

type World struct {
	viewPlane       ViewPlane
	backgroundColor *RGBColor
	objects         []GeometricObject
	tracer          Tracer
	sampler         PixelSampler
}

type RenderCallback func(row, col int, color *RGBColor)
type PixelSampler func(w *World, ray *Ray, row, col int) *RGBColor

func NewWorld() *World {
	w := World{}
	w.viewPlane.Hres = 200
	w.viewPlane.Vres = 200
	w.viewPlane.PixelSize = 1.0
	w.viewPlane.SetGamma(1.0)
	w.viewPlane.NumSamples = 16

	w.backgroundColor = Black()
	w.tracer = Tracer{&w}
	w.sampler = JitteredSampling

	sphere := NewSphere(Point3D{0, -25, 0}, 80)
	sphere.SetColor(&RGBColor{1, 0, 0})
	w.AddObject(sphere)

	sphere = NewSphere(Point3D{0, 30, 0}, 60)
	sphere.SetColor(&RGBColor{1, 1, 0})
	w.AddObject(sphere)

	plane := NewPlane(Point3D{0, 0, 0}, Normal{0, 1, 1})
	plane.SetColor(&RGBColor{0, 0.3, 0})
	w.AddObject(plane)

	return &w
}

func (w *World) GetResolution() (int, int) {
	return w.viewPlane.Hres, w.viewPlane.Vres
}

func (w *World) RenderScene(callback RenderCallback) {
	fmt.Printf("Starting...\n")
	ray := Ray{Direction: NewVector3D(0, 0, -1)}

	for r := 0; r < w.viewPlane.Vres; r++ {
		for c := 0; c < w.viewPlane.Hres; c++ {
			pixelColor := w.sampler(w, &ray, r, c)
			w.displayPixel(r, c, pixelColor, callback)
		}
	}

	fmt.Printf("Done!")
}

func (w *World) AddObject(obj GeometricObject) {
	w.objects = append(w.objects, obj)
}

func (w *World) GetObjects() []GeometricObject {
	return w.objects
}

func (w *World) GetBackgroundColor() *RGBColor {
	return w.backgroundColor
}

func (w *World) GetViewPlane() ViewPlane {
	return w.viewPlane
}

func (w *World) GetTracer() Tracer {
	return w.tracer
}

func (w *World) displayPixel(row, col int, color *RGBColor, callback RenderCallback) {
	x := col
	y := w.viewPlane.Vres - row - 1
	callback(x, y, color.MaxToOne())
}
