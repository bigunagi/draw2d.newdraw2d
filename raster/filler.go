package raster

import (
	"image"
)

type Rasterizer struct {
	img            *image.Alpha
	table1, table2 []float64
}

func NewRasterizer(width, height int) *Rasterizer {
	r := new(Rasterizer)
	r.img = image.NewAlpha(image.Rect(0, 0, width, height))
	r.table1 = make([]float64, height)
	r.table2 = make([]float64, height)
	return r
}

func (r *Rasterizer) Clear() {
	for i := 0; i < len(r.img.Pix); i++ {
		r.img.Pix[i] = 0
	}
}

func (r *Rasterizer) FillMonotone(p Polygon) {
	var xmin, ymin, xmax, ymax float64
	var ptYmin, ptYmax, i int
	xmin = p[0]
	ymin = p[1]
	xmax = xmin
	ymax = ymin
	ptYmin = 0
	ptYmax = 0
	var x, y float64
	for i = 2; i < len(p); i += 2 {
		x = p[i]
		y = p[i+1]
		if x > xmax {
			xmax = x
		} else if x < xmin {
			xmin = x
		}
		if y > ymax {
			ymax = y
			ptYmax = i
		} else if y < ymin {
			ymin = y
			ptYmin = i
		}
	}

	i = ptYmin
	j := ptYmin + 2
	if j >= len(p) {
		j = 0
	}
	for i != ptYmax {
		i = j
		j = j + 2
		if j >= len(p) {
			j = 0
		}
		r.edge(r.table1, p[i], p[i+1], p[j], p[j+1])
	}

	i = ptYmax
	j = ptYmax + 2
	if j >= len(p) {
		j = 0
	}
	for i != ptYmin {
		i = j
		j = j + 2
		if j >= len(p) {
			j = 0
		}
		r.edge(r.table2, p[i], p[i+1], p[j], p[j+1])
	}

	r.scan(int(ymin+0.5), int(ymax+0.5))

}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func (r *Rasterizer) edge(table []float64, x1, y1, x2, y2 float64) {
	var swap float64
	var idy, iy1, iy2 int
	if y2 < y1 {
		swap = x1
		x1 = x2
		x2 = swap
		swap = y1
		y1 = y2
		y2 = swap
	}
	iy1 = int(y1 + 0.5)
	iy2 = int(y2 + 0.5)
	idy = iy2 - iy1

	if idy == 0 {
		return
	}
	idy = max(2, idy-1)

	x := x1
	dx := (x2 - x1) / float64(idy)

	for iy1 < iy2 {
		table[iy1] = x
		x += dx
		iy1++
	}
}

func (r *Rasterizer) scan(ymin, ymax int) {
	var x1, x2, swap float64
	var idx, ix1, ix2 int

	pix := r.img.Pix[ymin*r.img.Stride:]
	for y := ymin; y < ymax; y++ {
		x1 = r.table1[y]
		x2 = r.table2[y]
		pix = pix[r.img.Stride:]

		if x2 < x1 {
			swap = x1
			x1 = x2
			x2 = swap
		}

		ix1 = int(x1 + 0.5)
		ix2 = int(x2 + 0.5)
		idx = ix2 - ix1

		if idx == 0 {
			continue
		}
		for ix1 < ix2 {
			pix[ix1] = 0xff
			ix1++
		}
	}
}
