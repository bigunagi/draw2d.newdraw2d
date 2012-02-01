package geometry

const (
	NearlyZero float64 = 1e-12
)

// Clamp f to be between min and max
func Clamp(s, min, max float64) float64 {
	if s < min {
		return min
	}
	if s > max {
		return max
	}
	return s
}

// Linearly interpolate between a and b by t percent (comprise from 0 to 1).
func Lerp(a, b, t float64) float64 {
	return a*(1.0-t) + b*t
}
