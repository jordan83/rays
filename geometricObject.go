package main

type GeometricObject interface {
	// TODO do i really want to pass the shade rec? Might be a better way.
	Hit(ray *Ray, s *ShadeRec) (bool, float64)
}
