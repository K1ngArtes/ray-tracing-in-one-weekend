package math

import (
	"math"
)

type Algebraic interface {
	Plus(b Algebraic) Algebraic;
}

type Vec3 struct {
	X, Y, Z float64
}

func (a *Vec3) Plus(b Algebraic) Algebraic {
	switch b.(type) {
	case *Vec3:
		return a.Add(b)
	default:
		panic("Not implemented")
	}
}

func (v *Vec3) Neg() Vec3 {
	return Vec3{
		-v.X,
		-v.Y,
		-v.Z,
	}
}

func (v *Vec3) Add(vec Vec3) Vec3 {
	return Vec3{
		v.X + vec.X,
		v.Y + vec.Y,
		v.Z + vec.Z,
	}
}

func (v *Vec3) Sub(vec Vec3) Vec3 {
	return Vec3{
		v.X - vec.X,
		v.Y - vec.Y,
		v.Z - vec.Z,
	}
}

func (v *Vec3) Mulf(num float64) Vec3 {
	return Vec3{
		v.X * num,
		v.Y * num,
		v.Z * num,
	}
}

func (v *Vec3) Mulv(vec Vec3) Vec3 {
	return Vec3{
		v.X * vec.X,
		v.Y * vec.Y,
		v.Z * vec.Z,
	}
}

func (v *Vec3) Div(num float64) Vec3 {
	return Vec3{
		v.X / num,
		v.Y / num,
		v.Z / num,
	}
}

func (v *Vec3) Dot() float64 {
	return v.X*v.X +
		v.Y*v.Y +
		v.Z*v.Z
}

func (v *Vec3) Cross(vec Vec3) Vec3 {
	return Vec3{
		v.Y*vec.Z - v.Z*vec.Y,
		v.Z*vec.X - v.X*vec.Z,
		v.X*vec.Y - v.Y*vec.X,
	}
}

func (v *Vec3) UnitVector() Vec3 {
	return v.Div(v.Len())
}

func (v *Vec3) Len() float64 {
	return math.Sqrt(v.Len_squared())
}

func (v *Vec3) Len_squared() float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

type Point3 = Vec3
type Color = Vec3
