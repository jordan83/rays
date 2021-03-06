package main

type Plane struct {
	point  Point3D
	normal Normal
	color  *RGBColor
}

func NewPlane(point Point3D, normal Normal) *Plane {
	return &Plane{point, normal, Black()}
}

func (p *Plane) Hit(ray *Ray, shadeRec *ShadeRec) (bool, float64) {
	t := p.point.Subtract(ray.Origin).DotNormal(p.normal) / (ray.Direction.DotNormal(p.normal))
	if t > K_EPSILON {
		shadeRec.Normal = p.normal
		shadeRec.LocalHitPoint = ray.Origin.Add(ray.Direction.ScalarMul(t))
		return true, t
	}
	return false, 0
}

func (p *Plane) SetColor(color *RGBColor) {
	p.color = color
}

func (p *Plane) GetColor() *RGBColor {
	return p.color
}
