package raster

import (
	"bufio"
	"code.google.com/p/draw2d/curve"
	"code.google.com/p/freetype-go/freetype/raster"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"testing"
)

var flattening_threshold float64 = 0.5

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

func TestFreetype(t *testing.T) {
	var p Path
	p.LineTo(10, 190)
	c := curve.CubicCurveFloat64{10, 190, 10, 10, 190, 10, 190, 190}
	c.Segment(&p, flattening_threshold)
	poly := Polygon(p.points)
	rgba := color.RGBA{0, 0, 0, 0xff}

	bounds := image.Rect(0, 0, 200, 200)
	mask := image.NewAlpha(bounds)
	img := image.NewRGBA(bounds)
	rasterizer := raster.NewRasterizer(200, 200)
	rasterizer.UseNonZeroWinding = false
	rasterizer.Start(raster.Point{raster.Fix32(10 * 256), raster.Fix32(190 * 256)})
	for j := 0; j < len(poly); j = j + 2 {
		rasterizer.Add1(raster.Point{raster.Fix32(poly[j] * 256), raster.Fix32(poly[j+1] * 256)})
	}
	painter := raster.NewAlphaSrcPainter(mask)
	rasterizer.Rasterize(painter)
	DrawSolidRGBA(img, mask, rgba)
	savepng("_testFreetype.png", img)
}

func TestFreetypeNonZeroWinding(t *testing.T) {
	var p Path
	p.LineTo(10, 190)
	c := curve.CubicCurveFloat64{10, 190, 10, 10, 190, 10, 190, 190}
	c.Segment(&p, flattening_threshold)
	poly := Polygon(p.points)
	rgba := color.RGBA{0, 0, 0, 0xff}

	bounds := image.Rect(0, 0, 200, 200)
	mask := image.NewAlpha(bounds)
	img := image.NewRGBA(bounds)
	rasterizer := raster.NewRasterizer(200, 200)
	rasterizer.UseNonZeroWinding = true
	rasterizer.Start(raster.Point{raster.Fix32(10 * 256), raster.Fix32(190 * 256)})
	for j := 0; j < len(poly); j = j + 2 {
		rasterizer.Add1(raster.Point{raster.Fix32(poly[j] * 256), raster.Fix32(poly[j+1] * 256)})
	}
	painter := raster.NewAlphaSrcPainter(mask)
	rasterizer.Rasterize(painter)
	DrawSolidRGBA(img, mask, rgba)

	savepng("_testFreetypeNonZeroWinding.png", img)
}

func TestSimpleRasterizer(t *testing.T) {
	bounds := image.Rect(0, 0, 200, 200)
	img := image.NewRGBA(bounds)
	mask := image.NewAlpha(bounds)
	var p Path
	p.LineTo(10, 190)
	c := curve.CubicCurveFloat64{10, 190, 10, 10, 190, 10, 190, 190}
	c.Segment(&p, flattening_threshold)
	poly := Polygon(p.points)
	rgba := color.RGBA{0, 0, 0, 0xff}
	r := NewRasterizer()
	r.Fill(mask, poly)
	DrawSolidRGBA(img, mask, rgba)
	savepng("_testSimpleRasterizer.png", img)
}

func TestRasterizer(t *testing.T) {
	img := image.NewRGBA(image.Rect(0, 0, 200, 200))
	var p Path
	p.LineTo(10, 190)
	c := curve.CubicCurveFloat64{10, 190, 10, 10, 190, 10, 190, 190}
	c.Segment(&p, flattening_threshold)
	poly := Polygon(p.points)
	color := color.RGBA{0, 0, 0, 0xff}
	tr := [6]float64{1, 0, 0, 1, 0, 0}
	r := NewRasterizer8BitsSample(200, 200)
	//PolylineBresenham(img, image.Black, poly...)


	r.RenderEvenOdd(img, color, []Polygon{poly}, tr)
	savepng("_testRasterizer.png", img)
}

func TestRasterizerNonZeroWinding(t *testing.T) {
	img := image.NewRGBA(image.Rect(0, 0, 200, 200))
	var p Path
	p.LineTo(10, 190)
	c := curve.CubicCurveFloat64{10, 190, 10, 10, 190, 10, 190, 190}
	c.Segment(&p, flattening_threshold)
	poly := Polygon(p.points)
	color := color.RGBA{0, 0, 0, 0xff}
	tr := [6]float64{1, 0, 0, 1, 0, 0}
	r := NewRasterizer8BitsSample(200, 200)
	//PolylineBresenham(img, image.Black, poly...)


	r.RenderNonZeroWinding(img, color, []Polygon{poly}, tr)
	savepng("_testRasterizerNonZeroWinding.png", img)
}

func BenchmarkFreetype(b *testing.B) {
	var p Path
	p.LineTo(10, 190)
	c := curve.CubicCurveFloat64{10, 190, 10, 10, 190, 10, 190, 190}
	c.Segment(&p, flattening_threshold)
	poly := Polygon(p.points)
	rgba := color.RGBA{0, 0, 0, 0xff}
	
	for i := 0; i < b.N; i++ {
		bounds := image.Rect(0, 0, 200, 200)
		rasterizer := raster.NewRasterizer(200, 200)
		img := image.NewRGBA(bounds)
		mask := image.NewAlpha(bounds)	
		rasterizer.UseNonZeroWinding = false
		rasterizer.Start(raster.Point{raster.Fix32(10 * 256), raster.Fix32(190 * 256)})
		for j := 0; j < len(poly); j = j + 2 {
			rasterizer.Add1(raster.Point{raster.Fix32(poly[j] * 256), raster.Fix32(poly[j+1] * 256)})
		}
		painter := raster.NewAlphaSrcPainter(mask)
		rasterizer.Rasterize(painter)
		DrawSolidRGBA(img, mask, rgba)
	}
}
func BenchmarkFreetypeNonZeroWinding(b *testing.B) {
	var p Path
	p.LineTo(10, 190)
	c := curve.CubicCurveFloat64{10, 190, 10, 10, 190, 10, 190, 190}
	c.Segment(&p, flattening_threshold)
	poly := Polygon(p.points)
	rgba := color.RGBA{0, 0, 0, 0xff}
	
	for i := 0; i < b.N; i++ {
		bounds := image.Rect(0, 0, 200, 200)
		rasterizer := raster.NewRasterizer(200, 200)
		img := image.NewRGBA(bounds)
		mask := image.NewAlpha(bounds)	
		rasterizer.UseNonZeroWinding = true
		rasterizer.Start(raster.Point{raster.Fix32(10 * 256), raster.Fix32(190 * 256)})
		for j := 0; j < len(poly); j = j + 2 {
			rasterizer.Add1(raster.Point{raster.Fix32(poly[j] * 256), raster.Fix32(poly[j+1] * 256)})
		}
		painter := raster.NewAlphaSrcPainter(mask)
		rasterizer.Rasterize(painter)
		DrawSolidRGBA(img, mask, rgba)
	}
}

func BenchmarkRasterizerNonZeroWinding(b *testing.B) {
	var p Path
	p.LineTo(10, 190)
	c := curve.CubicCurveFloat64{10, 190, 10, 10, 190, 10, 190, 190}
	c.Segment(&p, flattening_threshold)
	poly := Polygon(p.points)
	color := color.RGBA{0, 0, 0, 0xff}
	tr := [6]float64{1, 0, 0, 1, 0, 0}
	rasterizer := NewRasterizer8BitsSample(200, 200)
	for i := 0; i < b.N; i++ {
		img := image.NewRGBA(image.Rect(0, 0, 200, 200))
		
		rasterizer.RenderNonZeroWinding(img, color, []Polygon{poly}, tr)
	}
}

func BenchmarkRasterizer(b *testing.B) {
	var p Path
	p.LineTo(10, 190)
	c := curve.CubicCurveFloat64{10, 190, 10, 10, 190, 10, 190, 190}
	c.Segment(&p, flattening_threshold)
	poly := Polygon(p.points)
	color := color.RGBA{0, 0, 0, 0xff}
	tr := [6]float64{1, 0, 0, 1, 0, 0}
	rasterizer := NewRasterizer8BitsSample(200, 200)
	for i := 0; i < b.N; i++ {
		img := image.NewRGBA(image.Rect(0, 0, 200, 200))
		
		rasterizer.RenderEvenOdd(img, color, []Polygon{poly}, tr)
	}
}

func BenchmarkSimpleRasterizer(b *testing.B) {
	var p Path
	p.LineTo(10, 190)
	c := curve.CubicCurveFloat64{10, 190, 10, 10, 190, 10, 190, 190}
	c.Segment(&p, flattening_threshold)
	poly := Polygon(p.points)
	rgba := color.RGBA{0, 0, 0, 0xff}
	rasterizer := NewRasterizer()
	for i := 0; i < b.N; i++ {
		bounds := image.Rect(0, 0, 200, 200)
		img := image.NewRGBA(bounds)
		mask := image.NewAlpha(bounds)	
		rasterizer.Fill(mask, poly)
		DrawSolidRGBA(img, mask, rgba)
	}
}
