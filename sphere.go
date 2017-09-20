package main

import (
	"math"
)

type Sphere struct {
	center Point3D
	radius float64
	color  *RGBColor // TEMP
}

func NewSphere(center Point3D, radius float64) *Sphere {
	return &Sphere{center, radius, Black()}
}

func (s *Sphere) Hit(ray *Ray, shadeRec *ShadeRec) (bool, float64) {
	temp := ray.Origin.Subtract(s.center)
	a := ray.Direction.Dot(ray.Direction)
	b := temp.ScalarMul(2.0).Dot(ray.Direction)
	c := temp.Dot(temp) - (s.radius * s.radius)
	disc := b*b - 4.0*a*c

	if disc < 0 {
		return false, -1
	}

	forHit := func(tMin float64) (bool, float64) {
		shadeRec.Normal = NewNormalFromVector(temp.Add(ray.Direction.ScalarMul(tMin)).ScalarDiv(s.radius))
		shadeRec.LocalHitPoint = ray.Origin.Add(ray.Direction.ScalarMul(tMin))
		return true, tMin
	}

	e := math.Sqrt(disc)
	denom := 2.0 * a

	t := (-b - e) / denom // smaller root
	if t > K_EPSILON {
		return forHit(t)
	}

	t = (-b + e) / denom // larger root
	if t > K_EPSILON {
		return forHit(t)
	}

	return false, -1
}

func (s *Sphere) SetColor(color *RGBColor) {
	s.color = color
}

func (s *Sphere) GetColor() *RGBColor {
	return s.color
}
