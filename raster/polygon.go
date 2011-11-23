package raster

type Polygon []float64


func (p Polygon) ComputeBoundingRect() (xmin, ymin, xmax, ymax float64, ptXmin, ptYmin, ptXmax, ptYmax int) {
	xmin = p[0]; ymin = p[1]
	xmax = xmin; ymax = ymin
	ptYmin = 0 ; ptYmax = 0
	var x, y float64
	for i := 2; i < len(p); i += 2 {
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
	return 
}


