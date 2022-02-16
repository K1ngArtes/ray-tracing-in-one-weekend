package trace

import (
	"math"

	"github.com/K1ngArtes/ray-tracing-in-one-weekend/geom"
)

type Sphere struct {
	Center geom.Vec3
	Radius float64
}

func NewSphere(center geom.Vec3, radius float64) Sphere {
	return Sphere{
		Center: center,
		Radius: radius,
	}
}

func (sphere Sphere) Hit(r Ray, tMin float64, tMax float64, hit *Hit) bool {
	oc := r.Origin.Minus(sphere.Center)

	a := r.Dir.LenSq()
	halfB := oc.Dot(r.Dir)
	c := oc.LenSq() - sphere.Radius * sphere.Radius
	discriminant := halfB*halfB - a*c

	if (discriminant < 0) {
        return false;
    }

	// Calculate actual hit point
	sqrtD := math.Sqrt(discriminant)
	root := (-halfB-sqrtD) / a
	if root < tMin || root > tMax {
		root = (-halfB-sqrtD) / a
		if root < tMin || root > tMax {
			return false
		}
	}
	
	hit.T = root
	hit.Point = r.At(hit.T)
	outwardNormal := hit.Point.Minus(sphere.Center).Div(sphere.Radius)
	hit.SetFaceNormal(&r, &outwardNormal)

	return true
}