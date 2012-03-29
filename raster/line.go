// Copyright 2011 The draw2d Authors. All rights reserved.
// created: 27/05/2011 by Laurent Le Goff
package raster

import (
	"image"
	//"math"
	//"fmt"
)

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func DrawPolyline(img *image.Alpha, s ...float64) {
	var x1, y1, x2, y2 int
	x1 = int(s[0] + 0.5)
	y1 = int(s[1] + 0.5)
	for i := 2; i < len(s); i += 2 {
		x2 = int(s[i] + 0.5)
		y2 = int(s[i+1] + 0.5)
		DrawLine(img, x1, y1, x2, y2)
		x1 = x2
		y1 = y2
	}
}

// Uses Bresenham
func DrawLine(img *image.Alpha, x0, y0, x1, y1 int) {
	//fmt.Printf("Draw line: (%d, %d) (%d, %d)\n", x0, y0, x1, y1)
	dx := abs(x1 - x0)
	dy := abs(y1 - y0)
	var sx, sy int
	if x0 < x1 {
		sx = 1
	} else {
		sx = -1
	}
	if y0 < y1 {
		sy = 1
	} else {
		sy = -1
	}
	err := dx - dy

	y0 = y0 * img.Stride
	y1 = y1 * img.Stride
	var e2 int
	for {

		img.Pix[y0+x0] = 0xff
		if x0 == x1 && y0 == y1 {
			return
		}
		e2 = 2 * err
		if e2 > -dy {
			err = err - dy
			x0 = x0 + sx
		}
		if e2 < dx {
			err = err + dx
			y0 = y0 + sy*img.Stride
		}
	}
}

func fpart(x float64) float64 {
	return x - float64(int(x+1)-1)
}

func rfpart(x float64) float64 {
	return 1 - fpart(x)
}

func plot(img *image.Alpha, x, y int, brightness float64) {
	// for now do nothing
	img.Pix[y*img.Stride+x] = uint8(0xff * brightness)
}

func DrawPolylineAA(img *image.Alpha, s ...float64) {
	var x1, y1, x2, y2 float64
	x1 = s[0]
	y1 = s[1]
	for i := 2; i < len(s); i += 2 {
		x2 = s[i]
		y2 = s[i+1]
		DrawLineAA(img, x1, y1, x2, y2)
		x1 = x2
		y1 = y2
	}
}

// Uses Wu-antialiasing
func DrawLineAA(img *image.Alpha, x1, y1, x2, y2 float64) {
	//var swap float64
	dx := x2 - x1
	dy := y2 - y1
	/*if math.Fabs(dx) < math.Fabs(dy) {
			swap = x1
			x1 = y1
			y1 = swap
			swap = x2
			x2 = y2
			y2 = swap
			swap = dx
			dx = dy
			dy = swap
	    }
	    if x2 < x1 {
			swap = x1
			x1 = x2
			x2 = swap
			swap = y1
			y1 = y2
			y2 = swap
	    }*/
	gradient := dy / dx

	// handle first endpoint
	xend := int(x1 + 0.5) // round
	yend := y1 + gradient*(float64(xend)-x1)

	xgap := rfpart(x1 + 0.5) // reverse fraction part 
	xpxl1 := xend            // this will be used in the main loop
	ypxl1 := int(yend)
	yfparth := yend - float64(int(yend+1)-1) // fpart yend
	yrfparth := 1 - yfparth                  // rfpart yend
	plot(img, xpxl1, ypxl1, yrfparth*xgap)
	plot(img, xpxl1, ypxl1+1, yfparth*xgap)
	intery := yend + gradient // first y-intersection for the main loop

	// handle second endpoint
	xend = int(x2 + 0.5) // round
	yend = y2 + gradient*(float64(xend)-x2)
	xgap = fpart(x2 + 0.5)
	xpxl2 := xend                           // this will be used in the main loop
	ypxl2 := int(yend)                      // floor
	yfparth = yend - float64(int(yend+1)-1) // fpart yend
	yrfparth = 1 - yfparth                  // rfpart yend
	plot(img, xpxl2, ypxl2, yrfparth*xgap)
	plot(img, xpxl2, ypxl2+1, yfparth*xgap)

	// main loop
	for x := xpxl1 + 1; x < xpxl2; x++ {
		yfparth = intery - float64(int(intery+1)-1) // fpart intery
		yrfparth = 1 - yfparth                      // rfpart intery
		plot(img, x, int(intery), yrfparth)
		plot(img, x, int(intery)+1, yrfparth)
		intery = intery + gradient
	}
}
