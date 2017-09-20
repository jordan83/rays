package main

type ShadeRec struct {
	HitAnObject   bool
	LocalHitPoint Point3D
	Normal        Normal
	Color         *RGBColor
	World         *World
}

func NewShadeRec(world *World) ShadeRec {
	return ShadeRec{
		false,
		NewPoint3D(),
		NewNormal(),
		Black(),
		world,
	}
}
