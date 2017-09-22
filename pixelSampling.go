package main

import (
	"math"
	"math/rand"
)

const (
	ZW float64 = 100.0
)

type JitterProvider func() float64

func JitteredSampling(w *World, ray *Ray, r, c int) *RGBColor {
	return RegularSamplingWithJitter(w, ray, r, c, func() float64 { return rand.Float64() })
}

func RegularSampling(w *World, ray *Ray, r, c int) *RGBColor {
	return RegularSamplingWithJitter(w, ray, r, c, func() float64 { return 0.5 })
}

func RegularSamplingWithJitter(w *World, ray *Ray, r, c int, jitterProvider JitterProvider) *RGBColor {
	vp := w.GetViewPlane()
	tracer := w.GetTracer()

	color := w.GetBackgroundColor()
	pp := NewPoint3D()

	n := float64(int(math.Sqrt(float64(vp.NumSamples))))

	for p := 0; p < int(n); p++ {
		for q := 0; q < int(n); q++ {
			pp.X = float64(vp.PixelSize) * (float64(c) - 0.5*float64(vp.Hres) + (float64(q)+jitterProvider())/n)
			pp.Y = float64(vp.PixelSize) * (float64(r) - 0.5*float64(vp.Vres) + (float64(p)+jitterProvider())/n)

			ray.Origin = Point3D{pp.X, pp.Y, ZW}
			color = color.Add(tracer.TraceRay(ray))
		}
	}

	return color.DivideBy(float64(vp.NumSamples))
}

func RandomSampling(w *World, ray *Ray, r, c int) *RGBColor {
	vp := w.GetViewPlane()
	tracer := w.GetTracer()

	color := w.GetBackgroundColor()
	pp := NewPoint3D()

	for p := 0; p < vp.NumSamples; p++ {
		pp.X = float64(vp.PixelSize) * (float64(c) - 0.5*float64(vp.Hres) + rand.Float64())
		pp.Y = float64(vp.PixelSize) * (float64(r) - 0.5*float64(vp.Vres) + rand.Float64())
		ray.Origin = Point3D{pp.X, pp.Y, ZW}
		color = color.Add(tracer.TraceRay(ray))
	}

	return color.DivideBy(float64(vp.NumSamples))
}

func NoSampling(w *World, ray *Ray, r, c int) *RGBColor {
	hres := float64(w.viewPlane.Hres)
	vres := float64(w.viewPlane.Vres)

	ray.Origin = Point3D{float64(w.viewPlane.PixelSize) * (float64(c) - hres/2.0 + 0.5), float64(w.viewPlane.PixelSize) * (float64(r) - vres/2.0 + 0.5), ZW}
	return w.tracer.TraceRay(ray)
}
