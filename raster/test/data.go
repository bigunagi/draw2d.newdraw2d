package main

import (
	. "code.google.com/p/draw2d/raster"
	"image/color"
)

var (
	draws = []Draw{
		{
			"Triangle", 120, 120,
			[]Polygon{{10, 110, 110, 110, 60, 10}},
			color.NRGBA{0, 0, 0, 0xff},
		},
		{
			"Rectangle", 120, 120,
			[]Polygon{{10, 10, 110, 10, 110, 110, 10, 110}},
			color.RGBA{0, 0, 0, 0xff},
		},

		{
			"Star", 120, 120,
			[]Polygon{{60, 10, 30, 110, 110, 50, 10, 50, 90, 110}},
			color.NRGBA{0, 0, 0, 0xff},
		},
	}
)
