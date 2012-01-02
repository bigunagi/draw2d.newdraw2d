package geometry

// 2 dimensional point
type Point struct {
	X, Y Scalar
} 

func (p1 Point) NearlyEquals(p2 Point) bool {
	return (p1.X - p2.X).Abs() < NearlyZero && (p1.Y - p2.Y).Abs() < NearlyZero 
}

// A Line is defined by two point
type Line struct {
	P1, P2 Point
}

// 
type Polyline struct {
	vertices []Point
}

// This function append vertices of polyline in parameter and return the new polyline 
// that is the concatenation of the two polylines
func (p1 Polyline) Concat(p2 Polyline) Polyline {
	var p Polyline
	p.vertices = make([]Point, len(p1.vertices) + len(p2.vertices))
	copy(p.vertices,p1.vertices)
	copy(p.vertices[len(p1.vertices):], p2.vertices)
	return p
}

// Return the number of lines
func (p Polyline) GetLineCount() int {
	return len(p.vertices) - 1
}

// Return nth line of this polyline index begin at 0
func (p Polyline) GetLine(index int) Line {
	i := index * 2
	return Line{p.vertices[i], p.vertices[i+1]}
}

// Return the number of vertices
// Vertices count == Line count
func (p Polyline) GetVertexCount() int {
	return len(p.vertices)
}

// Return nth vertex of this polyline index begin at 0
func (p Polyline) GetVertex(index int) Point {
	return p.vertices[index]
}
// Return vertices of polyline
func (p Polyline) GetVertices() []Point {
	return p.vertices
}


// Close the polyline to make a polygon
func (p Polyline) ToPolygon() Polygon {
	polygon := Polygon{p.vertices}
	if p.vertices[0].NearlyEquals(p.vertices[len(p.vertices) - 1]) {
		return polygon
	}
	if cap(p.vertices) < len(p.vertices) {
		vertices := make([]Point, len(p.vertices)+2)
		copy(vertices, p.vertices)
		polygon.vertices = vertices
	}
	polygon.vertices = polygon.vertices[:len(p.vertices) + 2]
	// close the polyline
	polygon.vertices[len(polygon.vertices) - 2] = polygon.vertices[0]
	polygon.vertices[len(polygon.vertices) - 1] = polygon.vertices[1]
	return polygon
}

// A polygon is a closed path composed by straigth line segments called edges. 
// Edges are joined by vertices (Point).
// this is a 2 dimensional polygon 
type Polygon Polyline

type PolylineConverter interface {
	AsPolyline() Polyline
}
