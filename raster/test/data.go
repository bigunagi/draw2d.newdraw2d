package main

import (
	"code.google.com/p/draw2d/raster"
	"image/color"
)


var (
	draws = []Draw{
		{	
			"Triangle", 120, 120,
			[]raster.Polygon{{10, 110, 110, 110, 60, 10}},
			color.NRGBA{0, 0, 0, 0xff},
		},
		{
			"Rectangle", 120, 120,
			[]raster.Polygon{{10, 10, 110, 10, 110, 110, 10, 110}},
			color.RGBA{0, 0, 0, 0xff},
		},
		{
			"Rectangle2", 120, 120,
			[]raster.Polygon{{0, 0, 120, 0, 120, 120, 0, 120}},
			color.RGBA{0, 0, 0, 0xff},
		},
		{	
			"Triangle2", 120, 120,
			[]raster.Polygon{{-10, 110, 130, 110, 60, -10}},
			color.NRGBA{0, 0, 0, 0xff},
		},
	}
)