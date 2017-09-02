package raymarcher

import (
	"math"
)

// D3 is a 3-dimensional double-based vector.
type D3 struct {
	X float64
	Y float64
	Z float64
}

// I3 is a 3-dimensional integer-based vector.
type I3 struct {
	X int32
	Y int32
	Z int32
}

// Color is an RGB color. Each dimension is within the range 0 (lowest
// intensity) to 255 (highest intensity).
type Color I3

// Length calculates the distance covered by a vector.
func (v *D3) Length() float64 {
	return math.Sqrt((v.X * v.X) + (v.Y * v.Y) + (v.Z * v.Z))
}

// Mul returns a new vector which is multiplied by a factor of m.
func (v *D3) Mul(m float64) D3 {
	return D3{
		X: v.X * m,
		Y: v.Y * m,
		Z: v.Z * m,
	}
}

// Add returns a new vector which is added to o.
func (v *D3) Add(o D3) D3 {
	return D3{
		X: v.X + o.X,
		Y: v.Y + o.Y,
		Z: v.Z + o.Z,
	}
}
