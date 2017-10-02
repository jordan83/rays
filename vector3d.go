package main

import (
	"math"
)

type Vector3D struct{ X, Y, Z float64 }

func NewVector3D(x, y, z float64) Vector3D {
	return Vector3D{x, y, z}
}

func (v Vector3D) Negate() Vector3D {
	return Vector3D{-v.X, -v.Y, -v.Z}
}

func (v Vector3D) ScalarMul(x float64) Vector3D {
	return Vector3D{x * v.X, x * v.Y, x * v.Z}
}

func (v Vector3D) ScalarDiv(x float64) Vector3D {
	return v.ScalarMul(1 / x)
}

func (v Vector3D) LengthSquared() float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

func (v Vector3D) Add(other Vector3D) Vector3D {
	return Vector3D{v.X + other.X, v.Y + other.Y, v.Z + other.Z}
}

func (v Vector3D) Subtract(other Vector3D) Vector3D {
	return Vector3D{v.X - other.X, v.Y - other.Y, v.Z - other.Z}
}

func (v Vector3D) Dot(other Vector3D) float64 {
	return v.X*other.X + v.Y*other.Y + v.Z*other.Z
}

func (v Vector3D) DotNormal(n Normal) float64 {
	return v.X*n.X + v.Y*n.Y + v.Z*n.Z
}

func (v Vector3D) Cross(other Vector3D) Vector3D {
	return Vector3D{v.Y*other.Z - v.Z - other.Y, v.Z*other.X - v.X*other.Z, v.X*other.Y - v.Y*other.X}
}

func (v Vector3D) Normalize() Vector3D {
	length := math.Sqrt(v.LengthSquared())
	return Vector3D{v.X / length, v.Y / length, v.Z / length}
}
