package geometry

// A bezier of degree 3 is defined as 
// P(t) = P0*(1-t)^3 + 3*P1*t*(1-t)^2 + 3*P2*t^2*(1-t) + P3*t^3
type Curve3 [4]Vector

func (c Curve3) GetPoint(t float64) Vector {
	it := 1 - t
	it2 := it * it
	it3 := it2 * it
	t2 := t * t
	t3 := t2 * t
	x := c[0].X*it3 + 3*c[1].X*t*it2 + 3*c[2].X*t2*it + c[3].X*t3
	y := c[0].Y*it3 + 3*c[1].Y*t*it2 + 3*c[2].Y*t2*it + c[3].Y*t3
	return Vector{x, y}
}
