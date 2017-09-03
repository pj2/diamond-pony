package raymarcher

import (
	"math"
	"time"
)

var (
	NoPixel = D3{
		X: math.Inf(1),
		Y: math.Inf(1),
		Z: math.Inf(1),
	}
)

// Marcher is a ray-marcher which generates a 3D image.
type Marcher struct {
	Epsilon        float64
	MaxSteps       int
	MaxDepth       float64
	W              int
	H              int
	World          SDF
	CameraPosition D3
	Up             D3
	Target         D3
}

// Pixel is the result of a ray march.
type Pixel struct {
	X        int
	Y        int
	Position D3
}

// New creates a Marcher with default parameters.
func New() *Marcher {
	return &Marcher{
		Epsilon:  0.01,
		MaxSteps: 5,
		MaxDepth: 20,
		W:        300,
		H:        300,
		World:    Empty,
		CameraPosition: D3{
			X: 5,
			Y: 0,
			Z: 0,
		},
		Up:     D3{0.0, 1.0, 0.0},
		Target: D3{0.0, 0.0, 0.0},
	}
}

// March calculates the position (and distance) at which a ray collides with
// some geometry. If the ray hits nothing, the distance is positive infinity.
func (m *Marcher) March(origin, r D3) D3 {
	var next D3
	var distance float64 = 0.0
	var totalDistance float64 = 0.0
	var steps int = 0

	for {
		// We hit nothing.
		if totalDistance >= m.MaxDepth || steps >= m.MaxSteps {
			return NoPixel
		}

		next = origin.Add(r.Mul(distance))
		distance = m.World(next)
		totalDistance += distance

		// We hit something.
		if distance < m.Epsilon {
			break
		}

		steps++
	}

	return next
}

// RayDirection calculates the direction of ray of light for a given pixel.
func (m *Marcher) RayDirection(x, y int) D3 {
	dir := m.Target.Minus(m.CameraPosition).Normalize()
	// right := dir.Right()
	// up := dir.Up()
	right := D3{0, 0, 1} // TODO Don't hardcode :D
	up := D3{0, 1, 0}

	// Clamp to from -1.0 to 1.0 in local space
	mulX := -1.0 + (2.0 * (float64(x) / float64(m.W)))
	mulY := -1.0 + (2.0 * (float64(y) / float64(m.H)))

	// Map from screen-space to local-space
	offsetX := right.Mul(mulX)
	offsetY := up.Mul(mulY)

	ray := dir
	ray = ray.Add(offsetX)
	ray = ray.Add(offsetY)
	return ray.Normalize()
}

// Pixels produces the march algorithm output for all pixels continually
// until quit.
func (m *Marcher) Pixels(c chan Pixel, tick <-chan time.Time,
	quit chan struct{}) {
	for {
		select {
		case <-quit:
			break
		case <-tick:
			for y := 0; y < m.H; y++ {
				for x := 0; x < m.W; x++ {
					c <- Pixel{
						X: x,
						Y: y,
						Position: m.March(m.CameraPosition,
							m.RayDirection(x, y)),
					}
				}
			}
		}
	}
}

// Empty returns true if a pixel hit nothing.
func (p *Pixel) Empty() bool {
	return math.IsInf(p.Position.X, 0)
}
