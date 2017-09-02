package raymarcher

// SDF is a signed-distance function. It is a primary building block of a
// raymarcher, defining the world in terms of distances. It is used within the
// marching algorithm to determine where a ray collides (and from that, the
// ray's corresponding pixel color).
type SDF func(D3) float64

// Empty represents a lack of volume.
func Empty(position D3) float64 {
	return -1.0
}

// Sphere returns the signed distance to the boundary of a spherical volume.
func Sphere(position D3, radius float64) float64 {
	return position.Length() - radius
}
