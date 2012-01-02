package geometry

// 2 dimensional point
type Point struct {
	X, Y Scalar
} 

func (p1 Point) NearlyEquals(p2 Point) bool {
	return (p1.X - p2.X).Abs() < NearlyZero && (p1.Y - p2.Y).Abs() < NearlyZero 
}

func (p1 Point) Barycenter(p2 Point) Point {
	return Point{(p1.X + p2.X) / 2, (p1.Y + p2.Y) / 2}
}