package trace

import "github.com/K1ngArtes/ray-tracing-in-one-weekend/geom"

type Ray struct {
	Origin geom.Vec3
	Dir    geom.Vec3
}

func NewRay(origin geom.Vec3, dir geom.Vec3) Ray {
	return Ray{origin, dir}
}

func (r *Ray) At(t float64) geom.Vec3 {
	return r.Origin.Plus(r.Dir.Scaled(t))
}
