

const (
	ELLIPSE_MAGIC_NUMBER = Scalar(0.551784)
)

func (path Path) Ellipse(x Point, width, height Scalar) {
	float im = 1 - ELLIPSE_MAGIC_NUMBER;
	float rw = width / 2;
	float rh = height / 2;
	float mw = width * im;
	float mh = height * im;
	float x1 = x0 + mw;
	float y1 = y0 + mh;
	float x2 = x0 + rw;
	float y2 = y0 + rh;
	float x4 = x0 + width;
	float y4 = y0 + height;
	float x3 = x4 - mw;
	float y3 = y4 - mh;
	path.moveTo(x0, y2);
	path.cubicTo(x0, y1, x1, y0, x2, y0);
	path.cubicTo(x3, y0, x4, y1, x4, y2);
	path.cubicTo(x4, y3, x3, y4, x2, y4);
	path.cubicTo(x1, y4, x0, y3, x0, y2);
}


func (path Path) arcTo(float x1, float y1, float x2, float y2, float radius) {
	// some maths (see figure XXX)
	// let P0(x0, y0), P1(x1, y1), P2(x2, y2)
	// computes coords of P0 and P2 relative to P1:
	// (note: all _ terminated coordinates are relative to P1)
	float x0_ = x0 - x1, y0_ = y0 - y1; // coord of P0 relative to P1
	float x2_ = x2 - x1, y2_ = y2 - y1; // coord of P2 relative to P1
	// compute the two unit vectors:
	float d0 = (float) Math.sqrt(x0_ * x0_ + y0_ * y0_);
	if (d0 == 0) {
		lineTo(x1, y1);
		return;
	}
	float d2 = (float) Math.sqrt(x2_ * x2_ + y2_ * y2_);
	if (d2 == 0) {
		lineTo(x1, y1);
		return;
	}
	float xv0 = x0_ / d0, yv0 = y0_ / d0;
	float xv2 = x2_ / d2, yv2 = y2_ / d2;
	// compute the two normal unit vectors:
	float xn0 = - yv0, yn0 = xv0;
	float xn2 = - yv2, yn2 = xv2;
	// if necessary, reverse the orientation of the normal unit vectors, so that each vector point in the direction of the center of the circle:
	if (xn0 * x2_ + yn0 * y2_ < 0) {
		xn0 = - xn0; yn0 = - yn0;
	}
	if (xn2 * x0_ + yn2 * y0_ < 0) {
		xn2 = - xn2; yn2 = - yn2;
	}
	// compute  C (center of the circle);
	// a simple system of two equation must be solved:
	float det = xn0 * yn2 - xn2 * yn0;
	if (det == 0) {
		lineTo(x1, y1);
		return;
	}
	float q = radius / det;
	float xc_ = (yn2 - yn0) * q;
	float yc_ = (xn0 - xn2) * q;
//		float xc = x1 + xc_;
//		float yc = y1 + yc_;
	// compute the length of the segments P1Q0 = P1Q2, using a scalar product:
	float delta = xv0 * xc_ + yv0 * yc_;
	// compute Q0 and Q2:
	float xq0 = x1 + xv0 * delta;
	float yq0 = y1 + yv0 * delta;
	float xq2 = x1 + xv2 * delta;
	float yq2 = y1 + yv2 * delta;
	// end maths.
	
	// now we can draw:
	float magic = 0.55228475f; // see http://www.tinaja.com/glib/ellipse4.pdf
	float dd = delta * (1 - magic);
	path.lineTo(xq0, yq0);
	path.cubicTo(x1 + xv0 * dd, y1 + yv0 * dd, x1 + xv2 * dd, y1 + yv2 * dd, xq2, yq2);
	x0 = xq2;
	y0 = yq2;
}