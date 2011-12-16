// test project main.go
package main

import (
	"bufio"
	"code.google.com/p/draw2d/raster"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

var (
	draws = []Draw{
		{	
			"Triangle", 120, 120,
			[]raster.Polygon{{10, 110, 110, 110, 60, 10}},
			color.RGBA{0, 0, 0, 0xff},
		},
		{
			"Rectangle", 120, 120,
			[]raster.Polygon{{10, 10, 110, 10, 110, 110, 10, 110}},
			color.RGBA{0, 0, 0, 0xff},
		},
	}
)

func savepng(filePath string, m image.Image) {
	f, err := os.Create(filePath)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer f.Close()
	b := bufio.NewWriter(f)
	err = png.Encode(b, m)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	err = b.Flush()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

type Draw struct {
	Name          string
	Width, Height int
	Polygons      []raster.Polygon
	Color         color.RGBA
}

func (d Draw) save() {
	tr := [6]float64{1, 0, 0, 1, 0, 0} // identity matrix
	r := raster.NewRasterizer8BitsSample(d.Width, d.Height)
	img := image.NewRGBA(image.Rect(0, 0, d.Width, d.Height))
	r.RenderEvenOdd(img, &(d.Color), d.Polygons, tr)
	savepng("_test" + d.Name+".png", img)
}

func main() {
	for _, d := range draws {
		d.save()
	}
}
