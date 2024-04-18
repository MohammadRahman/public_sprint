package sprint

import "fmt"

type Point struct {
	X    float32
	Y    float32
	Text string
}

func roundFloat(f float32, places int) float32 {
	scale := float32(1)
	for i := 0; i < places; i++ {
		scale *= 10
	}
	return float32(int(f*scale+0.5)) / scale
}

func PointText(p Point) Point {

	xRounded := roundFloat(p.X, 2)
	yRounded := roundFloat(p.Y, 2)

	formattedText := fmt.Sprintf("Text at (%.2f, %.2f)", xRounded, yRounded)

	return Point{
		X:    p.X,
		Y:    p.Y,
		Text: formattedText,
	}
}
