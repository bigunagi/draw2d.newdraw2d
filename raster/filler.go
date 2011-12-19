package raster

import (
	"image"
	"fmt"
)

type Intersection struct {
	x float64
	winding int8
	next *Intersection
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

func (r *Rasterizer) edge(x1, y1, x2, y2 float64) {
	var swap, dy float64
	var iy1, iy2 int
	var winding int8 = 1
	if y2 < y1 {
		swap = x1
		x1 = x2
		x2 = swap
		swap = y1
		y1 = y2
		y2 = swap
		winding = -1
	}
	iy1 = int(y1 + 0.5)
	iy2 = int(y2 + 0.5)
	dy = y2 - y1

	if dy == 0 {
		return
	}
	//idy = max(2, idy-1)

	x := x1
	dx := (x2 - x1) / dy

	for iy1 < iy2 {
		r.insert(x, iy1, winding)
		x += dx
		iy1++
	}
}

func (r *Rasterizer) insert(x float64, y int, winding int8) {
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
	var i,j *Intersection
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
					ix1 = int(i.x + 0.5)
					ix2 = int(j.x + 0.5)
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
	var i,j *Intersection
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
					ix1 = int(i.x + 0.5)
					ix2 = int(j.x + 0.5)
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
