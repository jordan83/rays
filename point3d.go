package main

type Point3D struct {
	X, Y, Z float64
}

func NewPoint3D() Point3D {
	return Point3D{0, 0, 0}
}

func (p Point3D) Subtract(other Point3D) Vector3D {
	return NewVector3D(p.X-other.X, p.Y-other.Y, p.Z-other.Z)
}

func (p Point3D) Add(v Vector3D) Point3D {
	return Point3D{p.X + v.X, p.Y + v.Y, p.Z + v.Z}
}
