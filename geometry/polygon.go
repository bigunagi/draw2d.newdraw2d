package geometry

// 
type Polyline []Point

// This function append vertices of polyline in parameter and return the new polyline 
// that is the concatenation of the two polylines
func (p1 Polyline) Concat(p2 Polyline) Polyline {
	p := make(Polyline, len(p1) + len(p2), cap(p1) + cap(p2))
	copy(p, p1)
	copy(p[len(p1):], p2)
	return p
}

// Return the number of lines
func (p Polyline) GetLineCount() int {
	return len(p) - 1
}

// Return nth line of this polyline index begin at 0
func (p Polyline) GetLine(index int) (p1, p2 Point) {
	i := index * 2
	return p[i], p[i+1]
}

// Return the number of vertices
// Vertices count == Line count
func (p Polyline) GetVertexCount() int {
	return len(p)
}

// Return nth vertex of this polyline index begin at 0
func (p Polyline) GetVertex(index int) Point {
	return p[index]
}

// Return vertices of polyline
func (p Polyline) GetVertices() []Point {
	return p
}


// Close the polyline to make a polygon
func (p Polyline) ToPolygon() Polygon {
	var polygon Polygon
	if p[0].NearlyEquals(p[len(p) - 1]) {
		polygon := make(Polygon, len(p))
		copy(polygon, p)
		return polygon
	}
	if cap(p) < len(p) {
		polygon := make(Polygon, len(p)+2)
		copy(polygon, p)
	}
	polygon = polygon[:len(p) + 2]
	// close the polyline
	polygon[len(polygon) - 2] = polygon[0]
	polygon[len(polygon) - 1] = polygon[1]
	return polygon
}

// A polygon is a closed path composed by straigth line segments called edges. 
// Edges are joined by vertices (Point).
// this is a 2 dimensional polygon 
type Polygon Polyline

type PolylineConverter interface {
	ToPolyline() Polyline
}
