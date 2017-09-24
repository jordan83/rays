package main

import (
	"fmt"
)

const (
	ZW float64 = 100.0
)

type World struct {
	viewPlane       ViewPlane
	backgroundColor *RGBColor
	objects         []GeometricObject
	tracer          Tracer
	sampler         *Sampler
}

type RenderCallback func(row, col int, color *RGBColor)

func NewWorld() *World {
	w := World{}
	w.viewPlane.Hres = 200
	w.viewPlane.Vres = 200
	w.viewPlane.PixelSize = 1.0
	w.viewPlane.SetGamma(1.0)
	w.viewPlane.NumSamples = 16

	w.backgroundColor = Black()
	w.tracer = Tracer{&w}
	w.sampler = NewSampler(w.viewPlane.NumSamples)

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
	vp := w.viewPlane
	ray := Ray{Direction: NewVector3D(0, 0, -1)}
	pp := NewPoint3D()

	for r := 0; r < w.viewPlane.Vres; r++ {
		for c := 0; c < w.viewPlane.Hres; c++ {

			pixelColor := Black()
			for j := 0; j < vp.NumSamples; j++ {
				sample := w.sampler.SampleUnitSquare()

				pp.X = float64(vp.PixelSize) * (float64(c) - 0.5*float64(vp.Hres) + sample.X)
				pp.Y = float64(vp.PixelSize) * (float64(r) - 0.5*float64(vp.Vres) + sample.Y)
				ray.Origin = Point3D{pp.X, pp.Y, ZW}

				pixelColor = pixelColor.Add(w.tracer.TraceRay(&ray))
			}

			w.displayPixel(r, c, pixelColor.DivideBy(float64(vp.NumSamples)), callback)
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
