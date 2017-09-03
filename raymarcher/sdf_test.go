package raymarcher

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestEmpty(t *testing.T) {
	a := Empty(D3{1000000, 200000000, 300000000})
	assert.Equal(t, math.Inf(1), a)
}

func TestSphere(t *testing.T) {
	a := Sphere(D3{5, 0, 0}, 2)
	assert.Equal(t, float64(3), a)
}
