package raymarcher

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

const (
	EPSILON = 0.001
)

func TestLength(t *testing.T) {
	v := D3{3, 4, 5}
	assert.InEpsilon(t, 7.07, v.Length(), EPSILON)
}

func TestMul(t *testing.T) {
	v := D3{2, 3, 4}
	assert.Equal(t, v.Mul(2), D3{4, 6, 8})
}

func TestAdd(t *testing.T) {
	v := D3{1, 2, 3}
	assert.Equal(t, v.Add(D3{1, 1, 1}), D3{2, 3, 4})
}

func TestMinus(t *testing.T) {
	v := D3{0, 0, 0}
	assert.Equal(t, v.Minus(D3{1, 2, 3}), D3{-1, -2, -3})
}

func TestNormalize(t *testing.T) {
	v := D3{51, 35, 90}
	v = v.Normalize()
	assert.InEpsilon(t, 0.467006, v.X, EPSILON)
	assert.InEpsilon(t, 0.320495, v.Y, EPSILON)
	assert.InEpsilon(t, 0.824129, v.Z, EPSILON)
}

func TestNormalizeZero(t *testing.T) {
	v := D3{0, 0, 0}
	v = v.Normalize()
	assert.False(t, math.IsNaN(v.X), "div by zero")
}

func TestCross(t *testing.T) {
	a := D3{51, 35, 90}
	b := D3{71, 53, 100}
	v := a.Cross(b)

	assert.Equal(t, v, D3{-1270, 1290, 218})
}

func TestRight(t *testing.T) {
	v := D3{0, 4, -2}
	assert.Equal(t, D3{1, 0, 0}, v.Right())
}

// func TestUp(t *testing.T) {
// 	v := D3{0, 4, -2}
// 	assert.Equal(t, D3{1, 0, 0}, v.Up())
// }
