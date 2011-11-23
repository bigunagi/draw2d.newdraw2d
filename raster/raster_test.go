package raster

import (
	"testing"
	"log"
	"image"
	"os"
	"bufio"
	"image/png"
	"code.google.com/p/draw2d.hg/curve"
	"image/draw"
)

var flattening_threshold float64 = 0.25

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

type Path struct {
	points []float64
}

func (p *Path) LineTo(x, y float64) {
	if len(p.points)+2 > cap(p.points) {
		points := make([]float64, len(p.points)+2, len(p.points)+32)
		copy(points, p.points)
		p.points = points
	} else {
		p.points = p.points[0 : len(p.points)+2]
	}
	p.points[len(p.points)-2] = x
	p.points[len(p.points)-1] = y
}

func TestLine(t *testing.T) {
	mask := image.NewAlpha(200, 200)
	var p Path
	p.LineTo(10, 190)
	c := curve.CubicCurveFloat64{10, 190, 10, 10, 190, 10, 190, 190}
	c.Segment(&p, flattening_threshold)
	poly := Polygon(p.points)
	DrawPolyline(mask, poly...)

	img := image.NewRGBA(200, 200)
	draw.DrawMask(img, mask.Bounds(), image.NewColorImage(image.RGBAColor{0, 0, 0, 0xff}), image.ZP, mask, image.ZP, draw.Src)

	savepng("_testRasterizer.png", img)
}

func TestLineAA(t *testing.T) {
	mask := image.NewAlpha(200, 200)
	var p Path
	p.LineTo(10, 190)
	c := curve.CubicCurveFloat64{10, 190, 10, 10, 190, 10, 190, 190}
	c.Segment(&p, flattening_threshold)
	poly := Polygon(p.points)
	DrawPolylineAA(mask, poly...)
	img := image.NewRGBA(200, 200)
	draw.DrawMask(img, mask.Bounds(), image.NewColorImage(image.RGBAColor{0, 0, 0, 0xff}), image.ZP, mask, image.ZP, draw.Over)
	savepng("_testRasterizerAA.png", img)
}

/*
func TestRasterizer(t *testing.T) {
	mask := image.NewRGBA(200, 200)
	var p Path
	p.LineTo(10, 190)
	c := curve.CubicCurveFloat64{10, 190, 10, 10, 190, 10, 190, 190}
	c.Segment(&p, flattening_threshold)
	poly := Polygon(p.points)
	color := image.NRGBAColor{0, 0, 0, 0xff}
	var r Rasterizer
	//PolylineBresenham(mask, image.Black, poly...)

	r.Fill(poly, color)
	savepng("_testRasterizer.png", mask)
}
*/
