package main

type Tracer struct {
	world *World
}

func (t Tracer) TraceRay(ray *Ray) *RGBColor {
	sr := t.hitBareBonesObjects(ray)

	if sr.HitAnObject {
		return sr.Color
	}
	return t.world.GetBackgroundColor()
}

func (t Tracer) hitBareBonesObjects(ray *Ray) ShadeRec {
	sr := NewShadeRec(t.world)
	tmin := K_HUGE_VALUE

	for _, obj := range t.world.GetObjects() {
		success, t := obj.Hit(ray, &sr)
		if success && t < tmin {
			sr.HitAnObject = true
			tmin = t
			sr.Color = obj.GetColor()
		}
	}
	return sr
}
