// test project main.go
package main

import (
	"image"
	"image/color"
	"os"
	"log"
	"image/png"
	"bufio"
	"code.google.com/p/draw2d/raster"
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
type Figure struct {
	name string
	polygon raster.Polygon
	color color.RGBA
}

func (f Figure)save(img *image.RGBA) {
	tr := [6]float64{1, 0, 0, 1, 0, 0}
	r := raster.NewRasterizer8BitsSample(200, 200)
	r.RenderEvenOdd(img, &(f.color), f.polygon, tr)
	savepng(f.name + ".png", img)
}

func main() {
	figure := Figure{"triangle",  
			raster.Polygon{10, 110, 110, 110, 60, 10},
			color.RGBA{0, 0, 0, 0xff}}
	figure.save(image.NewRGBA(image.Rect(0, 0, 200, 200)))
	figure = Figure{"rectangle",  
			raster.Polygon{10, 10, 110, 10, 110, 110, 10, 110},
			color.RGBA{0, 0, 0, 0xff}}
	figure.save(image.NewRGBA(image.Rect(0, 0, 200, 200)))
}
