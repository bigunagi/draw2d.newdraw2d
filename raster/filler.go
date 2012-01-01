package raster

import (
	"fmt"
	"image"
)

type Intersection struct {
	x       int
	winding int8
	next    *Intersection
}

type Rasterizer struct {
	table []*Intersection
}

func NewRasterizer() *Rasterizer {
	r := new(Rasterizer)
	return r
}

func (r *Rasterizer) Fill(img *image.Alpha, p Polygon, nonZeroWindingRule bool) {
	r.table = make([]*Intersection, img.Bounds().Dy())
	var xmin, ymin, xmax, ymax float64
	xmin = p[0]
	ymin = p[1]
	xmax = xmin
	ymax = ymin
	var x, y float64
	for i := 2; i < len(p); i += 2 {
		x, y = p[i], p[i+1]
		if x > xmax {
			xmax = x
		} else if x < xmin {
			xmin = x
		}
		if y > ymax {
			ymax = y
		} else if y < ymin {
			ymin = y
		}
	}
	prevX, prevY := p[0], p[1]
	for i := 2; i < len(p); i += 2 {
		x, y = p[i], p[i+1]
		r.edge(prevX, prevY, x, y)
		prevX, prevY = x, y
	}

	if nonZeroWindingRule {
		r.scanNonZero(img, int(ymin+0.5), int(ymax+0.5))
	} else {
		r.scanEvenOdd(img, int(ymin+0.5), int(ymax+0.5))
	}
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func (r *Rasterizer) edge(xf1, yf1, xf2, yf2 float64) {
	x1, y1, x2, y2 := int(xf1+0.5), int(yf1+0.5), int(xf2+0.5), int(yf2+0.5)
	dx := abs(x2 - x1)
	dy := abs(y2 - y1)
	if dy == 0 {
		return
	}
	var sx, sy int
	if x1 < x2 {
		sx = 1
	} else {
		sx = -1
	}
	if y1 < y2 {
		sy = 1
	} else {
		sy = -1
	}
	err := dx - dy
	r.insert(x1, y1, int8(sy))
	var e2 int
	for {
		if x1 == x2 && y1 == y2 {
			return
		}
		e2 = 2 * err
		if e2 > -dy {
			err = err - dy
			x1 = x1 + sx
		}
		if e2 < dx {
			err = err + dx
			y1 = y1 + sy
			r.insert(x1, y1, int8(sy))
		}
	}
}

func (r *Rasterizer) insert(x int, y int, winding int8) {
	i := &Intersection{x, winding, nil}
	if r.table[y] == nil {
		r.table[y] = i
		return
	}
	var prev *Intersection
	current := r.table[y]
	for current != nil && x > current.x {
		prev = current
		current = current.next
	}
	i.next = current
	if prev != nil {
		prev.next = i
		return
	}
}

func printIntersection(i *Intersection) {
	if i == nil {
		fmt.Print("nil")
	} else {
		for i != nil {
			fmt.Print(i.x, " ")
			i = i.next
		}
	}
	fmt.Println()
}

func (r *Rasterizer) scanEvenOdd(img *image.Alpha, ymin, ymax int) {
	var idx, ix1, ix2 int
	var i, j *Intersection
	fill := true
	pix := img.Pix[ymin*img.Stride:]
	for y := ymin; y < ymax; y++ {
		pix = pix[img.Stride:]
		i = r.table[y]
		if i != nil {
			fill = true
			j = i.next
			for j != nil {
				if fill {
					ix1 = i.x
					ix2 = j.x
					idx = ix2 - ix1
					if idx == 0 {
						continue
					}
					for ix1 < ix2 {
						pix[ix1] = 0xff
						ix1++
					}
				}
				fill = !fill
				i = j
				j = i.next
			}
		}
	}
}

func (r *Rasterizer) scanNonZero(img *image.Alpha, ymin, ymax int) {
	var ix1, ix2 int
	var i, j *Intersection
	pix := img.Pix[ymin*img.Stride:]
	var winding int8 = 0
	for y := ymin; y < ymax; y++ {
		pix = pix[img.Stride:]
		i = r.table[y]
		if i != nil {
			winding = i.winding
			j = i.next
			for j != nil {
				if winding != 0 {
					ix1 = i.x
					ix2 = j.x
					for ix1 < ix2 {
						pix[ix1] = 0xff
						ix1++
					}
				}
				winding = winding + j.winding
				i = j
				j = i.next
			}
		}
	}
}
