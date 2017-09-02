package raymarcher

import (
	"math"
)

// Marcher is a ray-marcher which generates a 3D image.
type Marcher struct {
	Epsilon  float64
	MaxSteps int
	MaxDepth float64
	W        int
	H        int
	World    SDF
}

// New creates a Marcher with default parameters.
func New() *Marcher {
	return &Marcher{
		Epsilon:  0.01,
		MaxSteps: 5,
		MaxDepth: 20,
		W:        64,
		H:        64,
		World:    Empty,
	}
}

// March calculates the distance a ray travels from the origin before hitting
// a boundary. If the ray hits nothing, the distance is positive infinity.
func (m *Marcher) March(origin, ray D3) float64 {
	var distance float64 = 1.0
	var steps int = 0
	for {
		// We hit nothing.
		if distance >= m.MaxDepth || steps >= m.MaxSteps {
			return math.Inf(1)
		}

		next := origin.Add(ray.Mul(distance))
		distance += m.World(next)

		// We hit something.
		if distance < m.Epsilon {
			break
		}

		steps += 1
	}

	return distance - 1.0
}

// Ray calculates the direction of ray of light for a given pixel.
func (m *Marcher) Ray(x, y int) (ray D3) {
	ray.X = float64(x) - (float64(m.W) / 2.0)
	ray.Y = float64(y) - (float64(m.H) / 2.0)
	ray.Z = 0.0
	return
}
