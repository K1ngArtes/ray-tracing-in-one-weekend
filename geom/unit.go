package geom

type Unit Vec3

func (u Unit) Dot(u2 Unit) float64 {
	return Vec3(u).Dot(Vec3(u2))
}

func (u Unit) Inv() Unit {
	return Unit(Vec3(u).Inv())
}
