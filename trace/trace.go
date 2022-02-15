package trace

import "github.com/K1ngArtes/ray-tracing-in-one-weekend/geom"

// Hit records the details of a Ray->Surface intersection.
type Hit struct {
	Point geom.Vec3
	Normal geom.Vec3
	t float64
}

type Hittable interface {
	Hit(r Ray, tMin float64, tMax float64, hitRecord *Hit) bool
}