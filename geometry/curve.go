package geometry

const (
	CurveRecursionLimit = 32
)

//	X1, Y1, X2, Y2, X3, Y3, X4, Y4 float64
type CubicCurve [4]Point

type LineTracer interface {
	LineTo(x, y Scalar)
}

func (c *CubicCurve) Subdivide() (c1, c2 CubicCurve) {
	// Calculate all the mid-points of the line segments
	//----------------------
	c1[0] = c[0]
	c2[3] = c[3]
	c1[1].X = (c[0].X + c[1].X) / 2
	c1[1].Y = (c[0].Y + c[1].Y) / 2
	x23 := (c[1].X + c[2].X) / 2
	y23 := (c[1].Y + c[2].Y) / 2
	c2[2].X = (c[2].X + c[3].X) / 2
	c2[2].Y = (c[2].Y + c[3].Y) / 2
	c1[2].X = (c1[1].X + x23) / 2
	c1[2].Y = (c1[1].Y + y23) / 2
	c2[1].X = (x23 + c2[2].X) / 2
	c2[1].Y = (y23 + c2[2].Y) / 2
	c1[3].X = (c1[2].X + c2[1].X) / 2
	c1[3].Y = (c1[2].Y + c2[1].Y) / 2
	c2[0] = c1[3]
	return
}

func (curve CubicCurve) ToPolyline() Polyline {
	var flattening_threshold Scalar = 0.5
	var curves [CurveRecursionLimit]CubicCurve
	curves[0] = curve
	i := 0
	// current curve
	var c CubicCurve

	var dx, dy, d2, d3 Scalar
	points := make([]Point, 32)

	count := 0

	for i >= 0 {
		c = curves[i]
		dx = c[3].X - c[0].X
		dy = c[3].Y - c[0].X

		d2 = ((c[1].X-c[3].X)*dy - (c[1].Y-c[3].Y)*dx).Abs()
		d3 = ((c[2].X-c[3].X)*dy - (c[2].Y-c[3].Y)*dx).Abs()

		if (d2+d3)*(d2+d3) < flattening_threshold*(dx*dx+dy*dy) || i == len(curves)-1 {
			points[count] = c[3]
			count++
			i--
		} else {
			// second half of bezier go lower onto the stack
			curves[i+1], curves[i] = c.Subdivide()
			i++
		}
	}
	return Polyline{points[:count]}
}

func (curve *CubicCurve) Segment(t LineTracer, flattening_threshold Scalar) {
	var curves [CurveRecursionLimit]CubicCurve
	curves[0] = *curve
	i := 0
	// current curve
	var c *CubicCurve

	var dx, dy, d2, d3 Scalar

	for i >= 0 {
		c = &curves[i]
		dx = c[3].X - c[0].X
		dy = c[3].Y - c[0].X

		d2 = ((c[1].X-c[3].X)*dy - (c[1].Y-c[3].Y)*dx).Abs()
		d3 = ((c[2].X-c[3].X)*dy - (c[2].Y-c[3].Y)*dx).Abs()

		if (d2+d3)*(d2+d3) < flattening_threshold*(dx*dx+dy*dy) || i == len(curves)-1 {
			t.LineTo(c[3].X, c[3].Y)
			i--
		} else {
			// second half of bezier go lower onto the stack
			curves[i+1], curves[i] = c.Subdivide()
			i++
		}
	}
}
/*
//X1, Y1, X2, Y2, X3, Y3 float64
type QuadCurve [3]Point

func (c *QuadCurve) Subdivide(c1, c2 *QuadCurve) {
	// Calculate all the mid-points of the line segments
	//----------------------
	c1[0], c1[1] = c[0], c[1]
	c2[4], c2[5] = c[4], c[5]
	c1[2] = (c[0] + c[2]) / 2
	c1[3] = (c[1] + c[3]) / 2
	c2[2] = (c[2] + c[4]) / 2
	c2[3] = (c[3] + c[5]) / 2
	c1[4] = (c1[2] + c2[2]) / 2
	c1[5] = (c1[3] + c2[3]) / 2
	c2[0], c2[1] = c1[4], c1[5]
	return
}

func (curve *QuadCurve) Segment(t LineTracer, flattening_threshold float64) {
	var curves [CurveRecursionLimit]QuadCurve
	curves[0] = *curve
	i := 0
	// current curve
	var c *QuadCurve
	var dx, dy, d float64

	for i >= 0 {
		c = &curves[i]
		dx = c[4] - c[0]
		dy = c[5] - c[1]

		d = math.Abs(((c[2]-c[4])*dy - (c[3]-c[5])*dx))

		if (d*d) < flattening_threshold*(dx*dx+dy*dy) || i == len(curves)-1 {
			t.LineTo(c[4], c[5])
			i--
		} else {
			// second half of bezier go lower onto the stack
			c.Subdivide(&curves[i+1], &curves[i])
			i++
		}
	}
}
*/
