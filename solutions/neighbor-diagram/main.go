package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"math/rand"
	"os"
	"time"
)

const ImageSize = 1000
const PointCount = 30

type Point struct {
	X float64
	Y float64
}

func RandomPoint() Point {
	return Point{rand.Float64(), rand.Float64()}
}

func (p Point) Distance(p1 Point) float64 {
	return math.Sqrt(math.Pow(p1.X-p.X, 2) + math.Pow(p1.Y-p.Y, 2))
}

func RandomColor() color.Color {
	return color.RGBA{
		R: uint8(rand.Intn(256)),
		G: uint8(rand.Intn(256)),
		B: uint8(rand.Intn(256)),
		A: 0xff,
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage:", os.Args[0], "<output.png>")
		os.Exit(1)
	}

	rand.Seed(time.Now().UnixNano())
	points := make([]Point, PointCount)
	colors := make([]color.Color, PointCount)
	for i := range points {
		points[i] = RandomPoint()
		colors[i] = RandomColor()
	}
	res := image.NewRGBA(image.Rect(0, 0, ImageSize, ImageSize))
	for x := 0; x < ImageSize; x++ {
		for y := 0; y < ImageSize; y++ {
			point := Point{float64(x) / (ImageSize - 1), float64(y) / (ImageSize - 1)}
			pickIndex := 0
			distance := point.Distance(points[0])
			for i, p := range points[1:] {
				if p.Distance(point) < distance {
					distance = p.Distance(point)
					pickIndex = i
				}
			}
			if distance < 0.005 {
				res.Set(x, y, color.Gray{Y: 0})
			} else {
				res.Set(x, y, colors[pickIndex])
			}
		}
	}
	output, err := os.Create(os.Args[1])
	if err != nil {
		panic(err)
	}
	png.Encode(output, res)
	output.Close()
}

func roundNum(n float64) int {
	if math.Abs(math.Ceil(n)-n) <= math.Abs(math.Floor(n)-n) {
		return int(math.Ceil(n))
	} else {
		return int(math.Floor(n))
	}
}
