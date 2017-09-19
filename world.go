package main

type World struct {
	viewPlane       ViewPlane
	backgroundColor RGBColor
	objects         []GeometricObject
	Sphere          Sphere
	tracer          Tracer
}

type RenderCallback func(row, col int, color *RGBColor)

func NewWorld() *World {
	w := World{}
	w.viewPlane.Hres = 200
	w.viewPlane.Vres = 200
	w.viewPlane.PixelSize = 1.0
	w.viewPlane.SetGamma(1.0)

	w.backgroundColor = *White()
	w.tracer = Tracer{&w}

	w.Sphere = Sphere{Point3D{0, 0, 0}, 85}

	return &w
}

func (w *World) GetResolution() (int, int) {
	return w.viewPlane.Hres, w.viewPlane.Vres
}

func (w *World) RenderScene(callback RenderCallback) {
	ray := Ray{Direction: NewVector3D(0, 0, -1)}
	zw := 100.0
	vres := float32(w.viewPlane.Vres)
	hres := float32(w.viewPlane.Hres)

	var r, c float32
	for r = 0.0; r < vres; r++ {
		for c = 0.0; c < hres; c++ {
			ray.Origin = Point3D{float64(w.viewPlane.PixelSize * (c - hres/2.0 + 0.5)), float64(w.viewPlane.PixelSize * (r - vres/2.0 + 0.5)), zw}
			pixelColor := w.tracer.TraceRay(&ray)
			w.displayPixel(int(r), int(c), pixelColor, callback)
		}
	}

}

func (w *World) displayPixel(row, col int, color *RGBColor, callback RenderCallback) {
	x := col
	y := w.viewPlane.Vres - row - 1
	callback(x, y, color)
}
