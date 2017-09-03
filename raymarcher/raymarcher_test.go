package raymarcher

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRayDirectionAbove(t *testing.T) {
	m := Marcher{
		CameraPosition: D3{0, 4, 0},
		Target:         D3{0, 0, 0},
		Up:             D3{0, 1, 0},
		W:              100,
		H:              100,
	}

	topLeft := m.RayDirection(0, 0)
	assert.True(t, topLeft.Y < 0, "y (%f) must be negative!", topLeft.Y)
}

func TestRayDirectionRight(t *testing.T) {
	m := Marcher{
		CameraPosition: D3{4, 0, 0},
		Target:         D3{0, 0, 0},
		Up:             D3{0, 1, 0},
		W:              100,
		H:              100,
	}

	topLeft := m.RayDirection(0, 0)
	assert.True(t, topLeft.X < 0, "x (%f) must be negative!", topLeft.X)
}

func TestRayDirectionAboveRight(t *testing.T) {
	m := Marcher{
		CameraPosition: D3{4, 4, 0},
		Target:         D3{0, 0, 0},
		Up:             D3{0, 1, 0},
		W:              100,
		H:              100,
	}

	topLeft := m.RayDirection(0, 0)
	assert.True(t, topLeft.X < 0, "x (%f) must be negative!", topLeft.X)
	assert.True(t, topLeft.Y < 0, "y (%f) must be negative!", topLeft.Y)
}

func TestRayDirectionAboveIn(t *testing.T) {
	m := Marcher{
		CameraPosition: D3{0, 4, 2},
		Target:         D3{0, 0, 0},
		Up:             D3{0, 1, 0},
		W:              100,
		H:              100,
	}

	topLeft := m.RayDirection(0, 0)
	assert.True(t, topLeft.Z < 0, "z (%f) must be negative!", topLeft.Z)
	assert.True(t, topLeft.Y < 0, "y (%f) must be negative!", topLeft.Y)
}

// func TestRayDirectionAboveIn(t *testing.T) {
// 	m := Marcher{
// 		CameraPosition: D3{0, 4, 2},
// 		Target:         D3{0, 0, 0},
// 		Up:             D3{0, 1, 0},
// 		W:              100,
// 		H:              100,
// 	}
//
// 	topLeft := m.RayDirection(0, 0)
// 	assert.True(t, topLeft.X > 0, "x (%f) should be positive", topLeft.X)
// 	assert.True(t, topLeft.Y < 0, "y (%f) must be negative!", topLeft.Y)
// 	assert.True(t, topLeft.Z < 0, "z (%f) must be negative!", topLeft.Z)
//
// 	topRight := m.RayDirection(m.W, 0)
// 	assert.True(t, topRight.X < 0, "x (%f) should be negative", topRight.X)
// 	assert.True(t, topRight.Y < 0, "y (%f) must be negative!", topRight.Y)
// 	assert.True(t, topRight.Z < 0, "z (%f) must be negative!", topRight.Z)
//
// 	assert.Equal(t, topLeft.X-topLeft.X, float64(0))
//
// 	center := m.RayDirection(m.W/2, m.H/2)
// 	assert.Equal(t, center.X, float64(0))
// }
