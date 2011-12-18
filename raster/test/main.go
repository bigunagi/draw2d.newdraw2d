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

type Draw struct {
	Name          string
	Width, Height int
	Polygons      []raster.Polygon
	Color         color.Color
}

func (d Draw) save() {
	tr := [6]float64{1, 0, 0, 1, 0, 0} // identity matrix
	r := raster.NewRasterizer8BitsSample(d.Width, d.Height)
	img := image.NewRGBA(image.Rect(0, 0, d.Width, d.Height))
	r.RenderEvenOdd(img, d.Color, d.Polygons, tr)
	savepng("_test" + d.Name+".png", img)
}

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


func main() {
	for _, d := range draws {
		d.save()
	}
}
