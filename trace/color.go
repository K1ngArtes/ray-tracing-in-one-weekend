package trace

import "github.com/K1ngArtes/ray-tracing-in-one-weekend/math"

var (
	black = Color{0, 0, 0}
	white = Color{1, 1, 1}
)

type Color math.Vec3

// R returns the first element (Red).
func (c Color) R() float64 {
	return c[0]
}

// G returns the second element (Green).
func (c Color) G() float64 {
	return c[1]
}

// B returns the third element (Blue).
func (c Color) B() float64 {
	return c[2]
}
