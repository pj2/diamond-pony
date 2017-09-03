package raymarcher

import (
	"math"
)

var (
	Up    = D3{0, 1, 0}
	Right = D3{1, 0, 0}
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
func (v D3) Length() float64 {
	return math.Sqrt((v.X * v.X) + (v.Y * v.Y) + (v.Z * v.Z))
}

// Mul returns a new vector which is multiplied by a factor of m.
func (v D3) Mul(m float64) D3 {
	return D3{
		X: v.X * m,
		Y: v.Y * m,
		Z: v.Z * m,
	}
}

func (v D3) Mulv(o D3) D3 {
	return D3{
		X: v.X * o.X,
		Y: v.Y * o.Y,
		Z: v.Z * o.Z,
	}
}

// Add returns a new vector which is added to o.
func (v D3) Add(o D3) D3 {
	return D3{
		X: v.X + o.X,
		Y: v.Y + o.Y,
		Z: v.Z + o.Z,
	}
}

// Minus returns a new vector which has o minused from it.
func (v D3) Minus(o D3) D3 {
	return D3{
		X: v.X - o.X,
		Y: v.Y - o.Y,
		Z: v.Z - o.Z,
	}
}

// Normalize returns a vector of length 1.
func (v D3) Normalize() D3 {
	l := v.Length()
	if l > 0 {
		return D3{
			X: v.X / l,
			Y: v.Y / l,
			Z: v.Z / l,
		}
	} else {
		return D3{0, 0, 0}
	}
}

// Cross calculates the cross-product of two vectors.
func (v D3) Cross(o D3) D3 {
	return D3{
		X: v.Y*o.Z - o.Y*v.Z,
		Y: o.X*v.Z - v.X*o.Z,
		Z: v.X*o.Y - o.X*v.Y,
	}
}

// Right calculates the relative right vector.
func (v D3) Right() D3 {
	return v.Cross(Up).Normalize()
}

// Up calculates the relative up vector.
func (v D3) Up() D3 {
	return v.Cross(Right).Normalize()
}
