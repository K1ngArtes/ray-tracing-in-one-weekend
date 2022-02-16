package geom

import "math"

type Algebraic interface {
	Plus(b Algebraic) Algebraic
}

type Vec3 [3]float64

func (v Vec3) X() float64 {
	return v[0]
}

func (v Vec3) Y() float64 {
	return v[1]
}

func (v Vec3) Z() float64 {
	return v[2]
}

func (v Vec3) Inv() Vec3 {
	return Vec3{-v[0], -v[1], -v[2]}
}

func (v Vec3) Len() float64 {
	return math.Sqrt(v.LenSq())
}

func (v Vec3) LenSq() float64 {
	return v[0]*v[0] + v[1]*v[1] + v[2]*v[2]
}

func (v Vec3) Plus(v2 Vec3) Vec3 {
	return Vec3{
		v[0] + v2[0],
		v[1] + v2[1],
		v[2] + v2[2],
	}
}

func (v Vec3) Minus(v2 Vec3) Vec3 {
	return Vec3{
		v[0] - v2[0],
		v[1] - v2[1],
		v[2] - v2[2],
	}
}

func (v Vec3) Times(v2 Vec3) Vec3 {
	return Vec3{
		v[0] * v2[0],
		v[1] * v2[1],
		v[2] * v2[2],
	}
}

func (v Vec3) Div(n float64) Vec3 {
	return Vec3{
		v[0] / n,
		v[1] / n,
		v[2] / n,
	}
}

func (v Vec3) Scaled(n float64) Vec3 {
	return Vec3{
		v[0] * n,
		v[1] * n,
		v[2] * n,
	}
}

func (v Vec3) Dot(v2 Vec3) float64 {
	return v[0]*v2[0] + v[1]*v2[1] + v[2]*v2[2]
}

func (v Vec3) Cross(v2 Vec3) Vec3 {
	return Vec3{
		v[1]*v2[2] - v[2]*v2[1],
		v[2]*v2[0] - v[0]*v2[2],
		v[0]*v2[1] - v[1]*v2[0],
	}
}

func (v Vec3) Unit() (v2 Vec3) {
	k := 1.0 / v.Len()
	v2[0] = v[0] * k
	v2[1] = v[1] * k
	v2[2] = v[2] * k
	return
}

func (v Vec3) Negate() Vec3 {
	return Vec3{
		-v[0],
		-v[1],
		-v[2],
	}
}
