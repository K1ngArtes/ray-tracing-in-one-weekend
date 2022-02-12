package trace

import "github.com/K1ngArtes/ray-tracing-in-one-weekend/math"

type Ray struct {
	Origin math.Vec3
	Dir    math.Vec3
}

func (r Ray) At(t float64) math.Vec3 {
	return r.Origin.Plus(r.Dir.Scaled(t))
}
