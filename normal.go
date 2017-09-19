package main

type Normal struct {
	X, Y, Z float64
}

func NewNormal() Normal {
	return Normal{0, 0, 0}
}

func NewNormalFromVector(v Vector3D) Normal {
	return Normal{
		v.X,
		v.Y,
		v.Z,
	}
}
