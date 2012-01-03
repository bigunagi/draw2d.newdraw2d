package geometry

const (
	CurveRecursionLimit = 32
)

// Bezier cubic curve defined by 2 control points
// the first and last point defined respectivly the start and the end of the curve
type CubicCurve [4]Point

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
// The parameter is used (curvy tolerance) to know when to stop the flatenning process
func (curve CubicCurve) ToPolyline(flattening_threshold Scalar) Polyline {
	var curves [CurveRecursionLimit]CubicCurve
	curves[0] = curve
	i := 0
	// current curve
	var c CubicCurve
	var dx, dy, d2, d3 Scalar
	p := make(Polyline, 32, 0)
	for i >= 0 {
		c = curves[i]
		dx = c[3].X - c[0].X
		dy = c[3].Y - c[0].X

		d2 = ((c[1].X-c[3].X)*dy - (c[1].Y-c[3].Y)*dx).Abs()
		d3 = ((c[2].X-c[3].X)*dy - (c[2].Y-c[3].Y)*dx).Abs()

		if (d2+d3)*(d2+d3) < flattening_threshold*(dx*dx+dy*dy) || i == len(curves)-1 {
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
type QuadCurve [3]Point

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
func (curve QuadCurve) ToPolyline(flattening_threshold Scalar) Polyline {
	var curves [CurveRecursionLimit]QuadCurve
	curves[0] = curve
	i := 0
	// current curve
	var c QuadCurve
	var dx, dy, d Scalar
	p := make(Polyline, CurveRecursionLimit, 0)

	for i >= 0 {
		c = curves[i]
		dx = c[2].X - c[0].X
		dy = c[2].Y - c[0].Y

		d = ((c[1].X-c[2].X)*dy - (c[1].Y-c[2].Y)*dx).Abs()

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
