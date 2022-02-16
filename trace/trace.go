package trace

import "github.com/K1ngArtes/ray-tracing-in-one-weekend/geom"

// Hit records the details of a Ray->Surface intersection.
type Hit struct {
	Point geom.Vec3
	normal geom.Vec3
	T float64
	IsFrontFace bool
}

type Hittable interface {
	Hit(r Ray, tMin float64, tMax float64, hitRecord *Hit) bool
}

func (hit Hit) SetFaceNormal(r *Ray, outwardNormal *geom.Vec3) {
	hit.IsFrontFace = r.Dir.Dot(*outwardNormal) < 0
	if(hit.IsFrontFace) {
		hit.normal = *outwardNormal
	} else {
		hit.normal = outwardNormal.Negate()
	}
}