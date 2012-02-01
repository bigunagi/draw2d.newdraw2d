package geometry

import (
	"math"
)

const (
	curveRecursionLimit = 32
)

// Bezier cubic curve defined by 2 control points
// the first and last point defined respectivly the start and the end of the curve
type CubicCurve [4]Vector

// Subdivide the curve into 2 CubicCurve that is equivalent
func (c CubicCurve) Subdivide() (c1, c2 CubicCurve) {
	// Calculate all the mid-points of the line segments
	//----------------------
	c1[0] = c[0]
	c2[3] = c[3]
	p := c[1].Center(c[2])
	c1[1] = c[0].Center(c[1])
	c2[2] = c[2].Center(c[3])
	c1[2] = c1[1].Center(p)
	c2[1] = c2[2].Center(p)
	c1[3] = c1[2].Center(c2[1])
	c2[0] = c1[3]
	return
}

// Flatten Curve into a Polyline
// flatteningThreshold is used to stop subdivision of curve. Give good result with the value 0.25 
// see http://www.antigrain.com/research/adaptive_bezier/index.html
func (curve CubicCurve) ToPolyline(flatteningThreshold float64) Polyline {
	var curves [curveRecursionLimit]CubicCurve
	curves[0] = curve
	i := 0
	// current curve
	var c CubicCurve
	var dx, dy, d2, d3 float64
	p := make(Polyline, 32, 0)
	for i >= 0 {
		c = curves[i]
		d := c[3].Sub(c[0])
		dn := d.Normal()

		d2 := math.Abs(c[1].Sub(c[3]).Dot(dn)) + math.Abs(c[2].Sub(c[3]).Dot(dn))

		if d2*d2 < flatteningThreshold*d.LengthSquare() || i == len(curves)-1 {
			p = append(p, c[3])
			i--
		} else {
			// second half of bezier go lower onto the stack
			curves[i+1], curves[i] = c.Subdivide()
			i++
		}
	}
	return p
}

// Bezier quadratic curve defined by 1 control point
// the first and last point defined respectivly the start and the end of the curve
type QuadCurve [3]Vector

// Subdivide the curve into 2 QuadCurve that is equivalent
func (c QuadCurve) Subdivide() (c1, c2 QuadCurve) {
	// Calculate all the mid-points of the line segments
	//----------------------
	c1[0] = c[0]
	c2[2] = c[2]
	c1[1] = c[0].Center(c[1])
	c2[1] = c[1].Center(c[2])
	c1[2] = c1[1].Center(c2[1])
	c2[0] = c1[2]
	return
}

// Flatten Curve into a Polyline
// The parameter is used (curvy tolerance) to know when to stop the flatenning process
func (curve QuadCurve) ToPolyline(flattening_threshold float64) Polyline {
	var curves [curveRecursionLimit]QuadCurve
	curves[0] = curve
	i := 0
	// current curve
	var c QuadCurve
	var dx, dy, d float64
	p := make(Polyline, curveRecursionLimit, 0)

	for i >= 0 {
		c = curves[i]
		dx = c[2].X - c[0].X
		dy = c[2].Y - c[0].Y

		d = math.Abs((c[1].X-c[2].X)*dy - (c[1].Y-c[2].Y)*dx)

		if (d*d) < flattening_threshold*(dx*dx+dy*dy) || i == len(curves)-1 {
			p = append(p, c[2])
			i--
		} else {
			// second half of bezier go lower onto the stack
			curves[i+1], curves[i] = c.Subdivide()
			i++
		}
	}
	return p
}
