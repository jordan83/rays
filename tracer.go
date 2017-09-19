package main

type Tracer struct {
	world *World
}

func (t Tracer) TraceRay(ray *Ray) *RGBColor {
	sr := NewShadeRec(t.world)
	success, _ := t.world.Sphere.Hit(ray, &sr)
	if success {
		return Red()
	}
	return Black()
}
