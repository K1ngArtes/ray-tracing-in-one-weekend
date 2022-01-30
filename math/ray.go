package math

type Ray struct {
	origin Point3
	dir    Vec3
}

func (r *Ray) At(t float64) Point3 {
	return r.origin.Add(r.dir.Mulf(t))
}
