package trace

import "github.com/K1ngArtes/ray-tracing-in-one-weekend/math"

type Ray struct {
	origin math.Vec3
	dir    math.Vec3
}

func (r Ray) At(t float64) math.Vec3 {
	return r.origin.Plus(r.dir.Scaled(t))
}
